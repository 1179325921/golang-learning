package main

import (
	"bufio"
	"fmt"
	"os/exec"
)

func main() {
	// os/exec API执行命令并建立管道
	cmd0 := exec.Command("echo", "-n", "My first command comes from golang")
	//if err := cmd0.Start(); err != nil {
	//	fmt.Printf("Erorr:The Command N0 can not be startup:%s\n", err)
	//	return
	//}
	stdout0, err := cmd0.StdoutPipe() // StdoutPipe方法返回一个输出管道
	if err != nil {
		fmt.Printf("Error: Could not obtain the stdout pipe for command N0:%s\n", err)
		return
	}
	outputBuf0 := bufio.NewReader(stdout0)
	output0, _, err := outputBuf0.ReadLine()
	if err != nil {
		fmt.Printf("Error:could not read data from pipe :%s\n", err)
		return
	}
	fmt.Printf("%s\n", string(output0))
}
