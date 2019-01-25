package main

import (
	"./cmd"
)

func main() {
	cmd.InitConfig()
	cmd.Execute()
}
