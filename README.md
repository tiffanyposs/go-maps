# Go Maps

Maps are like objects.

The differnce with Maps vs Structs is that Maps require you to have the same type for all the keys and all the values respectively.

## How Maps Work

```go

func main() {
	// declare a map where all the keys are of type string
	// and all of the values are of type strings
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
        "white": "#ffffff",
	}

    // must access using the square brackets
    // cannot do colors.green, for example, with maps
	green := colors["green"]
	
    // delete "white" from colors map
	delete(colors, "white")

	fmt.Println(colors)
}

```

Two other ways to init a map using empty map to init:

```go
func main() {
    var colors = map[string]string

    colors["white"] = "#ffffff"
    colors["green"] = "#00ff00"
}

```

or with the built in `make` method (this is new)

```go
func main() {
    colors := make(map[string]string)

    colors["white"] = "#ffffff"
    colors["green"] = "#00ff00"
}

```