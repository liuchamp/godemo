package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/mem"
	"os"
)

func main() {

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/info", func(c *gin.Context) {
		pid := os.Getpid()
		v, _ := mem.VirtualMemory()
		//cm,_:=cpu.Info()
		fmt.Printf("进程 PID: %d \n", pid)

		//prc := exec.Command("bash", "-c", "top -n 1 -p "+string(pid))
		//out, err := prc.CombinedOutput()
		//if err != nil {
		//	panic(err)
		//}
		//fmt.Println(string(out))
		c.JSON(200, gin.H{
			"Total":       v.Total,
			"Free":        v.Free,
			"UsedPercent": v.UsedPercent,
		})
	})
	r.Run(":18880")
}
