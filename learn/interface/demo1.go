package main

import "fmt"

type VowelsFinder interface {
	FindVowels() []rune
}

type MyString string

// MyString implements VowelsFinder
// ms receiver
func (ms MyString) FindVowels() []rune {
	var vowels []rune
	for _, r := range ms {
		if isVowel(r) {
			vowels = append(vowels, r)
		}
	}
	return vowels
}

func isVowel(word rune) bool {
	if word == 'a' || word == 'e' || word == 'i' || word == 'o' || word == 'u' {
		return true
	}
	return false
}

func main() {
	Name := MyString("qwerpvnavopqtqpeovba")
	var v VowelsFinder
	v = Name // valid only if MyString implements VowelsFinder
	fmt.Printf("Vowels are %c ", v.FindVowels())
}
