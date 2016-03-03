package main

import (
	"fmt"
	"./config"
)

func main() {
	fmt.Println("Starting Server... ")
	config.InitMVCRequestConfig()
}
