package main

import "fmt"

func main() {
	// declare a map where all the keys are of type string
	// and all of the values are of type strings
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"white": "#ffffff",
	}

	printMap(colors)
}

func printMap(c map[string]string) {
	for key, value := range c {
		fmt.Println("Hex code for", key, "is", value)
	}
}

// $ go run main.go
