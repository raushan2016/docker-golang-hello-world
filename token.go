package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type responseJson struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    string `json:"expires_in"`
	ExpiresOn    string `json:"expires_on"`
	NotBefore    string `json:"not_before"`
	Resource     string `json:"resource"`
	TokenType    string `json:"token_type"`
}

func main1() {

	// Create HTTP request for a managed services for Azure resources token to access Azure Resource Manager
	var msi_endpoint *url.URL
	msi_endpoint, err := url.Parse("http://169.254.169.254/metadata/identity/oauth2/token?api-version=2018-02-01")
	if err != nil {
		fmt.Println("Error creating URL: ", err)
		return
	}
	msi_parameters := url.Values{}
	msi_parameters.Add("resource", "https://management.azure.com/")
	msi_endpoint.RawQuery = msi_parameters.Encode()
	req, err := http.NewRequest("GET", msi_endpoint.String(), nil)
	if err != nil {
		fmt.Println("Error creating HTTP request: ", err)
		return
	}
	req.Header.Add("Metadata", "true")

	// Call managed services for Azure resources token endpoint
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error calling token endpoint: ", err)
		return
	}

	// Pull out response body
	responseBytes, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		fmt.Println("Error reading response body : ", err)
		return
	}

	// Unmarshall response body into struct
	var r responseJson
	err = json.Unmarshal(responseBytes, &r)
	if err != nil {
		fmt.Println("Error unmarshalling the response:", err)
		return
	}

	// Print HTTP response and marshalled response body elements to console
	fmt.Println("Response status:", resp.Status)
	fmt.Println("access_token: ", r.AccessToken)
	fmt.Println("refresh_token: ", r.RefreshToken)
	fmt.Println("expires_in: ", r.ExpiresIn)
	fmt.Println("expires_on: ", r.ExpiresOn)
	fmt.Println("not_before: ", r.NotBefore)
	fmt.Println("resource: ", r.Resource)
	fmt.Println("token_type: ", r.TokenType)
}
