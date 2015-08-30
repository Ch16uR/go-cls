package main

import (
	_ "runtime"
)

func init() {
	//runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	server()
}
