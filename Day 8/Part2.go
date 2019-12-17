package main

import (
	"fmt"
	"github.com/labstack/gommon/color"
	"io/ioutil"
	"strconv"
	//"time"
)

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func reverse(s []string) []string {
    for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
        s[i], s[j] = s[j], s[i]
    }
    return s
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
	reverse(collection)

	currLayer := []int{}
	for _,s := range collection[0] {
		d,_ := strconv.Atoi(string(s))
		currLayer = append(currLayer, d)
	}

	// i,j := 0,0
	// for i < 6 {
	// 	j = 0
	// 	for j < 25 {
	// 		if currLayer[25*i+j] == 0 {
	// 			color.Printf(color.BlackBg("  "))
	// 		} else if currLayer[25*i+j] == 1 {
	// 			color.Printf(color.RedBg("  "))
	// 		} else {
	// 			color.Printf(color.BlueBg("  "))
	// 		}
	// 		j++
	// 	}
	// 	i++
	// 	fmt.Println("\n")
	// }

	for _,layer := range collection {
		i,j := 0,0
		for i < 6 {
			j = 0
			for j < 25 {
				if layer[25*i+j] == '2' {
				} else if layer[25*i+j] == '1' {
					currLayer[25*i+j] = 1
				} else {
					currLayer[25*i+j] = 0
				}
				j++
			}
			i++
		}
	}
	i,j := 0,0
	for i < 6 {
		j = 0
		for j < 25 {
			if currLayer[25*i+j] == 0 {
				color.Printf(color.BlackBg("  "))
			} else if currLayer[25*i+j] == 1 {
				color.Printf(color.RedBg("  "))
			} else {
				color.Printf(color.WhiteBg("  "))
			}
			j++
		}
		i++
		fmt.Println("\n")
	}

}
