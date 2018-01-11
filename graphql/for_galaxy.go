package graphql

import (
	"log"
	"strings"
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"errors"
)

func GetCompanyFromGCbySocialReason(socialReason string) (*MiniCompany, error) {
	var user MiniCompany

	log.Printf("Searching with '%s' socialReason", socialReason)

	endpoint := "http://13.90.253.208:60000/simple/v1/cjcae1ay000en0189jqrz4n2q"

	vars := map[string]interface{}{
		"social": socialReason,
	}

	dataVars, err := json.Marshal(vars)
	if err != nil {
		return &user, err
	}

	dataVarsString := strings.Replace(string(dataVars), "\u0022", "\u005c\u0022", -1)

	fullBody := fmt.Sprintf("{\"query\": \"%s\", \"variables\":\"%s\"}", CompanyBySocialReason, dataVarsString)
	var buf = strings.NewReader(fullBody)

	resp, err := http.Post(endpoint, "application/json", buf)

	if err != nil {
		return &user, err
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return &user, err
	}

	log.Printf("%s", respBody)

	type Response struct {
		Data map[string]MiniCompany `json:"data"`
		Error []string `json:"error"`
	}

	var finalResp Response

	err = json.Unmarshal(respBody, &finalResp)
	if err != nil {
		return &user, err
	}

	log.Println(finalResp)


	user = finalResp.Data["Company"]

	if strings.EqualFold(user.Email, "") {
		return &user, errors.New("user not exist")
	}
	return &user, nil

}

func GetCompanyFromGCbyRuc(ruc string) (*MiniCompany, error) {
	var user MiniCompany

	log.Printf("Searching with '%s' ruc", ruc)

	endpoint := "http://13.90.253.208:60000/simple/v1/cjcae1ay000en0189jqrz4n2q"

	vars := map[string]interface{}{
		"ruc": ruc,
	}

	dataVars, err := json.Marshal(vars)
	if err != nil {
		return &user, err
	}

	dataVarsString := strings.Replace(string(dataVars), "\u0022", "\u005c\u0022", -1)

	fullBody := fmt.Sprintf("{\"query\": \"%s\", \"variables\":\"%s\"}", CompanyByRuc, dataVarsString)
	var buf = strings.NewReader(fullBody)

	resp, err := http.Post(endpoint, "application/json", buf)

	if err != nil {
		return &user, err
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return &user, err
	}

	log.Printf("%s", respBody)

	type Response struct {
		Data map[string]MiniCompany `json:"data"`
		Error []string `json:"error"`
	}

	var finalResp Response

	err = json.Unmarshal(respBody, &finalResp)
	if err != nil {
		return &user, err
	}

	log.Println(finalResp)


	user = finalResp.Data["Company"]

	if strings.EqualFold(user.Email, "") {
		return &user, errors.New("user not exist")
	}
	return &user, nil

}


func CheckIfCompanyExistBysocialReason(socialReason string) (bool, error) {
	_, err := GetCompanyFromGCbySocialReason(socialReason)
	if err != nil {
		return false, err
	}
	return true, nil
}

func CheckIfCompanyExistByRuc(ruc string) (bool, error) {
	_, err := GetCompanyFromGCbyRuc(ruc)
	if err != nil {
		return false, err
	}
	return true, nil
}
