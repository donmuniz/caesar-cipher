package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

type cipherHttp struct {
	cipher *cipher
}

type cipherBody struct {
	Msg string `json:"msg"`
}

func (cHttp *cipherHttp) handleReq(w http.ResponseWriter, r *http.Request, decode bool) {
	if r.Method == "POST" {
		defer r.Body.Close()

		request := cipherBody{}
		err := json.NewDecoder(r.Body).Decode(&request)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			http.Error(w, "400 bad request", http.StatusBadRequest)
			return
		}

		msg := strings.ToUpper(string(request.Msg))

		if decode {
			msg = cHttp.cipher.Decode(msg)
		} else {
			msg = cHttp.cipher.Encode(msg)
		}

		responseInstance := cipherBody{msg}
		responseJson, _ := json.Marshal(responseInstance)

		w.WriteHeader(http.StatusOK)
		w.Write(responseJson)
	} else {
		w.WriteHeader(http.StatusNotFound)
		http.Error(w, "404 page not found", http.StatusNotFound)
	}
}

func (cHttp *cipherHttp) Decode(w http.ResponseWriter, r *http.Request) {
	cHttp.handleReq(w, r, true)
}

func (cHttp *cipherHttp) Encode(w http.ResponseWriter, r *http.Request) {
	cHttp.handleReq(w, r, false)
}
