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

func calc(collection []string)  (int,int) {

	j:=0
	k:=0

	for j <= 99 {
		k = 0
		for k <= 99 {
			d := []int{}
			for _,s := range collection {
				dat,_ := strconv.Atoi(s)
				d = append(d,dat)
			}
			
			d[1] = j
			d[2] = k

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
			if d[0] == 19690720 {
				return j,k
			}
			k++
		}
		j++
	}
	return 0,0
}


func main() {
	dat,err := ioutil.ReadFile("./data.txt")
	check(err)

	list := string(dat)
	collection := strings.Split(list, ",")
	fmt.Println(calc(collection))

}