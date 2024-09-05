package main

import (
	"fmt"
	"interforestaces/conifer"
	"interforestaces/deciduous"
	"interforestaces/trees"
)

func main() {
	fmt.Println("Welcome to the forest!")
	fmt.Println("Let's make some trees. We'll watch them become big and mighty!")
	d := deciduous.New("Acer", "saccharum", "sugar maple")
	trees.Lifecycle(d)
	c := conifer.New("Pinus", "elliottii", "slash pine")
	trees.Lifecycle(c)
	fmt.Println("Woah! How fun was that?")
	fmt.Println("Thanks for visiting!")
}
