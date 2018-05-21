package main

import "fmt"

func main() {
	fmt.Println(Hello("Ndi"))
}

// Hello says hello to the world
func Hello(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}
