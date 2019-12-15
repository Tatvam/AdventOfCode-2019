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
func cac(d []int, op int,p int) int {
	if p == 1 {
		return op
	}
	return d[op]

}
func calc(collection []string) {

	input := 5
	d := []int{}
	for _,s := range collection {
		dat,_ := strconv.Atoi(s)
		d = append(d,dat)
	}
	i:=0
	for {
		opcode,p1,p2 := opr(d[i])
		if opcode == 99 {
			break
		}
		if opcode == 1 {
			op1,op2,op3 := d[i+1],d[i+2],d[i+3]
			if op3 == 677 {
				fmt.Println(cac(d,op1,p1) + cac(d,op2,p2))
			}
			d[op3] = cac(d,op1,p1) + cac(d,op2,p2) 
			i+=4
		} else if opcode == 2 {
			op1,op2,op3 := d[i+1],d[i+2],d[i+3]
			if op3 == 677 {
				fmt.Println(cac(d,op1,p1) * cac(d,op2,p2))
			}
			d[op3] = cac(d,op1,p1) * cac(d,op2,p2) 
			i+=4
		} else if opcode == 3 {
			d[d[i+1]] = input
			i+=2
		} else if opcode == 4 {
			op1 := d[i+1]
			fmt.Println(cac(d,op1,p1))
			i+=2
		} else if opcode == 5 {
			op1,op2 := d[i+1],d[i+2]
			if cac(d,op1,p1) != 0 {
				i = cac(d,op2,p2)
				
			} else {
				i+=3
			}
		}  else if opcode == 6 {
			op1,op2 := d[i+1],d[i+2]
			if cac(d,op1,p1) == 0 {
				i = cac(d,op2,p2)
			} else {
				i+=3
			}
		}  else if opcode == 7 {
			op1,op2,op3 := d[i+1],d[i+2],d[i+3]
			if cac(d,op1,p1) < cac(d,op2,p2) {
				d[op3] = 1
			}else {
				d[op3] = 0
			}
			i+=4

		}  else if opcode == 8 {
			op1,op2,op3 := d[i+1],d[i+2],d[i+3]

			if cac(d,op1,p1) == cac(d,op2,p2) {
				d[op3] = 1
			}else {
				d[op3] = 0
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