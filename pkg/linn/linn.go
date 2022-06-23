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

type authData struct {
	Server string
	Token  string
}

var serverAuthData authData

//InitResult : struct for Init functions channel
type InitResult struct {
	authData authData
    //Publicly accessable error for logging
	Err      error
}

//Init function called from main to auth with Linnworks and keep connection alive.
func Init(config core.Config, channel chan InitResult) {

	var authChannel = make(chan error)
	go auth(config, authChannel)
	err := <-authChannel
	if err != nil {
		channel <- InitResult{serverAuthData, err}
	}

	channel <- InitResult{serverAuthData, nil}

	go scheduleReAuth(config)
}

func auth(config core.Config, channel chan error) {

	jsonEncoding, err := json.Marshal(config.Linn)
	if err != nil {
		channel <- err
	}

	body := bytes.NewBuffer([]byte("request=" + string(jsonEncoding)))

	client := &http.Client{
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest("POST", "https://api.linnworks.net/api/Auth/AuthorizeByApplication", body)
	if err != nil {
		channel <- err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("Accept", "application/json")

	res, err := client.Do(req)
	if err != nil {
		channel <- err
	}

	defer res.Body.Close()

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		channel <- err
	}

	_ = json.Unmarshal(resBody, &serverAuthData)

	channel <- nil
}

func scheduleReAuth(config core.Config) {
	duration := time.Duration(time.Second * 20)
	var wrap func()
	wrap = func() {
		fmt.Println("Auth wrap!")
		var newAuthChannel = make(chan error)
		go auth(config, newAuthChannel)
		err := <-newAuthChannel
		if err != nil {
			log.Fatalf("Linn Init Auth Error: %s", err.Error())
		}
		time.AfterFunc(duration, wrap)

	}
	wrap()
}

//PostReq is a public function for making API post request to Linnworks servers.
//Uses private stored authData for authentication
func PostReq(url string, bodyData string) (string, error) {

	fmt.Println(serverAuthData)

	body := bytes.NewBuffer([]byte(bodyData))

	client := &http.Client{
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest("POST", serverAuthData.Server+url, body)
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", serverAuthData.Token)

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
	}
        
	return "", errors.New(res.Status)

}
