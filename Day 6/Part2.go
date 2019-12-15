package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"container/list"
)

var ans int

func createGraph(m map[string][]string, collection []string,visited map[string]int) {
	for _,s := range collection {
		nodes := strings.Split(s,")")
		m[nodes[0]] = append(m[nodes[0]],nodes[1])
		m[nodes[1]] = append(m[nodes[1]],nodes[0])
		visited[nodes[0]] = 0
		visited[nodes[1]] = 0
	}
}

func bfs (m map[string][]string, visited map[string]int, start string) {
	visited[start] = 1
	level := make(map[string]int)
	queue := list.New()

	queue.PushBack(start)
	level[start] = 0
	for queue.Len() > 0 {
		s := queue.Front()
		queue.Remove(s)
		n := s.Value.(string)
		//fmt.Println(n)
		visited[n] = 1
		for _,v := range m[n] {
			if visited[v] == 0 {
				queue.PushBack(v)
				level[v] = level[n] + 1
			//	fmt.Println(v,level[v])
				if v == "SAN" {
					ans = level[v] - 2
					return
				}
			}
		} 
	}

}

func calc(collection []string) {
	m := make(map[string][]string)
	visited := make(map[string]int)
	createGraph(m,collection,visited)

	// for _,s := range m {
	// 	for _,t := range s {
	// 		fmt.Println(t)
	// 	}
	// }
	start := "YOU"
	bfs(m,visited,start)
}

func main() {
	dat,_ := ioutil.ReadFile("./data.txt")
	list := string(dat)
	ans = 0
	collection := strings.Split(list, "\n")
	calc(collection)
	fmt.Println(ans)
}