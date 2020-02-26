package main

import (
	"fmt"
	"os"
	"os/exec"
)

//***将文件编译为Linux文件可执行

//filepath: 要编译的文件的路径
func build(filepath string) {
	_ = os.Setenv("CGO_ENABLED", "0")
	_ = os.Setenv("GOARCH", "amd64")
	_ = os.Setenv("GOOS", "linux")

	arg := []string{"build", filepath}
	if err := exec.Command("go", arg...).Run(); err != nil {
		fmt.Println("编译失败:", err)
	} else {
		fmt.Println("编译成功")
	}
}

func main() {
	build(`D:\Go_Code\note-gin\main.go`)
}
