package main

func main() {
	d := splitBytes([]byte("abcd:=efghijkl"), ":=")
	for index, line := range d {
		println(index, string(line))
	}
}
