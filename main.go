package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	var whatIsIt string

	secret := "aWFuZ25vVzpOQU06RU5JTDp0YTpzdTpuaW9K"

	sd, _ := base64.StdEncoding.DecodeString(secret)
	whatIsIt = string(sd)
	fmt.Println("Decode Base64: " + whatIsIt)
	fmt.Println("Answer: " + Reverse(whatIsIt))

}

func Reverse(input string) (result string) {
	for _, c := range input {
		result = string(c) + result
	}
	return
}
