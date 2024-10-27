package main

import (
	"fmt"

	rtfServer "rtfServer/server"
)

func main() {
	rtfServer.LaunchServer()
	fmt.Println("server running")
}
