package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

var maxSyscalls = 0

// SYSCALLFILE is the syscall file name
const SYSCALLFILE = "syscalls"

func main() {
	var systemCalls []string
	f, err := os.Open(SYSCALLFILE)
	if err != nil {
		fmt.Println("err 1")
		log.Fatalln(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.Replace(line, " ", "", -1)
		line = strings.Replace(line, "SYS_", "", -1)
		temp := strings.ToLower(strings.Split(line, "=")[0])
		systemCalls = append(systemCalls, temp)
		maxSyscalls++
	}

	counter := make([]int, maxSyscalls)
	var regs syscall.PtraceRegs
	cmd := exec.Command(os.Args[1], os.Args[2:]...)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{Ptrace: true}

	err = cmd.Start()
	if err != nil {
		fmt.Println("start err:", err)
	}

	err = cmd.Wait()
	if err != nil {
		fmt.Println("wait err:", err)
	}

	pid := cmd.Process.Pid
	fmt.Println("Process ID:", pid)

	before := true
	forCount := 0
	for {
		if before {
			err := syscall.PtraceGetRegs(pid, &regs)
			if err != nil {
				break
			}
			if regs.Orig_rax > uint64(maxSyscalls) {
				fmt.Println("Unknown:", regs.Orig_rax)
				return
			}

			counter[regs.Orig_rax]++
			forCount++
		}

		err := syscall.PtraceSyscall(pid, 0)
		if err != nil {
			fmt.Println("PtraceSysCall:", err)
			return
		}

		_, err = syscall.Wait4(pid, nil, 0, nil)
		if err != nil {
			fmt.Println("Wait4:", err)
			return
		}
		before = !before
	}

	for i, x := range counter {
		if x != 0 {
			fmt.Println(systemCalls[i], "->", x)
		}
	}
	fmt.Println("total systemcalls:", forCount)
}
