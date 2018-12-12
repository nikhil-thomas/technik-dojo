package main

import (
	"fmt"
	"log"
	"os"
	"os/user"
)

func main() {
	fmt.Println("userid:", os.Getuid())
	fmt.Println("effective group id:", os.Getegid())

	var u *user.User
	u, err := user.Current()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Group ids")
	groupIds, err := u.GroupIds()
	if err != nil {
		log.Fatalln(err)
	}
	for _, gid := range groupIds {
		fmt.Print(gid, " ")
	}
	fmt.Print("\n")
}
