package linnroutes

import (
	"fmt"
	"net/http"
	"quaysports.com/server/pkg/linn"
	"strings"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	trimmedPath := strings.Trim(r.URL.Path[1:], "/")
	path := strings.Split(trimmedPath, "/")
	fmt.Println(len(path))
	fmt.Println(path)
	if len(path) == 2 {
		switch path[1] {
		case "PostalServices":
			channel := make(chan string)
			go linn.GetPostalServices(channel)
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintf(w, "%s", <-channel)
		default:
			fmt.Println("no extension")
		}
	} else {
		fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
	}
}
