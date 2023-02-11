package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title  string   `json:"title"`
	Year   int      `json:"year"`
	Price  int      `json:"price"`
	Actors []string `json:"actors"`
}

func main() {

	mo := Movie{"喜剧之王", 200, 10, []string{"xingye", "chengzhipeng"}}

	//json编码过程
	jsonstr, err := json.Marshal(mo)
	if err != nil {
		fmt.Println("json marshal failed")
		return
	}

	fmt.Printf("jsonstr = %s\n", jsonstr)

	//json解码过程

	mymovie := mo
	json.Unmarshal(jsonstr, &mymovie)

	fmt.Printf("%v\n	", mymovie)

}
