package main

import (
	"fmt"
)

func check (l int) bool {
	flag := 0
	prev := l%10
	l=l/10
	for l > 0 {
		curr := l%10
		if(prev == curr){
			flag = 1
		}
		if(prev < curr){
			flag = 0;
			break;
		}
		prev = curr
		l = l/10

	}
	if flag==1 {
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