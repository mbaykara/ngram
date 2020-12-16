package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
	"unicode"

	"github.com/bbalet/stopwords"
)


func main() {
 	if len(os.Args) < 2 {
		fmt.Println("Provide a proper filename")
		os.Exit(1)
	}
	 
	d, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Given filename is not exist: ", os.Args[1])
		panic(err)
	}

	fmt.Println(kgram(string(d),2))
}

func tokenize(data string) map[string]int{

	a := regexp.MustCompile(` `)
	result := a.Split(data,-1)
	words := result
	m := make(map[string]int)
	for _, word := range words {
		m[word]++
	}
	return m
}
/*
 following method is inspired by https://github.com/rafatbiin/gongram/blob/master/ngram.go
*/
func puncAndStopword(data string) string {
	data = strings.TrimSpace(data)
	data = stopwords.CleanString(data,"de",true)
	builder := strings.Builder{}
	builder.Grow(len(data))
	for _, c := range data {
		if !unicode.IsPunct(c) {
			builder.WriteRune(c)
		}
	}
	return builder.String()
}
 
func kgram(data string, k int) ([]string, error) {
	if k > len(data) {
		fmt.Println("The given kgram is bigger than corpus")
	}
	data = puncAndStopword(data)
	s := strings.Split(data," ")
	result := []string{}
	for i := 0 ; i<(len(s)-k+1);i++ {
		ngram := strings.Join(s[i:i+k], " ")
		result = append(result, ngram)
	}
	return result, nil
}
