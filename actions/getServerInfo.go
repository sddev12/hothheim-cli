package actions

import (
	"encoding/json"
	"fmt"
	"hothheim-cli/utils"
	"io"
	"log"
	"net/http"
)

func ExecuteServerInfo(secret string) {
	utils.ClearConsole()
	utils.Banner()
	doneChan := make(chan bool)
	go utils.LoadAnimation(doneChan)

	errorChan := make(chan error)
	responseChan := make(chan GetStatusResponse)

	go func() {
		getStatusResponse, responseError := GetServerInfo(secret)

		errorChan <- responseError
		responseChan <- getStatusResponse
	}()

	respError := <-errorChan
	getStatusResponse := <-responseChan
	doneChan <- true

	if respError != nil {
		if responseError, ok := respError.(ResponseError); ok {
			fmt.Println(responseError.FriendlyErrorMessage)
			fmt.Println(responseError.ErrorCode)
		} else {
			fmt.Println(respError.Error())
		}
	} else {
		utils.PrintServerStatus(getStatusResponse.CurrentState)
	}
}

func GetServerInfo(secret string) (GetStatusResponse, error) {

	url := utils.BuildUrl(secret, "GetStatus")

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("ERROR: Failed to get server information")
	}

	if resp.StatusCode == 401 {
		log.Fatal("ERROR: Incorrect password")
	}

	if resp.StatusCode == 500 {

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal("Error: Failed to read response body")
		}
		defer resp.Body.Close()

		var responseError ResponseError

		parseErr := json.Unmarshal(body, &responseError)
		if parseErr != nil {
			log.Fatal("ERROR: Filed to unmarshall response body")
		}

		return GetStatusResponse{}, responseError
	}

	var getStatusResponse GetStatusResponse

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error: Failed to read response body")
	}
	defer resp.Body.Close()

	getStatusResponseError := json.Unmarshal(body, &getStatusResponse)
	if getStatusResponseError != nil {
		log.Fatalf("Error: Failed to unmarshal response body: %s", getStatusResponseError)
	}

	return getStatusResponse, nil
}
