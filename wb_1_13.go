package main

func main() {
	var (
		first  int
		second int
	)

	first = 9
	second = 21
	println(first, second)

	first, second = second, first
	println(first, second)

	first += second
	second -= first
	second = -second
	first -= second
	println(first, second)
}
