package main

import "github.com/asalih/elasticsearch-lim/app"

func main() {
	var timeh = &TimeHandler{}
	go timeh.InitTime()

	InitServer()

}
