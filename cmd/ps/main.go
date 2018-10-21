package main

import "payments-simple/internal/app/ps"

func main() {
	ps.LoadConfig()
	ps.RunServer()
}
