package main

import (
    "net/http"
    "fmt"
    "log"
    "github.com/gorilla/mux"
)

const Version = "0.0.1"

func handleRequests() {
    
    myRouter := mux.NewRouter().StrictSlash(true)
    myRouter.HandleFunc("/", showVersion)
    myRouter.HandleFunc("/brute/{type}", bruteForce)
    myRouter.HandleFunc("/localip", getLocalIp)
    myRouter.HandleFunc("/frontend/", serveStatic)
    log.Fatal(http.ListenAndServe(":10000", myRouter))
    
}

func main(){
    fmt.Println("Network Monitor v2.0 Initialized...")
    fmt.Println("-----------------------------------")
    fmt.Println("* Author : Elliot Forbes")
    
    handleRequests()
}