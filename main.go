package main

func main() {
	var timeh = &TimeHandler{}
	go timeh.InitTime()

	InitServer()

}
