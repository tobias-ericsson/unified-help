package helpers

import (
	"fmt"
	"os/exec"
)

func Exec(cmd string, args []string) error {
	path, err := exec.LookPath(cmd)
	if err != nil {
		fmt.Println(err)
		return err
	}

	out, err := exec.Command(cmd, args[1]).Output()
	if err != nil {
		//fmt.Println(err)
		return err
	}
	fmt.Printf("From %s at %s\n", cmd, path)
	fmt.Printf("%s\n", out)
	return nil
}
