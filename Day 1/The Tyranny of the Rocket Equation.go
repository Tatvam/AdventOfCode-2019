package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
)

	
func check(e error) {
    if e != nil {
        panic(e)
    }
}

func calc(collection []string)  int{
	var res int
	res = 0
	for _,s := range collection {
		dat,_ := strconv.Atoi(s)
		res += dat/3 - 2
	}
	return res
}


func main() {
	dat,err := ioutil.ReadFile("./input.txt")
	check(err)

	list := string(dat)
	collection := strings.Split(list, "\n")
	result := calc(collection)
	fmt.Println(result)

}