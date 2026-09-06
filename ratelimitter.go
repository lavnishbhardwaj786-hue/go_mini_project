package main

import (
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"
)

type limitter struct {
	mu      sync.Mutex
	limit   int
	content map[string]int
	time    time.Duration
}

func (l *limitter) timehandle() {
	ticker := time.NewTicker(10 * time.Second)
	for {
		<-ticker.C
		fmt.Println("the 10 second done time to reset ")
		l.reset()
	}
}
func (l *limitter) countcheck(content string) bool {
	count := l.content[content]
	if count < 5 {
		return true
	} else {
		fmt.Println("not more than 5 request at a time")
		return false
	}
}
func (l *limitter) Usercame(limt int, content string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	// 1. Update the limit on the existing struct instance
	l.limit = limt
	// 2. CRITICAL: Initialize the map if it hasn't been initialized yet
	if l.content == nil {
		l.content = make(map[string]int)
	}
	l.content[content] += 1
}
func (l *limitter) reset() {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.content = map[string]int{}
}
func (l *limitter) userHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	check := l.countcheck(id)
	if check == false {
		fmt.Println("too many attempts just lay low right now")
	}
	l.Usercame(5, id)
	fmt.Println(l.content)
}
func main() {
	impo := limitter{}
	mux := http.NewServeMux()
	mux.HandleFunc("GET /users/{id}", impo.userHandler)
	// it is here because we only want one ticker not one per request
	go impo.timehandle()
	err := http.ListenAndServe(":8000", mux)
	if err != nil {
		fmt.Println("this is serious server is not working")
		os.Exit(1)
	} else {
		fmt.Println(" server is running on 8000 visit /users")
	}
}
