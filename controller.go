package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func FindPairs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//reading the request body
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		//writing error response in-case of an error
		w.WriteHeader(400)
		err := Error{}
		err.Code = "ServerError"
		err.Message = "error while reading request body"
		log.Fatal(err)
		json.NewEncoder(w).Encode(err)
	}

	//unmarshalling the request body to the relevant request struct
	var request Request
	err = json.Unmarshal(reqBody, &request)
	if err != nil {
		//writing error response in-case of an error
		w.WriteHeader(400)
		err := Error{}
		err.Code = "ServerError"
		err.Message = "error while unmarshalling request body"
		log.Fatal(err)
		json.NewEncoder(w).Encode(err)
	}

	//calling the function to calculate the answer
	result := request.calculate()

	//writing success response
	log.Println(result)
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(result)
}
