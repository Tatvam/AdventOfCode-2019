package main

import (
	"fmt"
)

func check (l int) bool {
	m := make(map[int]int)
	flag := 0
	flag1 := 1
	prev := l%10
	l=l/10
	for l > 0 {
		curr := l%10
		if(prev == curr){
			m[curr] = m[curr] + 1
		}
		if(prev < curr){
			flag1 = 0;
			break;
		}
		prev = curr
		l = l/10

	}
	for _,s := range m {
		if(s == 1){
			flag = 1
		}
	}
	if flag == 1 && flag1 == 1{
		return true
	}
	return false
}

func main() {
	l := 123257
	r := 647015
	ans := 0
	for i := l ; i<=r ; i++ {
		if check(i) {
			ans++;
		}
	}
	fmt.Println(ans)
}