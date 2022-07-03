package main

var vowels = map[rune]bool{
	'a': true,
	'e': true,
	'i': true,
	'o': true,
	'u': true,
}

func GetCount(str string) (count int) {
	for _, char := range str {
		if vowels[char] {
			count++
		}
	}
	return
}

func main() {

}
