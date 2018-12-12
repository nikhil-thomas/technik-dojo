package main

//#include <stdio.h>
//void callC() {
//    printf("Calling C Code!\n");
//}
import "C"
import "fmt"

func main() {
	fmt.Println("Go statement")
	C.callC()
	fmt.Println("Go statement")
}
