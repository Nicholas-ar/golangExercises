package main

import (
	"fmt"
	"runtime"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
)

func main() {
	fmt.Println("=== System Information ===")

	// OS and Platform info
	hostInfo, _ := host.Info()
	fmt.Printf("OS: %s\n", runtime.GOOS)
	fmt.Printf("Platform: %s\n", hostInfo.Platform)
	fmt.Printf("Platform Version: %s\n", hostInfo.PlatformVersion)
	fmt.Printf("Host Name: %s\n", hostInfo.Hostname)

	// CPU info
	cpuInfo, _ := cpu.Info()
	fmt.Printf("\n=== CPU Information ===\n")
	for _, cpu := range cpuInfo {
		fmt.Printf("Model: %s\n", cpu.ModelName)
		fmt.Printf("Cores: %d\n", cpu.Cores)
		fmt.Printf("MHz: %.2f\n", cpu.Mhz)
	}

	// Memory info
	memInfo, _ := mem.VirtualMemory()
	fmt.Printf("\n=== Memory Information ===\n")
	fmt.Printf("Total: %.2f GB\n", float64(memInfo.Total)/(1024*1024*1024))
	fmt.Printf("Used: %.2f GB\n", float64(memInfo.Used)/(1024*1024*1024))
	fmt.Printf("Free: %.2f GB\n", float64(memInfo.Free)/(1024*1024*1024))
	fmt.Printf("Usage: %.2f%%\n", memInfo.UsedPercent)

	// Disk info
	partitions, _ := disk.Partitions(false)
	fmt.Printf("\n=== Disk Information ===\n")
	for _, partition := range partitions {
		diskUsage, _ := disk.Usage(partition.Mountpoint)
		fmt.Printf("\nMount Point: %s\n", partition.Mountpoint)
		fmt.Printf("Total: %.2f GB\n", float64(diskUsage.Total)/(1024*1024*1024))
		fmt.Printf("Used: %.2f GB\n", float64(diskUsage.Used)/(1024*1024*1024))
		fmt.Printf("Free: %.2f GB\n", float64(diskUsage.Free)/(1024*1024*1024))
		fmt.Printf("Usage: %.2f%%\n", diskUsage.UsedPercent)
	}
}
