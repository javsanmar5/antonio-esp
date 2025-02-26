package main

import (
	"fmt"
	"os/exec"
)

func Run(file string) {

    out, err := exec.Command("python", file).Output()
    check(err)

    fmt.Println(string(out))
}
