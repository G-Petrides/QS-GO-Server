package linn

import (
	"log"
)

func GetPostalServices(channel chan string) {
	postalServices, err := PostReq("/api/PostalServices/GetPostalServices", "")
	if err != nil {
		log.Fatalf("linn HTTP Error: %s", err.Error())
	}
	channel <- postalServices
}
