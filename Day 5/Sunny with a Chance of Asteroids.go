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

func opr(a int) (int,int,int) {
	p,r,op := 0,0,0
	op = a%100
	a = a/100
	p = a%10
	a = a/10
	r = a%10
	return op,p,r
}
func cac(d []int, p int, i int) int {
	if p == 1 {
		return d[i]
	}
	return d[d[i]]

}
func calc(collection []string) {

	input := 5
	d := []int{}
	for _,s := range collection {
		dat,_ := strconv.Atoi(s)
		d = append(d,dat)
	}
//	fmt.Println(d)
	i:=0
	for {
		opcode,p1,p2 := opr(d[i])
		fmt.Println(opcode,p1,p2)
		if opcode == 99 {
			break;
		}
		if opcode == 1 {
			d[d[i+3]] = cac(d,p1,i+1) + cac(d,p2,i+2) 
			i+=4
		} else if opcode == 2 {
			d[d[i+3]] = cac(d,p1,i+1) * cac(d,p2,i+2)  
			i+=4
		} else if opcode == 3 {
			d[d[i+1]] = input
			i+=2
		} else if opcode == 4 {
			fmt.Println(d[d[i+1]])
			i+=2
		} else if opcode == 5 {
			if cac(d,p1,i+1) > 0 {
				i = cac(d,p2,i+2)
			} else {
				i+=3
			}
		}  else if opcode == 6 {
			if cac(d,p1,i+1)  == 0 {
				i = cac(d,p2,i+2)
			} else {
				i+=3
			}
		}  else if opcode == 7 {
			if cac(d,p1,i+1) < cac(d,p2,i+2) {
				d[d[i+3]] = 1
			}else {
				d[d[i+3]] = 0
			}
			i+=4

		}  else if opcode == 8 {
			if cac(d,p1,i+1) == cac(d,p2,i+2) {
				d[d[i+3]] = 1
			}else {
				d[d[i+3]] = 0
			}
			i+=4
		}

	}
}


func main() {
	dat,err := ioutil.ReadFile("./data.txt")
	check(err)

	list := string(dat)
	collection := strings.Split(list, ",")
	calc(collection)

}