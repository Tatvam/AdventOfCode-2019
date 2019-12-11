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

	collection[1] = "12"
	collection[2] = "2"
	d := []int{}
	for _,s := range collection {
		dat,_ := strconv.Atoi(s)
		d = append(d,dat)
	}
	fmt.Println(d)

	for i:=0; i<len(d); i=i+4 {
		opcode := d[i]
		if opcode == 99 {
			break;
		}

		if opcode == 1 {
			d[d[i+3]] = d[d[i+1]] + d[d[i+2]] 
		} else if opcode == 2 {
			d[d[i+3]] = d[d[i+1]] * d[d[i+2]]  
		}

	}
	return d[0]
}


func main() {
	dat,err := ioutil.ReadFile("./data.txt")
	check(err)

	list := string(dat)
	collection := strings.Split(list, ",")
	fmt.Println(calc(collection))

}