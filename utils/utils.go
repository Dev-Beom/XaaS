package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
)

func InterfaceToBuffer(model interface{}) (bytes.Buffer, error) {
	buffer := new(bytes.Buffer)
	err := json.NewEncoder(buffer).Encode(model)
	if err != nil {
		return bytes.Buffer{}, err
	}
	return *buffer, nil
}

func LoggingByTest(msg string, code int, body string) {
	fmt.Println("[Test Message]\t" + msg)
	fmt.Println("[Status Code]\t" + strconv.Itoa(code))
	fmt.Println("[Response Body]\t\n" + body)
}
