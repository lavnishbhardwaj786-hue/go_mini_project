package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

type data struct {
	word      string
	frequency int
}

func getfilename() string {
	if len(os.Args) < 2 {
		fmt.Println("plese enter the file name")
		os.Exit(1)
	}
	name := os.Args[1]
	return name
}
func fileread() string {
	value := getfilename()
	data, err := os.ReadFile(value)
	if err != nil {
		fmt.Println("not able to read file or find file", err)
	}
	return string(data)
}
func datatostruct() []data {
	information := fileread()
	datatobesortedmap := make(map[string]int)
	words := strings.Fields(information)
	for _, word := range words {
		datatobesortedmap[word] += 1
	}
	datatobesorted := make([]data, 0, 10)
	for key, value := range datatobesortedmap {
		datatobesorted = append(datatobesorted, data{
			word:      key,
			frequency: value,
		})
	}
	slices.SortFunc(datatobesorted, func(a, b data) int {
		if a.frequency != b.frequency {
			return b.frequency - a.frequency
		}

		return strings.Compare(a.word, b.word)
	})
	return datatobesorted
}
func main() {
	mydata := datatostruct()
	limit := min(10, len(mydata))
	for i := 0; i < limit; i++ {
		fmt.Println("index", i, "value", mydata[i])
	}
}
