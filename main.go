package main

import (
	"io"
	"net/http"
)

func serverfeed(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Simple web server!,its pretty snazzy right?, or if you prefer yo can add some html like in static-file ")
}

func main() {
	http.HandleFunc("/", serverfeed)
	http.ListenAndServe(":8000", nil)
}
