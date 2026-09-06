package main

import "fmt"

func main() {
	names := []string{"Ildar", "Go", "World"}

	for i, name := range names {
		greeting := greet(name)
		fmt.Printf("%d: %s\n", i, greeting)
	}
}

func greet(name string) string {
	return "Hello, " + name + "!"
}
