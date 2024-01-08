package main

import (
	"fmt"
	"unicode/utf8"
)

func printBytes(s string) {
	fmt.Printf("Bytes: ")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
}

func printChars(s string) {
	fmt.Printf("Characters: ")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c ", s[i])
	}
}
func charsAndBytePosition(s string) {
	for index, rune := range s {
		fmt.Printf("%c starts at byte %d\n", rune, index)
	}
}
func main() {
	name := "Abhinav Gautam"
	fmt.Println(name)
	printBytes(name)
	fmt.Printf("\n")
	printBytes(name)
	println()
	charsAndBytePosition(name)
	println("String length")
	println()
	fmt.Printf("String: %s\n", name)
	fmt.Printf("Length: %d\n", utf8.RuneCountInString(name))
	fmt.Printf("Number of bytes: %d\n", len(name))

	fmt.Printf("\n")
	word2 := "Pets"
	fmt.Printf("String: %s\n", word2)
	fmt.Printf("Length: %d\n", utf8.RuneCountInString(word2))
	fmt.Printf("Number of bytes: %d\n", len(word2))

}
