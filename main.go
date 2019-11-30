package main

import "/server"

func main() {
	r := server.SetupRouter()
	r.Run(":8080")
}