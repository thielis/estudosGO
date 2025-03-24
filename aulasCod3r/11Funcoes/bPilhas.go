package main

import "runtime/debug"

func func4() {
	debug.PrintStack()
}

func func3() {
	func4()
}

func func2() {
	func3()
}

func func1() {
	func2()
}

func main() {
	func1()
}
