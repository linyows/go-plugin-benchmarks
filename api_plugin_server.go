package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const APIPluginServerAddr = ":8010"

func apiReq(prefix string) []byte {
	url := fmt.Sprintf("http://localhost:9000/?prefix=%s", prefix)
	res, err := http.Get(url)
	if err != nil {
		fmt.Printf("GET error: %s\n", err)
		return []byte("")
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("read error: %s\n", err)
		return []byte("")
	}

	return body
}

func ServeAPIPlugin() {
	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		te := apiReq(Name)
		if len(te) > 0 {
			fmt.Fprintln(w, te)
		}
	})
	http.ListenAndServe(APIPluginServerAddr, nil)
}
