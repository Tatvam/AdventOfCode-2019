package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
	"math"
)

var d1 int = 0
var d2 int = 0

type Pair struct {
	a int 
	b int 
}
	
func check(e error) {
    if e != nil {
        panic(e)
    }
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Max(x, y int) int {
    if x < y {
        return y
    }
    return x
}

func Min(x, y int) int {
    if x > y {
        return y
    }
    return x
}

func calculateCoord1 (dir string, coord []Pair, i int,lenint int) Pair{
	if dir == "R" {
		return Pair{coord[i].a+lenint,coord[i].b}
	}else if dir == "L" {
		return Pair{coord[i].a-lenint,coord[i].b}
	}else if dir == "U" {
		return Pair{coord[i].a,coord[i].b+lenint}
	}else {
		return Pair{coord[i].a,coord[i].b-lenint}
	}
}

func calculateCoord2 (dir string, currcord Pair, lenint int) Pair{
	if dir == "R" {
		return Pair{currcord.a+lenint,currcord.b}
	}else if dir == "L" {
		return Pair{currcord.a-lenint,currcord.b}
	}else if dir == "U" {
		return Pair{currcord.a,currcord.b+lenint}
	}else {
		return Pair{currcord.a,currcord.b-lenint}
	}
}

func minimize(set1,set2 []string) {
	coord := []Pair{}
	//res := []int{}
	coord = append(coord, Pair{0,0})


	for i,s := range set1 {
		dir := s[0:1]
		le := s[1:]
		lenint,_ := strconv.Atoi(le)
		coord = append(coord,calculateCoord1(dir,coord,i,lenint))	
	}
	currcord := Pair{0,0}
	var ans int
	ans = math.MaxInt64
	for _,s := range set2 {
		dir := s[0:1]
		le := s[1:]
		lenint,_ := strconv.Atoi(le)
		prevcord := currcord
		currcord = calculateCoord2(dir,currcord,lenint)
		d1 = d1+Abs(lenint)
		d2 = 0
		leng := len(coord)
		for j,_ := range coord {
			if j == leng-1 {
				break;
			}
			d2 += Abs(coord[j].a - coord[j+1].a) + Abs(coord[j].b - coord[j+1].b)
			if dir == "R" || dir == "L" {
				if coord[j].a == coord[j+1].a {
					if coord[j].b > coord[j+1].b {
						if currcord.b <= coord[j].b && currcord.b >= coord[j+1].b {
							if (prevcord.a < coord[j].a && currcord.a > coord[j].a) || (prevcord.a > coord[j].a && currcord.a < coord[j].a){
								ans = Min(ans,d1+d2-Abs(coord[j+1].a-currcord.a)-Abs(coord[j+1].b-currcord.b))
							}
						}
					}else{
						if currcord.b <= coord[j+1].b && currcord.b >= coord[j].b {
							if (prevcord.a < coord[j].a && currcord.a > coord[j].a) || (prevcord.a > coord[j].a && currcord.a < coord[j].a){
								ans = Min(ans,d1+d2-Abs(coord[j+1].a-currcord.a)-Abs(coord[j+1].b-currcord.b))
							}
						}
					}
				}
			}else{
				if coord[j].b == coord[j+1].b {
					if coord[j].a > coord[j+1].a {
						if currcord.a <= coord[j].a && currcord.a >= coord[j+1].a {
							if (prevcord.b < coord[j].b && currcord.b > coord[j].b) || (prevcord.b > coord[j].b && currcord.b < coord[j].b){
								ans = Min(ans,d1+d2-Abs(coord[j+1].b-currcord.b)-Abs(coord[j+1].a-currcord.a))
							}
						}
					}else{
						if currcord.a<= coord[j+1].a && currcord.a >= coord[j].a {
							if (prevcord.b < coord[j].b && currcord.b > coord[j].b) || (prevcord.b > coord[j].b && currcord.b < coord[j].b){
								ans = Min(ans,d1+d2-Abs(coord[j+1].b-currcord.b)-Abs(coord[j+1].b-currcord.b))
							}
						}
					}
				}
			}

		}
	}
	fmt.Println(ans)

}


func main() {
	dat,err := ioutil.ReadFile("./data.txt")
	check(err)

	list := string(dat)
	collection := strings.Split(list, "\n")
	set1  := strings.Split(collection[0], ",")
	set2 := strings.Split(collection[1], ",")
	minimize(set1,set2)

}