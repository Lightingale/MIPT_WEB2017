package main

import 	(
	//"fmt"
	"unicode"
	"strings"
)

func RemoveEven(arr []int) (result []int) {
	for _, value := range arr {
		if (value % 2 != 0) {
			result = append(result, value)
		}
	}
	return;
}

func PowerGenerator(base int) func() int {
	pow := 1;
	return func() int {
		pow *= base
		return pow
	}
}

func DifferentWordsCount(str string) (count int) {
	str = strings.ToLower(str) + "\n"
	words := make(map[string]bool)
	var word string
	for _, value := range str {
		if unicode.IsLetter(value) {
			word += string(value)
		} else if word != "" {
			if _, ok := words[word]; !ok {
				words[word] = true
				count += 1
			}
			word = ""
		}
	}
	return 
}

//  func main() {
//  	fmt.Println(DifferentWordsCount("Hello, world!HELLO  wOrlD...12"))
//  }
