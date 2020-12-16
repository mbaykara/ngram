package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"

	"github.com/bbalet/stopwords"
)


func main() {
	if len(os.Args) < 2 {
		fmt.Println("Provide a proper filename")
		os.Exit(1)
	}
	d, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("The file name is not available: ", os.Args[1])
		panic(err)
	}

	data :=tokenize(string(d))
	ngram(data,2)

}

func tokenize(data string) map[string]int{
	//Data will be cleaned or not
	cleanData := stopwords.CleanString(string(data),"de",true)
	fmt.Println(cleanData)
	a := regexp.MustCompile(` `)
	result := a.Split(data,-1)
	words := result

	m := make(map[string]int)
	for _, word := range words {
		m[word]++
	}
	return m
		//if _, ok := m[word];ok { m[word]++}
}

