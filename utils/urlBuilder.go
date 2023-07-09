package utils

import "fmt"

func BuildUrl(secret string, action string) string {

	functionUrl := "https://hothheim-function-start-stop.azurewebsites.net/api/startstophothheimserver"

	functionCodeParam := fmt.Sprintf("code=%s", secret)

	actionParam := fmt.Sprintf("action=%s", action)

	return fmt.Sprintf("%s?%s&%s", functionUrl, functionCodeParam, actionParam)
}
