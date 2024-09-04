package main

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/mem"
	"os/exec"
	"strings"
)

func fetchMemInfo() (*mem.VirtualMemoryStat, error) {
	memInfo, memErr := mem.VirtualMemory()
	if memErr != nil {
		return nil, fmt.Errorf("memory error: %v", memErr)
	}

	return memInfo, nil
}

func fetchCPUInfo() (string, error) {
	out, err := exec.Command("sysctl", "-n", "machdep.cpu.brand_string").Output()
	if err != nil {
		fmt.Println("cpu error", err)
		return "", err
	}
	cpuInfo := strings.TrimSpace(string(out))

	return cpuInfo, nil
}

func collectedData() (uint64, uint64, uint64, float64, uint64, uint64, uint64, uint64, uint64, uint64, uint64, uint64, uint64, string) {
	cpuInfo, cpuErr := fetchCPUInfo()
	if cpuErr != nil {
		fmt.Println(cpuErr)
		return 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, ""
	}

	memInfo, memErr := fetchMemInfo()
	if memErr != nil {
		fmt.Println(memErr)
		return 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, ""
	}

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

	// for cpu
	cpuChipInfo := cpuInfo

	return memTotal, memFree, memAvailable, usedPercent, memBuffers, memActive, memCached, memDirty, memInactive, memHighTotal, memLowTotal, memHighFree, memLowFree, cpuChipInfo
}
