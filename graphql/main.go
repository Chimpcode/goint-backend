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


type MiniCompany struct {
	Id             string
	CommercialName string
	Email          string
	Password       string
	Ruc string
	SocialReason string

}

var companyByEmail = `query ($email: String!){Company (email: $email) {id commercialName ruc socialReason email password}}`

// GC = Graphcool
func GetCompanyFromGCbyEmail(email string) (*MiniCompany, error) {
	var user MiniCompany

	log.Printf("Searching with '%s' email", email)

	endpoint := "http://13.90.253.208:60000/simple/v1/cjcae1ay000en0189jqrz4n2q"

	vars := map[string]interface{}{
		"email": email,
	}

	dataVars, err := json.Marshal(vars)
	if err != nil {
		return &user, err
	}

	dataVarsString := strings.Replace(string(dataVars), "\u0022", "\u005c\u0022", -1)

	fullBody := fmt.Sprintf("{\"query\": \"%s\", \"variables\":\"%s\"}", companyByEmail, dataVarsString)
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
