package main

import (
	"../Distributor"
	"../pgdatabase"
	"fmt"
	"../filter"
	"os"
	//"time"
	"time"
	"encoding/json"
)

func main() {

	Dist := new(Distributor.Distributor)
	Dist.GetPosts()

	db := pgdatabase.NewDAO()

	filter := new(filter.Filter)
	filter.InitFilter(os.Args[1])
	hit, miss := 0, 0

	start := time.Now()
	var elapsed time.Duration
	for _,Post := range Dist.Posts {
		data, _ := json.Marshal(Post)
		db.AddPost(data)
		//fmt.Println(Post.ToString())
		//if filter.ContainsDangerWord(Post.RawText) {
		//	data, _ := json.Marshal(Post)
		//	db.AddPost(data)
		//	hit++
		//} else {
		//	miss++
		//}
		//if elapsed > time.Second {
		//	break
		//}
	}

	elapsed = time.Since(start)
	fmt.Println("Total: ", hit + miss)
	fmt.Println("True: ", hit)
	fmt.Println("False: ", miss)
	fmt.Println("Elapsed: ", elapsed)
}
