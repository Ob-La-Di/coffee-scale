package main

func main() {
	scale, err := NewScale()

	if err != nil {
		panic(err)
	}

	println(scale.getWeight())
}
