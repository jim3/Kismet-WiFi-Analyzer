package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	SSID       string `json:"kismet.device.base.name"`
	Encryption string `json:"kismet.device.base.crypt"`
	MacAddr    string `json:"kismet.device.base.macaddr"`
}

func upload2parser(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("uploadfile")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Print out header for debugging
	fmt.Fprintf(w, "%v", header.Header)

	// Create a nil slice
	var resp []Response

	// Create a new decoder
	decoder := json.NewDecoder(file)

	// Decode the response body into the resp slice
	if err := decoder.Decode(&resp); err != nil {
		fmt.Println(err)
	}

	// If no error occurs, we can use the slice of items in our program
	for _, v := range resp {
		fmt.Println(v.SSID)
		fmt.Println(v.Encryption)
		fmt.Println(v.MacAddr)
	}

	fmt.Println(resp) // [{CatheadBiscuits WPA2 WPA2-PSK AES-CCMP 38:3F:B3:84:63:F8}]
}

func main() {
	// Create a new servemux
	mux := http.NewServeMux()
	mux.HandleFunc("/upload", upload2parser)
	mux.Handle("/", http.FileServer(http.Dir("/github.com/<USERNAME>/project-dir")))

	// Create a new http.Server struct
	s := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	s.ListenAndServe()
}
