package util

import (
	"log"

	cpu "github.com/shirou/gopsutil/cpu"
)

func GetCPUStatus() float64 {
	var cpuUsage float64
	data, cpuStatusErr := cpu.Times(true)
	var valueRetrieved bool
	for valueRetrieved == false {

		for i := range data {
			if data[i].CPU == "_Total" {
				cpuUsage = data[i].User
				valueRetrieved = true
				break
			}
		}

		if cpuStatusErr != nil || len(data) < 1 {
			if cpuStatusErr != nil {
				log.Println("Error retrieving CPU stats:", cpuStatusErr)
			}

			data, cpuStatusErr = cpu.Times(true)
			continue
		}
	}

	return cpuUsage
}
