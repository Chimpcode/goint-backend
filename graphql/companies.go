package graphql

import (
	"io/ioutil"
	"log"
	"strings"
	"fmt"
	"net/http"
	"encoding/json"
	"errors"
)

// var endpoint string = "https://api.graph.cool/simple/v1/cjdchobh82rgk0146bw4p4kyu"
var endpoint string = "https://api.graph.cool/simple/v1/cjlborupt7iji0181iefxheiv"

// GC = Graphcool
func GetCompanyFromGCbyEmail(email string) (*MiniCompany, error) {
	var user MiniCompany


	log.Printf("Searching with '%s' email", email)


	vars := map[string]interface{}{
		"email": email,
	}

	dataVars, err := json.Marshal(vars)
	if err != nil {
		return &user, err
	}

	dataVarsString := strings.Replace(string(dataVars), "\u0022", "\u005c\u0022", -1)

	fullBody := fmt.Sprintf("{\"query\": \"%s\", \"variables\":\"%s\"}", CompanyByEmail, dataVarsString)
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

