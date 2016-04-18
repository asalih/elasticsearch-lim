package main

import "github.com/asalih/elasticsearch-lim/app"

func main() {
	var timeh = app.TimeHandler{}
	go timeh.InitTime()

	app.InitServer()

}
