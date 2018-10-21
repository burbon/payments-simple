package main

import "payments-simple/internal/app/ps"

func main() {
	ps.LoadConfig()

	session := ps.CreateSession()
	defer session.Close()

	ps.RunServer()
}
