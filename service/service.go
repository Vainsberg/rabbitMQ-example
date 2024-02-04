package service

import "fmt"

func ProcessMessage(message []byte) {
	messageStr := string(message)
	fmt.Println(messageStr)
}
