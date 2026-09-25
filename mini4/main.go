package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type apidata struct {
	Id      int    `json:"id"`
	Name    string `json:"Name"`
	Work    string `json:"Work"`
	Salary  int    `json:"Salary"`
	Company string `json:"Company"`
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		log.Println("Request:", r.Method, r.URL.Path)

		next.ServeHTTP(w, r)
	})
}

func (apidata) datarequest(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		data, err := read("data.json")
		if err != nil {
			fmt.Println("not able to read the file ")
			return
		}
		w.Header().Set("Content-Type", "application/json")
		var info apidata
		err = json.Unmarshal(data, &info)
		if err != nil {
			fmt.Println("not able to unmarshall json->go", err)
			return
		}
		some, err := json.Marshal(info)
		if err != nil {
			fmt.Println("not able to marshall go -> json")
			return
		}
		w.Write(some)
	}
	if r.Method == http.MethodPut {
		defer r.Body.Close()

		w.Header().Set("Content-Type", "application/json")

		var change apidata

		err := json.NewDecoder(r.Body).Decode(&change)
		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		some, err := json.Marshal(change)
		if err != nil {
			http.Error(w, "Could not marshal JSON", http.StatusInternalServerError)
			return
		}
		fmt.Println("JSON to write:", string(some))
		fmt.Println("About to write file...")

		err = write(some)

		fmt.Println("WriteFile executed")

		if err != nil {
			fmt.Println("WRITE ERROR:", err)
			http.Error(w, "Could not write file", http.StatusInternalServerError)
			return
		}

		fmt.Println("File written successfully")
		w.Write(some)
	}

}

var mu sync.Mutex

func read(name string) ([]byte, error) {
	defer mu.Unlock()
	mu.Lock()
	data, err := os.ReadFile(name)
	if err != nil {
		if data == nil {
			fmt.Println("file is empty")
		}
		fmt.Println("we can't read file:")
		return data, err
	}
	return data, err
}
func write(data []byte) error {
	defer mu.Unlock()
	mu.Lock()
	err := os.WriteFile("data.json", data, 0644)
	if err != nil {
		fmt.Println("we can't write in the file:")
		return err
	}
	return err
}
func main() {
	api := apidata{}
	mux := http.NewServeMux()
	// mux - ye router ha
	mux.HandleFunc("/api", api.datarequest)
	loggedMux := loggingMiddleware(mux)
	server := &http.Server{
		Addr:    ":8080",
		Handler: loggedMux,
	} // our own created server
	go func() {
		if err := server.ListenAndServe(); err != nil &&
			err != http.ErrServerClosed {
			log.Println(err)
		}
	}() // it is used because to keep out server running and behind it we can perform shut down and some reqest which are in the run can be completed
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	<-sig // it will wait until signal arrives
	ctx, channel := context.WithTimeout(context.Background(), 5*time.Second)
	defer channel()
	if err := server.Shutdown(ctx); err != nil {
		log.Println("Shutdown error:", err)
	} //block the request and let running request finish then shutdown
}