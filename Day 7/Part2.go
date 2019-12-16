package main

import (
	"fmt"
	"sort"
	"math"
	"strings"
	"strconv"
	"io/ioutil"
)

var ans int = math.MinInt64

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

func NextPermutation(x sort.Interface) bool {
	n := x.Len() - 1
	if n < 1 {
		return false
	}
	j := n - 1
	for ; !x.Less(j, j+1); j-- {
		if j == 0 {
			return false
		}
	}
	l := n
	for !x.Less(j, l) {
		l--
	}
	x.Swap(j, l)
	for k, l := j+1, n; k < l; {
		x.Swap(k, l)
		k++
		l--
	}
	return true
}

func Max(x, y int) int {
    if x < y {
        return y
    }
    return x
}

func Amplifiers (program[] int, phase []int) int{

	ea := make(chan int, 1)
	ab := make(chan int)
	bc := make(chan int)
	cd := make(chan int)
	de := make(chan int)

	halt := make(chan bool)

	go calc(program, phase, ea, ab, halt)
	go calc(program, phase, ab, bc, halt)
	go calc(program, phase, bc, cd, halt)
	go calc(program, phase, cd, de, halt)
	go calc(program, phase, de, ea, halt)

	ea<-phase[0]
	ab<-phase[1]
	bc<-phase[2]
	cd<-phase[3]
	de<-phase[4]

	ea<-0

	for i := 0; i < 5; i++ {
		<-halt
	}

	return <-ea


}

func calc(program []int, x []int, input <-chan int, output chan<- int, halt chan<- bool) {
	d := make([]int, len(program))
	copy(d, program)
	i:=0
	for {
		opcode,p1,p2 := opr(d[i])
		if opcode == 99 {
			halt <- true
			break
		}
		if opcode == 1 {
			op1,op2,op3 := d[i+1],d[i+2],d[i+3]
			d[op3] = cac(d,op1,p1) + cac(d,op2,p2) 
			i+=4
		} else if opcode == 2 {
			op1,op2,op3 := d[i+1],d[i+2],d[i+3]
			d[op3] = cac(d,op1,p1) * cac(d,op2,p2) 
			i+=4
		} else if opcode == 3 {
			d[d[i+1]] =  <-input
			i+=2
		} else if opcode == 4 {
			op1 := d[i+1]
			output <- cac(d,op1,p1) 
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
	//fmt.Println(ans)
}
func main() {
	x := []int{5,6,7,8,9}
	dat,_ := ioutil.ReadFile("./data.txt")

	list := string(dat)
	collection := strings.Split(list, ",")
	d := []int{}
	for _,s := range collection {
		dat,_ := strconv.Atoi(s)
		d = append(d,dat)
	}
	for i := 1; NextPermutation(sort.IntSlice(x)); i++ {
		ans = Max(ans,Amplifiers(d, x))
	}
	fmt.Println(ans)

}