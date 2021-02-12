package utils

import (
	"log"
	"net/http"
	"strings"
)

// CheckErr 에러체크
func CheckErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

// CheckCode Response 체크
func CheckCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request Failed with Status:", res.StatusCode)
	}
}

// CleanString 문자열 앞뒤의 모든 스페이스 등을 다 잘라내어 깨끗하게 함
func CleanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}
