package main

import "fmt"

func main() {
	// declare a map where all the keys are of type string
	// and all of the values are of type strings
	// colors := map[string]string{
	// 	"red":   "#ff0000",
	// 	"green": "#00ff00",
	// }

	// init 0 value map
	// var colors = map[string]string

	// init 0 value map
	colors := make(map[string]string)

	colors["white"] = "#ffffff"
	colors["green"] = "#00ff00"

	green := colors["green"]
	// cannot do colors.green with maps

	delete(colors, "white")

	fmt.Println(colors)
}

// $ go run main.go
