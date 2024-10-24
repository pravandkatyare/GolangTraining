package main

import (
	"fmt"
	"os/exec"
)

// func main() {
// 	cmd := exec.Command("ls", "-l")
// 	out, err := cmd.Output()
// 	if err != nil {
// 		fmt.Println("Error : ", err)
// 	}

// 	fmt.Println("Output: ", out)
// }

// func main() {
// 	cmd := exec.Command("git", "status")
// 	out, err := cmd.CombinedOutput()
// 	if err != nil {
// 		fmt.Println("Error : ", err)
// 	}

// 	fmt.Println("Output: ", string(out))
// }

// func main() {
// 	fmt.Println("Started execution")
// 	cmd := exec.Command("sleep", "5")

// 	err := cmd.Start()
// 	if err != nil {
// 		fmt.Println("Error: ", err)
// 		return
// 	}
// 	fmt.Println("cmd.Process.Pid: ", cmd.Process.Pid)
// 	err = cmd.Wait()
// 	if err != nil {
// 		fmt.Println("Error: ", err)
// 	}
// }

func main() {
	cmd := exec.Command("git", "status")
	err := cmd.Run()
	if exitErr, ok := err.(*exec.ExitError); ok {
		fmt.Println("Command exited with status: ", exitErr.Error())
	}
}
