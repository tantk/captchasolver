package main

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	ctl:= ctrl{}
	ctl.StartServer()
}
