package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"os/exec"
)

func DeleteSlotnameHandler(name string) bool {
	cmd := exec.Command("sed", "-i", "s/^slot=.*/"+"slot="+name+"/g", "pg.sh")
	//fmt.Println(cmd, cmd.Stdout)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	e := cmd.Run()
	fmt.Println("执行命令为：", cmd)
	if e != nil {
		fmt.Println("错误信息为：", e)
		return false
	}
	cmd2 := exec.Command("bash", "pg.sh")
	cmd2.Stdout = os.Stdout
	cmd2.Stderr = os.Stderr
	e2 := cmd2.Run()
	fmt.Println("执行命令为：", cmd2)
	if e2 != nil {
		fmt.Println("错误信息为：", e2)
		return false
	}
	return true
}
func StopProcessFunc() {
	pid := os.Getpid()            //获取当前进程的pid
	p, err := os.FindProcess(pid) //根据pid获取当前进程的Process对象
	if err != nil {
		panic(err)
	}
	err = p.Kill() //杀掉此进程
	if err != nil {
		panic(err)
	}
}
func SlotHandler(c *gin.Context) {
	slotname := c.Param("name")
	if slotname == "stop" {
		fmt.Println("进程退出...")
		StopProcessFunc()
	}
	fmt.Println("slotname: ", slotname)

	if !DeleteSlotnameHandler(slotname) {
		fmt.Println("删除插槽失败，请检查插槽名是否存在...")
		c.JSON(400, gin.H{
			"err": "删除插槽失败，请检查插槽是否存在",
		})
		return
	} else {
		c.JSON(200, gin.H{
			"msg": "操作成功",
		})
	}

}
func main() {
	r := gin.Default()
	r.GET("/slot/:name", SlotHandler)
	r.Run(":9990")
}
