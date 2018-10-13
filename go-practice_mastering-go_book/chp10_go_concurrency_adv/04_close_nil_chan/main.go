package main

func main() {
	var c chan string
	close(c)

	// output
	// panic: close of nil channel
}
