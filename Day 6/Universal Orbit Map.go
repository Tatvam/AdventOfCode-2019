package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var ans int

func createGraph(m map[string][]string, collection []string,visited map[string]int) {
	for _,s := range collection {
		nodes := strings.Split(s,")")
		m[nodes[0]] = append(m[nodes[0]],nodes[1])
		//m[nodes[1]] = append(m[nodes[1]],nodes[0])
		visited[nodes[0]] = 0
		visited[nodes[1]] = 0
	}
}

func dfs (m map[string][]string, visited map[string]int, start string, k int) {
	visited[start] = 1

	for _,s := range m[start] {
		if visited[s] == 0 {
			dfs(m, visited, s,k+1)
		}
	}
	ans+=k
}

func calc(collection []string) {
	m := make(map[string][]string)
	visited := make(map[string]int)
	createGraph(m,collection,visited)

	for _,s := range m {
		for _,t := range s {
			fmt.Println(t)
		}
	}
	k := 0
	start := "COM"
	dfs(m,visited,start,k)
}

func main() {
	dat,_ := ioutil.ReadFile("./data.txt")
	list := string(dat)
	ans = 0
	collection := strings.Split(list, "\n")
	calc(collection)
	fmt.Println(ans)
}