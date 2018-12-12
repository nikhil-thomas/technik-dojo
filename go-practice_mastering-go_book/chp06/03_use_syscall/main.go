package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {
	pid, r2, errNo := syscall.Syscall(39, 0, 0, 0)

	fmt.Println("pid: ", pid)
	fmt.Println("r2: ", r2)
	fmt.Println("errNo: ", errNo)

	uid, r2, errNo := syscall.Syscall(24, 0, 0, 0)

	fmt.Println("uid: ", uid)
	fmt.Println("r2: ", r2)
	fmt.Println("errNo: ", errNo)

	fmt.Print("\n\n")

	message := []byte("HELLO\n")
	fd := 1
	syscall.Write(fd, message)
	fmt.Print(":::::\n")

	fmt.Println("Using syscall.Exec()")
	command := "/bin/ls"
	env := os.Environ()
	syscall.Exec(command, []string{"ls", "-a", "-x"}, env)
}
