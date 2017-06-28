package utils

import (
	"fmt"

	"github.com/shirou/gopsutil/mem"
)

func GetMemInfo() (*mem.VirtualMemoryStat)  {

	v, _ := mem.VirtualMemory()
	fmt.Println(v)
	return v
}