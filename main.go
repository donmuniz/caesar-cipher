package main

import (
	"log"
	"net/http"
)

func main() {
	port := getenv("PORT")
	pass := getenv("CIPHER_PASS")

	startChar := 'A'
	endChar := 'Z'

	cipherInstance := cipher{pass, startChar, endChar}
	cipherHttp := cipherHttp{cipher: &cipherInstance}

	router := http.NewServeMux()
	router.Handle("/encode", http.HandlerFunc(cipherHttp.Encode))
	router.Handle("/decode", http.HandlerFunc(cipherHttp.Decode))

	log.Printf("starting server at port: %s\n", port)
	http.ListenAndServe(":"+port, router)
}
