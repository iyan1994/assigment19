package main

import "fmt"

func ReverseString(text string) string {
	valRune := []rune(text)

	for i, j := 0, len(valRune)-1; i < j; i, j = i+1, j-1 {
		valRune[i], valRune[j] = valRune[j], valRune[i]

	}
	return string(valRune)

}

func main() {

	kalimat := "belajar git"

	fmt.Println(ReverseString(kalimat))
}
