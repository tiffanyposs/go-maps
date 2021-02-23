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

## Iterating on a map

```go
func main() {
	// declare a map where all the keys are of type string
	// and all of the values are of type strings
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"white": "#ffffff",
	}

	for key, value := range colors {
		fmt.Println("Hex code for", key, "is", value)
	}
}

```

## Differences between Structs & Maps

**Maps** - closely related properties, like colors

* All keys must be the same type
* All values must be the same type
* Keys are indexed, we can iterate over them
* Used to represent a collection of related properties
* Don't need to know all the keys at complile time
* Reference type

**Structs** - collection of properties for a "thing"

* Values can be of different type
* Keys don't support indexing, we CANNOT iterate over them
* You need to know all the fields at time of compile
* Used to represent a "thing" with a lot of different properties
* Value type