package metricsCollector

import (
	"bytes"
	"fmt"
	"github.com/shirou/gopsutil/v3/mem"
	"os/exec"
	"strconv"
	"strings"
)

func fetchMemInfo() (*mem.VirtualMemoryStat, error) {
	memInfo, memErr := mem.VirtualMemory()
	if memErr != nil {
		return nil, fmt.Errorf("memory error: %v", memErr)
		// change to log info
	}

	return memInfo, nil
}

func fetchCPUInfo(key string) (string, error) {
	cmd := exec.Command("sysctl", "-n", key)
	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("failed to fetch %s: %v", key, err)
		// change to log info
	}

	return strings.TrimSpace(out.String()), nil
}

func collectedData() (*mem.VirtualMemoryStat, int, int, int) {
	memInfo, memErr := fetchMemInfo()
	if memErr != nil {
		fmt.Println(memErr)
		// change to log info
		return nil, 0, 0, 0
	}

	coreCntStr, err := fetchCPUInfo("machdep.cpu.core_count")
	if err != nil {
		fmt.Println("Error fetching core count:", err)
		// change to log info
		return nil, 0, 0, 0
	}
	coreCnt, _ := strconv.Atoi(coreCntStr)

	// CPU thread count
	threadCntStr, err := fetchCPUInfo("machdep.cpu.thread_count")
	if err != nil {
		fmt.Println("Error fetching thread count:", err)
		return nil, 0, 0, 0
	}
	threadCnt, _ := strconv.Atoi(threadCntStr)

	// physical cores (active)
	actCoresStr, err := fetchCPUInfo("hw.ncpu")
	if err != nil {
		fmt.Println("Error fetching active cores:", err)
		// change to log info
		return nil, 0, 0, 0
	}
	actCores, _ := strconv.Atoi(actCoresStr)

	// TODO: fix the cpu metrics -> fixed

	// returns the whole memInfo struct slice
	return memInfo, coreCnt, threadCnt, actCores
}
