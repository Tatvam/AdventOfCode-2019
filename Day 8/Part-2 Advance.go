package main

import (
	"fmt"
	"github.com/labstack/gommon/color"
	"io/ioutil"
	"strconv"
	tm "github.com/buger/goterm"
    "time"
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
	tm.Clear() 
	currLayer := []int{}
	for _,s := range collection[0] {
		d,_ := strconv.Atoi(string(s))
		currLayer = append(currLayer, d)
	}

	i,j := 0,0
	for i < 6 {
		j = 0
		for j < 25 {
			tm.MoveCursor(j + 1, i + 1 )
			if currLayer[25*i+j] == 2 {
				tm.Printf(color.BlackBg(" "))
			} else if currLayer[25*i+j] == 1 {
				tm.Printf(color.RedBg(" "))
			} else {
				tm.Printf(color.WhiteBg(" "))
			}
			j++
			time.Sleep(time.Second/100)
			tm.Flush()
		}
		i++
		fmt.Println("\n")
	}

	for _,layer := range collection {
		i,j := 0,0
		for i < 6 {
			j = 0
			for j < 25 {
				tm.MoveCursor(j + 1, i + 1 )
				if layer[25*i+j] == '2' {
				} else if layer[25*i+j] == '1' {
					tm.Printf(color.RedBg(" "))
				} else {
					tm.Printf(color.WhiteBg(" "))
				}
				j++
				time.Sleep(time.Second/1000)
				tm.Flush()
			}
			i++
		}
	}

}
