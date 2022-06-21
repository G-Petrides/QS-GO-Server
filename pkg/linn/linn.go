package linn

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"quaysports.com/server/pkg/core"
	"time"
)

type AuthData struct {
	Server string
	Token  string
}

var AuthProperties AuthData

func Init(config core.Config, channel chan bool) {

	err := Auth(config)
	fmt.Println("Linn Auth!")
	if err != nil {
		log.Fatalf("Linn Auth Init Error: %s", err.Error())
	}
	channel <- true

	for range time.Tick(time.Second * 20) {
		err := Auth(config)
		fmt.Println("Linn Auth!")
		if err != nil {
			log.Fatalf("Linn Auth Init Error: %s", err.Error())
		}
	}
}

func Auth(config core.Config) error {

	jsonEncoding, err := json.Marshal(config)
	if err != nil {
		return err
	}

	body := bytes.NewBuffer([]byte("request=" + string(jsonEncoding)))

	res, err := http.Post("https://api.linnworks.net/api/AuthProperties/AuthorizeByApplication", "application/json", body)
	if err != nil {
		return err
	}

	buffer := new(bytes.Buffer)
	_, err = buffer.ReadFrom(res.Body)
	if err != nil {
		return err
	}

	_ = json.Unmarshal(buffer.Bytes(), &AuthProperties)

	return nil
}

func PostReq(url string, bodyData string) (string, error) {

	if AuthProperties.Server != "" {
		url = AuthProperties.Server + url
		fmt.Println(url)
	}

	body := bytes.NewBuffer([]byte(bodyData))

	client := &http.Client{
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("Accept", "application/json")
	if AuthProperties.Token != "" {
		fmt.Println(AuthProperties.Token)
		req.Header.Set("Authorization", AuthProperties.Token)
	}
	res, err := client.Do(req)
	if res.StatusCode == http.StatusOK {
		if err != nil {
			return "", err
		}

		defer res.Body.Close()

		resBody, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return "", nil
		}

		return string(resBody), nil
	} else {
		return "", errors.New(res.Status)
	}
}
