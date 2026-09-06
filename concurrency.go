package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"
)

// we are working as a client in it as we have to open those url concurrently
func requesting() (resp *http.Response, err error) {
	// in this way we use much time as we go for one request wait and then other and wait so we will use gruetine from now on for these things
	urls := []string{
		"https://google.com",
		"https://github.com",
		"https://example.com",
		"https://demoqa.com/",
		"https://go.dev",
		"https://golang.org",
		"https://www.microsoft.com",
		"https://www.apple.com",
		"https://www.amazon.com",
		"https://www.wikipedia.org",
		"https://www.reddit.com",
		"https://www.cloudflare.com",
		"https://www.mozilla.org",
		"https://www.python.org",
		"https://nodejs.org",
		"https://www.npmjs.com",
		"https://stackoverflow.com",
		"https://www.linkedin.com",
		"https://www.oracle.com",
		"https://www.ibm.com",
		"https://www.netflix.com",
		"https://this-domain-does-not-exist-123456789.com",
		"https://another-invalid-domain-987654321.com",
		"https://example.com/nonexistent-page-123456",
		"https://httpstat.us/500",
	}
	start := time.Now()
	for i := 0; i < len(urls); i++ {
		start1 := time.Now()
		resp, err = http.Get(urls[i])
		urlprocesstime := time.Since(start1)
		if err != nil {
			fmt.Println("can give requests"+"the url is ", i, "the error is ", err)
			continue
		}
		fmt.Println(resp.StatusCode, "which url", urls[i], urlprocesstime)
		resp.Body.Close()
	}
	// do something that takes time
	elapsed := time.Since(start)
	fmt.Println(elapsed)
	return
}
func checkURL(url string, ctx context.Context) {
	start := time.Now()
	// we can't use go as go http.get we will loos resp,err so we create a funciton
	// which is gruotine it self
	//

	//resp, err := http.Get(url)

	//
	// we will use request with context insted of normal url
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil) // it use to create request

	if err != nil {
		fmt.Println(url, "failed to create request:", err)
		return
	}

	resp, err := http.DefaultClient.Do(req) // send the request
	elapsed := time.Since(start)

	if err != nil {
		fmt.Println(url, "failed:", err)
		return
	}
	defer resp.Body.Close()
	fmt.Println(
		url,
		resp.StatusCode,
		"time:",
		elapsed,
	)
}
func requestinggrutine() {
	// ye url launcher ha to ye kuch na bhi return kre to bhi koi faraq nahi padega
	urls := []string{
		"https://google.com",
		"https://github.com",
		"https://example.com",
		"https://demoqa.com/",
		"https://go.dev",
		"https://golang.org",
		"https://www.microsoft.com",
		"https://www.apple.com",
		"https://www.amazon.com",
		"https://www.wikipedia.org",
		"https://www.reddit.com",
		"https://www.cloudflare.com",
		"https://www.mozilla.org",
		"https://www.python.org",
		"https://nodejs.org",
		"https://www.npmjs.com",
		"https://stackoverflow.com",
		"https://www.linkedin.com",
		"https://www.oracle.com",
		"https://www.ibm.com",
		"https://www.netflix.com",
		"https://this-domain-does-not-exist-123456789.com",
		"https://another-invalid-domain-987654321.com",
		"https://example.com/nonexistent-page-123456",
		"https://httpstat.us/500",
	}
	start := time.Now()
	var wg sync.WaitGroup
	// wg.done() funciton ke under ka counter to wo har gruotine me jaye or us ke kam hone ka us ko pata ho
	// so we used it like that and we can used it in a way like use the defer in checkurl as wg.done()
	// context.background is used when you are creating top level operation yourslef
	// in normal you can drive context from incoming request ctx:=request.context() because request life cycle control operation
	for i := 0; i < len(urls); i++ {
		// to launch n gruoting we can tell this eralier as well wg.add(len(url))
		// now we don't required to do anything like we.add(1)
		wg.Add(1)
		go func() {
			defer wg.Done()
			ctx, cancle := context.WithTimeout(context.Background(), 3*time.Second)
			defer cancle()
			checkURL(urls[i], ctx)
		}()
		// in previous verison for loop i will be treated seprately for all gruotine
	}
	wg.Wait()
	elapsed := time.Since(start)
	fmt.Println(elapsed)
}
func main() {
	fmt.Println("before using concurrency")
	value, err := requesting()
	fmt.Println(value)
	if err != nil {
		fmt.Println("error is there ding", err)
	}
	fmt.Println("this is after using concurrency")
	requestinggrutine()
}

// if you use http.client it has timeout function in it as well we can use it in here but context is powerful as it's not a time function
