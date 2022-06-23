package main

import (
	"fmt"
	"log"
	"net/http"
	"quaysports.com/server/pkg/core"
	"quaysports.com/server/pkg/linn"
	"quaysports.com/server/pkg/routes"
)

func main() {
	config, err := core.GetConfig()
	if err != nil {
		fmt.Println(err)
	}

    fmt.Println(config)
	fmt.Println("Server Started!")
    done := make(chan linn.InitResult)
	go linn.Init(config, done)
    result := <-done
    if result.Err != nil {
        log.Fatalf("Linn Init Error: %s", result.Err.Error())
    }

	http.HandleFunc("/Linn/", linnroutes.Handler)
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("HTTP Server error: %s", err.Error())
	}

	postData, err := linn.PostReq("/api/Inventory/GetPostalServices", "")
	if err != nil {
		log.Fatalf("Linn HTTP error: %s", err.Error())
	}
	fmt.Println(postData)
}
