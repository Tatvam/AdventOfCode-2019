package main

import (
	"fmt"
	"io/ioutil"
	"math"
)

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func main() {
	dat,_  := ioutil.ReadFile("./data.txt")
	list := string(dat)
	collection := []string{}
	for i := 0; i < len(list); i +=150 {
		batch := list[i:min(i+150, len(list))]
		collection = append(collection,batch)	
	}
	//layers := [][]strings{}
	minZero := math.MaxInt64
	oneTotal := 0
	twoTotal := 0
	for _,layer := range collection {
		zeroNum := 0
		oneNum := 0
		twoNum := 0
		for _,val := range layer {
			if val == '0' {
				zeroNum++
			} else if val == '1' {
				oneNum++
			} else {
				twoNum++
			}
		}
		if(minZero > zeroNum) {
			minZero,oneTotal,twoTotal = zeroNum,oneNum,twoNum
		}
	}
	fmt.Println(oneTotal*twoTotal)
}