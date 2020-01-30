package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/process"
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

		//cm,_:=cpu.Info()
		fmt.Printf("进程 PID: %d \n", pid)

		//prc := exec.Command("bash", "-c", "top -n 1 -p "+string(pid))
		//out, err := prc.CombinedOutput()
		//if err != nil {
		//	panic(err)
		//}
		//fmt.Println(string(out))
		ps, err := process.NewProcess(int32(pid))
		if err != nil {
			fmt.Println(err)
		}
		cpup, err := ps.CPUPercent()
		mif, err := ps.MemoryInfo()
		pofs, err := ps.OpenFiles()
		for _, pof := range pofs {
			fmt.Println(pof.Path, "----->", pof.Fd)
		}
		c.JSON(200, gin.H{
			"cpus":     cpup,
			"men":      mif.VMS,
			"fileinfo": len(pofs),
		})
	})
	r.Run(":18880")
}
