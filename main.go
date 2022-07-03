package main

func GetCount(str string) (count int) {
	for _, char := range str {
		if char == 'a' {
			count++
		}
	}
	return
}

func main() {

}
