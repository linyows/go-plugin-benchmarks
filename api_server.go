package main

import (
	"fmt"
	"net/http"
)

const APIServerAddr = ":9000"

func ServeAPI() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()

		if len(q) == 0 {
			return
		}

		for k, v := range q {
			if k == "prefix" {
				fmt.Fprintf(w, "%s, http api\n", v)
			}
		}
	})
	http.ListenAndServe(APIServerAddr, nil)
}
