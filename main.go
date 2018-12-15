package main

import (
	"fmt"
	"github.com/pbnjay/memory"
	"github.com/shirou/gopsutil/cpu"
	"log"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

/* get windows Disk Size
func windisk {
	h := syscall.MustLoadDLL("kernel32.dll")
	c := h.MustFindProc("GetDiskFreeSpaceExW")

	lpFreeBytesAvailable := int64(0)
    lpTotalNumberOfBytes := int64(0)
    lpTotalNumberOfFreeBytes := int64(0)
    r2, _, err := c.Call(uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr("C:"))),
        uintptr(unsafe.Pointer(&lpFreeBytesAvailable)),
        uintptr(unsafe.Pointer(&lpTotalNumberOfBytes)),
        uintptr(unsafe.Pointer(&lpTotalNumberOfFreeBytes)))
}
*/

func getDiskSize() {
	var stat syscall.Statfs_t
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	syscall.Statfs(wd, &stat)
	fmt.Printf("Disk:%dGB\n", (stat.Blocks*uint64(stat.Bsize))/1024/1024/1024.0)
}

// 윈도우즈에서 그래픽카드를 가지고오는 함수
func getWinGPU() string {
	Info := exec.Command("cmd", "/C", "wmic path win32_VideoController get name")
	Info.SysProcAttr = &syscall.SysProcAttr{}
	History, err := Info.Output()
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(string(History), "Name", "", -1)
}

// mac, linux에서만 정보를 가지고 올 수 있다.
func getGraphicCard() {
	// $ glxinfo | grep "OpenGL renderer string"
	findStr := "OpenGL renderer string:"
	results, err := exec.Command("glxinfo").Output()
	if err != nil {
		log.Fatal(err)
	}
	for _, line := range strings.Split(string(results), "\n") {
		if !strings.Contains(line, findStr) {
			continue
		}
		fmt.Println(strings.TrimLeft(line, findStr))
		break
	}
}

func getCpu() {
	is, err := cpu.Info()
	if err != nil {
		log.Fatal(err)
	}
	for _, i := range is {
		fmt.Println(i.ModelName)
	}
}

func getMemSize() {
	fmt.Printf("Mem:%dG\n", memory.TotalMemory()/1024/1024/1024.0)
}

func main() {
	getCpu()
	getGraphicCard()
	getMemSize()
	getDiskSize()
}
