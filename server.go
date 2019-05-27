package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type todo struct {
	Text string
	// Date *time.Time
}

// Result sent to client
type Result struct {
	Result  []Match `json:"result"`
	Pending bool    `json:"pending"`
}

func startServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)

		if (*r).Method == "OPTIONS" {
			return
		}

		setDefaultHeaders(&w)

		// process incoming data
		decoder := json.NewDecoder(r.Body)
		var data todo
		err := decoder.Decode(&data)
		if err != nil {
			panic(err)
		}

		matches := findMatches(data.Text)
		result := Result{matches, false}

		b, err := json.Marshal(result)
		if err != nil {
			panic(err)
		}

		fmt.Fprint(w, string(b))
	})

	log.Fatal(http.ListenAndServe(":1000", nil))
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func setDefaultHeaders(w *http.ResponseWriter) {
	(*w).Header().Set("Content-Type", "application/json; charset=UTF-8")
}
