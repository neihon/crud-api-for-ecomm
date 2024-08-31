package main

import (
	"fmt"
	"runtime"
)

func detectOs() {
	systemOs := runtime.GOOS
	fmt.Println(systemOs)
	// maybe wrong output? maybe right?
}

func main() {
	detectOs()

}
