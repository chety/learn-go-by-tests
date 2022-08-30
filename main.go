package main

import "fmt"

const englishGreeting = "Hello,"

func Hello(name string) string {
	if len(name) == 0 {
		name = "World!"
	}
	return fmt.Sprintf("%s %s", englishGreeting, name)
}

func main() {
	fmt.Println(Hello("Mirko"))
}
