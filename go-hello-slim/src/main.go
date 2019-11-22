package main

import (
		"io"
        "log"
        "net/http"
)

func MyHandleFunc(w http.ResponseWriter, req *http.Request) {
        io.WriteString(w, "Hello World\n")
}

func main() {
        http.HandleFunc("/", MyHandleFunc)
        log.Fatal(http.ListenAndServe(":3000", nil))
}