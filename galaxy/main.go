package galaxy

import (
	"github.com/minio/minio-go"
	"../utils"
	"io/ioutil"
	"os"
	"time"
	"strings"
	"log"
	"io"
	"regexp"
	"strconv"
	"errors"
)

var GointStorage *minio.Client

func InitGointStorage(config *utils.GointConfig) error{

	endpoint := config.Storage.Endpoint
	accessKeyID := config.Storage.AccessKey
	secretAccessKey := config.Storage.SecretKey
	useSSL := false

	var err error
	GointStorage, err = minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
	if err != nil {
		return err
	}
	return nil
}

func fixName (name string) (string, error) {
	fixedName := strings.Replace(name, " ", "_", -1)
	var extension string
	if strings.Contains(fixedName, ".") {
		parts := strings.Split(fixedName, ".")
		extension = parts[len(parts)-1]

		if !strings.EqualFold(extension, "png") {
			return "", errors.New("currently, png image type only accepted")
		}
		return fixedName,  nil
	}
	return fixedName + ".png", nil

}

func UploadPhoto(image io.Reader, name string, bucketName string) (string, error) {

	// Preparing the temporal file to upload
	data, err := ioutil.ReadAll(image)
	if err != nil {
		return "", err
	}

	file, err := ioutil.TempFile(os.TempDir(), "goint")
	if err != nil {
		return "", err
	}
	defer os.Remove(file.Name())

	_, err = file.Write(data)
	if err != nil {
		return "", err
	}

	// Preparing the lcp bucket (A unique bucket), I don't know if it is the most efficient way
	exists, err := GointStorage.BucketExists(bucketName)
	if err != nil {
		return "", err
	}
	if !exists {
		GointStorage.MakeBucket(bucketName, bucketsLocation)
	}

	// Generating the name of file:


	// Creating the file

	imageName, err  := fixName(name)
	if err != nil {
		return "", err
	}

	// Check if object already exist

	log.Println(imageName)
	fileForTest, _ := ioutil.TempFile(os.TempDir(), "test_goint")

	err = GointStorage.FGetObject(bucketName, imageName, fileForTest.Name(), minio.GetObjectOptions{})
	fileForTest.Close()
	os.Remove(fileForTest.Name())
	log.Println(err)
	if err == nil { // It's triggered cause image already exists, then I need change the name
		// add a number to suffix
		chunk := strings.Split(imageName, ".")
		onlyName := chunk[0]
		extension := chunk[len(chunk)-1]

		reg, _ := regexp.Compile("\\w+_[0-9]+")
		if reg.MatchString(onlyName) {
			parts := strings.Split(onlyName, "_")
			count, _ := strconv.Atoi(parts[len(parts)-1])
			count = count + 1

			nextName := strings.Join(parts[:len(parts)-1], "_") + "_" + strconv.Itoa(count)

			imageName = nextName + "." + extension

		} else {
			imageName = onlyName + "_1." +  extension
		}
	} else {
		if strings.Contains(err.Error(), "no such file") || strings.Contains(err.Error(), "not exist") {
		} else {
			return "", err
		}

	}


	filePath := file.Name()
	contentType := "image/png"

	// Upload the zip file with FPutObject
	_, err = GointStorage.FPutObject(
		bucketName,
		imageName,
		filePath,
		minio.PutObjectOptions{ContentType:contentType},
	)
	if err != nil {
		return "", err
	}

	return imageName, nil

}

func DownloadPhoto(name string, bucketName string) (string, error) {
	exists, err := GointStorage.BucketExists(bucketName)
	if err != nil {
		return "", err
	}
	if !exists {
		return "", errors.New("bucket not exist, then neither exist your file")
	}

	objectPhoto, err := GointStorage.GetObject(bucketName, name, minio.GetObjectOptions{})
	if err != nil {
		return "", err
	}

	data, err := ioutil.ReadAll(objectPhoto)
	if err != nil {
		return "", err
	}

	file, err := ioutil.TempFile(os.TempDir(), "goint")
	if err != nil {
		return "", err
	}

	// Very interesting use for 10 seconds of life for temporal file
	defer func() {
		go func() {
			time.Sleep(10)
			os.Remove(file.Name())
		}()
	}()

	_, err = file.Write(data)
	if err != nil {
		return "", err
	}

	return file.Name(), nil
}

func GetAllImagesFromBucket(bucketName string) []string {
	var done = make(chan struct{}, 1)
	fileStream := GointStorage.ListObjectsV2(bucketName, "", true, done)

	allFiles := make([]string, 0)
	for fileInfo := range fileStream {
		name := fileInfo.Key
		log.Println(fileInfo.Key)
		allFiles = append(allFiles, bucketName + "/" + name)
	}

	return allFiles
}