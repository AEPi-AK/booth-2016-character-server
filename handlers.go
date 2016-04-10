package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"io"
	"io/ioutil"

	"github.com/gorilla/mux"
)

func CharacterCreate(w http.ResponseWriter, r *http.Request) {
    var requestData CreateRequest 
    body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
    if err != nil {
        panic(err)
    }
    if err := r.Body.Close(); err != nil {
        panic(err)
    }
    if err := json.Unmarshal(body, &requestData); err != nil {
        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        w.WriteHeader(422) // unprocessable entity
        if err := json.NewEncoder(w).Encode(err); err != nil {
            panic(err)
        }
    }

    character := CreateNewCharacter(requestData.Data)
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusCreated)
    if err := json.NewEncoder(w).Encode(character); err != nil {
        panic(err)
    }
}


func CharacterUpdate(w http.ResponseWriter, r *http.Request) {
    var requestData UpdateRequest 
    body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
    if err != nil {
        panic(err)
    }
    if err := r.Body.Close(); err != nil {
        panic(err)
    }
    if err := json.Unmarshal(body, &requestData); err != nil {
        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        w.WriteHeader(422) // unprocessable entity
        if err := json.NewEncoder(w).Encode(err); err != nil {
            panic(err)
        }
    }

    character := UpdateCharacter(requestData)
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusCreated)
    if err := json.NewEncoder(w).Encode(character); err != nil {
        panic(err)
    }
}

func CharacterShow(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	identifier := vars["identifier"]

	result, err := FindCharacter(identifier)

	if err != nil {
		fmt.Fprintln(w, "Character not found!")
		return
	}

	resultString,_ := json.Marshal(result)
	fmt.Fprintln(w, "", string(resultString))
}
