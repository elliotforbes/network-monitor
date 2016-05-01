package main

import (
    "fmt"
    "os/exec"
    "log"
    "net"
    "net/http"
    "encoding/json"
)

func getLocalIp(w http.ResponseWriter, r *http.Request){
    
    ifaces, err := net.Interfaces()
    if err != nil {
        log.Fatal(err)
    }
      
    for _, i := range ifaces {
        addrs, err := i.Addrs()
    
        if err != nil {
            log.Fatal(err)
        }
                
        for _, addr := range addrs {
            var ip net.IP
                        
            switch v := addr.(type) {
            case *net.IPNet:
                    ip = v.IP
            case *net.IPAddr:
                    ip = v.IP
            }
            // process IP address
            if ip != nil {
                fmt.Println(ip.String())
                fmt.Println(ip.IsMulticast())
            }
        }
    }
    
    fmt.Fprintf(w, "Local IP: ")
}

func crackWep(){
    fmt.Println("Cracking WEP Password...")
    cmd := exec.Command("pwd")
    stdout, err := cmd.StdoutPipe()
    
    if err != nil {
        log.Fatal(err)
    }
    if err := cmd.Start(); err != nil {
        log.Fatal(err)
    }
    
    fmt.Println(json.NewDecoder(stdout))
   
}

func crackWpa(){
    fmt.Println("Cracking WPA Password...")
}

