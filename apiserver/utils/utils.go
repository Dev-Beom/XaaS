package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func InterfaceToBuffer(model interface{}) (bytes.Buffer, error) {
	buffer := new(bytes.Buffer)
	err := json.NewEncoder(buffer).Encode(model)
	if err != nil {
		return bytes.Buffer{}, err
	}
	return *buffer, nil
}

func GetFileExtension(filename string) (string, error) {
	slice := strings.Split(filename, ".")
	if len(slice) < 1 {
		return "", errors.New("파일의 확장자를 추출할 수 없습니다")
	}
	return slice[len(slice)-1], nil
}

func LoggingByTest(msg string, code int, body string) {
	fmt.Println("[Test Message]\t" + msg)
	fmt.Println("[Status Code]\t" + strconv.Itoa(code))
	fmt.Println("[Response Body]\t\n" + body)
}
