package main

import "github.com/mj37yhyy/embassy/embctl/cmd/pkg"

func main() {
	fErr := cmd.Execute()
	if fErr != nil {
		panic(fErr)
	}
}
