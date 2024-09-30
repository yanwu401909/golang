package main

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
)

func main() {
	// 获取CPU信息
	count, _ := cpu.Counts(false)
	cpus, _ := cpu.Info()
	fmt.Println("CPU Info:")
	for _, cpu := range cpus {
		fmt.Printf("- CPU: %s\n", cpu.ModelName)
	}
	fmt.Printf("- CPU CORES: %v\n", count)
	cpuPercent, _ := cpu.Percent(time.Second*1, false)
	for _, p := range cpuPercent {
		fmt.Printf("- CPU Percent: %0.1f %% \n", p)
	}
	// 获取内存信息
	memInfo, _ := mem.VirtualMemory()
	fmt.Println("\nMemory Info:")
	fmt.Printf("- Total: %0.1fG, \n- Available: %0.1fG, \n- Used: %0.1fG, \n- UsedPercent: %0.1f %%\n",
		byte2G(memInfo.Total), byte2G(memInfo.Available), byte2G(memInfo.Used), memInfo.UsedPercent)
	//获取磁盘信息
	fmt.Println("\nDisk Info:")
	disks, _ := disk.Partitions(true)
	for _, d := range disks {
		stat, _ := disk.Usage(d.Mountpoint)
		fmt.Printf("- %v(%v) %0.1fG %0.1fG %0.1fG %0.1f %% \n", d.Mountpoint, d.Fstype, byte2G(stat.Total), byte2G(stat.Free), byte2G(stat.Used), stat.UsedPercent)
	}
}

func byte2G(v uint64) float64 {
	return float64(v) / (1024 * 1024 * 1024)
}
