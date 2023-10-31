package main

import (
	"github.com/dalas98/mekari-test/server"
)

func main() {
	r := server.NewServer()
	r.Run()
}
