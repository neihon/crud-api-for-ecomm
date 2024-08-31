package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func getOSVersion() (string, error) {
	osVersion := exec.Command("sw_vers", "-productVersion")
	output, err := osVersion.Output()
	if err != nil {
		return "", err
	}
	fmt.Println(string(output))
	return strings.TrimSpace(string(output)), nil
}

func getSysArch() (string, error) {
	sysArchitecture := exec.Command("uname", "-a")
	output, err := sysArchitecture.Output()
	if err != nil {
		return "", err
	}
	fmt.Println(string(output))
	return strings.TrimSpace(string(output)), nil
}

func readMemData() {}

func main() {
	getOSVersion()
	getSysArch()
}

// get data from RAM dump
// analyse said data
