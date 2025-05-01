package main

import (
	"fmt"
	"log"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/net"
)

// Function to get CPU Usage
func getCPUUsage() {
	percent, err := cpu.Percent(0, true)
	if err != nil {
		log.Println("Error fetching CPU usage:", err)
		return
	}
	for idx, p := range percent {
		fmt.Printf("CPU Core %d Usage: %.2f%%\n", idx, p)
	}
}

// Function to get Memory Usage
func getMemoryUsage() {
	v, err := mem.VirtualMemory()
	if err != nil {
		log.Println("Error fetching memory usage:", err)
		return
	}
	fmt.Printf("Memory Usage: %.2f%%, Total: %.2f GB, Free: %.2f GB\n",
		v.UsedPercent, float64(v.Total)/1024/1024/1024, float64(v.Free)/1024/1024/1024)
}

// Function to get Disk Usage
func getDiskUsage() {
	diskStats, err := disk.Usage("/")
	if err != nil {
		log.Println("Error fetching disk usage:", err)
		return
	}
	fmt.Printf("Disk Usage: %.2f%%, Total: %.2f GB, Free: %.2f GB\n",
		diskStats.UsedPercent, float64(diskStats.Total)/1024/1024/1024, float64(diskStats.Free)/1024/1024/1024)
}

// Function to get Network Stats
func getNetworkStats() {
	interfaces, err := net.Interfaces()
	if err != nil {
		log.Println("Error fetching network stats:", err)
		return
	}
	for _, iface := range interfaces {
		fmt.Printf("Network Interface: %s, Address: %s\n", iface.Name, iface.HardwareAddr)
	}
}

func main() {
	// Infinite loop for periodic extraction of system metrics
	for {
		// Get stats
		fmt.Println("\nFetching system metrics...")
		getCPUUsage()
		getMemoryUsage()
		getDiskUsage()
		getNetworkStats()

		// Wait for 5 seconds before fetching again
		time.Sleep(5 * time.Second)
	}
}
