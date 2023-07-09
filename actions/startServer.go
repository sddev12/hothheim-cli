package actions

import (
	"encoding/json"
	"fmt"
	"hothheim-cli/utils"
	"io"
	"log"
	"net/http"
)

func ExecuteStartServer(secret string) {
	utils.ClearConsole()
	utils.Banner()
	doneChan := make(chan bool)
	go utils.LoadAnimation(doneChan)

	errorChan := make(chan error)
	responseChan := make(chan bool)

	go func() {
		response, responseError := StartServer(secret)

		errorChan <- responseError
		responseChan <- response
	}()

	respError := <-errorChan
	response := <-responseChan
	doneChan <- true

	if respError != nil {
		if responseError, ok := respError.(ResponseError); ok {
			fmt.Println(responseError.FriendlyErrorMessage)
			fmt.Println(responseError.ErrorCode)
		} else {
			fmt.Println(respError.Error())
		}
	} else if response {
		ExecuteServerInfo(secret)
	}
}

func StartServer(secret string) (bool, error) {

	url := utils.BuildUrl(secret, "Start")

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("ERROR: Failed to request server start")
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

		return false, responseError
	}

	return true, nil
}
