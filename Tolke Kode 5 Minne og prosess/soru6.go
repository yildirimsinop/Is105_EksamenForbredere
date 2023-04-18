package main

import (
	"fmt"
	"os"
)

func main() {
	var attr = os.ProcAttr{}
	process, err := os.StartProcess("/bin/sleep", []string{"/bin/sleep", "5"}, &attr)
	if err == nil {
		fmt.Println("process started with pid", process.Pid)
		state, err := process.Wait()
		if err != nil {
			fmt.Println(err.Error())
		}

		fmt.Println("process exited with state", state)
	} else {
		fmt.Println(err.Error())
	}
}
