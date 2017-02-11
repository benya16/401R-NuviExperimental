package main

import (
	"../Distributor"
	"fmt"
	"../filter"
	"os"
	"time"
)

func main() {

	Dist := new(distributor.Distributor)
	Dist.GetPosts()

	filter := new(filter.Filter)
	filter.InitFilter(os.Args[1])
	hit, miss := 0, 0

	start := time.Now()
	var elapsed time.Duration
	for _,Post := range Dist.Posts {
		elapsed = time.Since(start)
		//fmt.Println(Post.ToString())
		if filter.ContainsDangerWord(Post.RawText) {
			hit++
		} else {
			miss++
		}
		if elapsed > time.Second {
			break
		}
	}


	fmt.Println("Total: ", hit + miss)
	fmt.Println("True: ", hit)
	fmt.Println("False: ", miss)
	fmt.Println("Elapsed: ", elapsed)
}
