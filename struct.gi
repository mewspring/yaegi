package main

type T struct {
	f int
	g int
}

func main() {
	a := T{ 7, 8 }
	//a := T{}
	println(a.f, a.g)
}