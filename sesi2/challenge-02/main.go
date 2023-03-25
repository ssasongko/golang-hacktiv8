package main

import "challenge-02/routers"

func main() {
	var PORT = ":8080"

	routers.StartServer().Run(PORT)
}
