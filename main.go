package main

import (
    "net/http"
    "fmt"
    "log"
    // "encoding/json"
    "github.com/gorilla/mux"
)

func showVersion(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Version 0.0.1")
    fmt.Fprintf(w, "Version 0.0.1")
}

func bruteForce(w http.ResponseWriter, r *http.Request) {
    pathVars := mux.Vars(r)
    
    // type := pathVars["type"]
    
    fmt.Println("Attempting To Brute Force Password...")
    fmt.Println("Network Type: " + pathVars["type"])
    
    networkType := pathVars["type"]
    
    if networkType == "wep" {
        fmt.Fprintf(w, "Cracking Wep Password")
        crackWep()
    }
    
    if networkType == "wpa" {
        fmt.Fprintf(w, "Cracking WPA Password")
        crackWpa()
    }
    
}

func handleRequests() {
    
    myRouter := mux.NewRouter().StrictSlash(true)
    myRouter.HandleFunc("/", showVersion)
    myRouter.HandleFunc("/brute/{type}", bruteForce)
    myRouter.HandleFunc("/localip", getLocalIp)
    
    log.Fatal(http.ListenAndServe(":10000", myRouter))
    
}

func main(){
    fmt.Println("Network Monitor v2.0 Initialized...")
    fmt.Println("-----------------------------------")
    fmt.Println("* Author : Elliot Forbes")
    
    handleRequests()
}