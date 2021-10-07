package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/healthz", healthz)
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/getEnvs", getEnvs)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}

}

func getEnvs(w http.ResponseWriter, r *http.Request) {
	env := os.Getenv("GOPATH")
	if env != "" {
		io.WriteString(w, fmt.Sprintf("env info %s\n", env))
	} else {
		io.WriteString(w, "env is nil\n")
	}
	fmt.Println("request ip: ", r.Host, "retcode: ", http.StatusOK)
}

func headers(w http.ResponseWriter, r *http.Request) {
	for k, v := range r.Header {
		w.Header().Add(k, v[0])
	}
	fmt.Println("request ip: ", r.Host, "retcode: ", http.StatusOK)
}

func healthz(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "200\n")
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("entering root handler")
	user := r.URL.Query().Get("user")
	if user != "" {
		io.WriteString(w, fmt.Sprintf("hello [%s]\n", user))
	} else {
		io.WriteString(w, "hello [stranger]\n")
	}
	io.WriteString(w, "===================Details of the http request header:============\n")
	for k, v := range r.Header {
		io.WriteString(w, fmt.Sprintf("%s=%s\n", k, v))
	}
}
