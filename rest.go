package main

import (
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
)

func showVersion(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Version " + Version)
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

func serveStatic(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, r.URL.Path[1:])
}