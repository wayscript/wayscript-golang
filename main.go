package main

import (
	"fmt"

	wayscript "github.com/wayscript/wayscript-golang/wayscript"
)

func main() {
	event, err := wayscript.GetEvent()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Event:")
	fmt.Println(event)

	userString, err := wayscript.GetUserByApplicationKey()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("User String:")
	fmt.Println(userString)

	fmt.Println("http-trigger: ")

	var jsonData = []byte(`
		{
			"data": "data response from http-trigger goes here",
			"headers": {"header":"value"},
			"status_code": 200
		}
	`)
	triggerString, err := wayscript.SendHttpTriggerResponse(jsonData)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(triggerString)
	}

}
