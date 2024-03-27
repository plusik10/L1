package main

func Way1(a int, b int) (int, int) {
	a = a + b
	b = a - b
	a = a - b
	return a, b
}

func Way2(a int, b int) (int, int) {
	a, b = b, a
	return a, b
}

func Way3(a int, b int) (int, int) {
	a = a * b
	b = a / b
	a = a / b
	return a, b
}

func Way4(a int, b int) (int, int) {
	a = a &^ b
	b = b &^ a
	a = a &^ b
	return a, b
}

func Way5(a int, b int) (int, int) {
	a = a - b
	b = a + b
	a = -a + b
	return a, b
}
