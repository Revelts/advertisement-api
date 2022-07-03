package main

import (
	"advertisement-api/Connection"
	"advertisement-api/Routes"
)

func main() {
	Connection.InitializeConnection()
	Routes.HandleRequests()
}
