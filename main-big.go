package main

import (
        "fmt"
        "github.com/validatepolicy/service"
        "github.com/validatepolicy/dbclient"
)

var appName = "validatepolicy"

func main() {
        fmt.Printf("Starting %v\n", appName)
        initializeBoltClient()
        service.StartWebServer("6767")
}

func initializeBoltClient() {
        service.DBClient = &dbclient.BoltClient{}
        service.DBClient.OpenBoltDb()
        service.DBClient.Seed()
}
