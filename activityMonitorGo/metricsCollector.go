package main

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
	}

	return strings.TrimSpace(out.String()), nil
}

func collectedData() (uint64, uint64, uint64, float64, uint64, uint64, uint64, uint64, uint64, uint64, uint64, uint64, uint64, int, int, int) {
	memInfo, memErr := fetchMemInfo()
	if memErr != nil {
		fmt.Println(memErr)
		return 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0
	}

	coreCntStr, err := fetchCPUInfo("machdep.cpu.core_count")
	if err != nil {
		fmt.Println("Error fetching core count:", err)
		return 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0
	}
	coreCnt, _ := strconv.Atoi(coreCntStr)

	// CPU thread count
	threadCntStr, err := fetchCPUInfo("machdep.cpu.thread_count")
	if err != nil {
		fmt.Println("Error fetching thread count:", err)
		return 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0
	}
	threadCnt, _ := strconv.Atoi(threadCntStr)

	// physical cores (active)
	actCoresStr, err := fetchCPUInfo("hw.ncpu")
	if err != nil {
		fmt.Println("Error fetching active cores:", err)
		return 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0
	}
	actCores, _ := strconv.Atoi(actCoresStr)

	// TODO: fix the cpu metrics
	// for mem
	memTotal := memInfo.Total
	memFree := memInfo.Free
	memAvailable := memInfo.Available //totalMem - freeMem
	usedPercent := memInfo.UsedPercent
	memBuffers := memInfo.Buffers
	memActive := memInfo.Active
	memCached := memInfo.Cached
	memDirty := memInfo.Dirty
	memInactive := memInfo.Inactive
	memHighTotal := memInfo.HighTotal
	memLowTotal := memInfo.LowTotal
	memHighFree := memInfo.HighFree
	memLowFree := memInfo.LowFree

	// TOneverDO: use again -> for cpu
	//coreCnt := machdep.cpu.core_count
	//threadCnt := machdep.cpu.thread_count
	//actCores := hw.ncpu

	return memTotal, memFree, memAvailable, usedPercent, memBuffers, memActive, memCached, memDirty, memInactive, memHighTotal, memLowTotal, memHighFree, memLowFree, coreCnt, threadCnt, actCores
}
