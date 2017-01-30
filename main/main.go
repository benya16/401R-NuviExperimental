package main

import (
	"../Distributor"
	"fmt"
)

func main() {

	Dist := new(distributor.Distributor)
	Dist.GetPosts()

	for _,Post := range Dist.Posts {
		fmt.Println(Post.ToString())
	}

}
