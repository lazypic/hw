package main

import (
	_ "bytes"
	"fmt"
	"github.com/pbnjay/memory"
	"github.com/shirou/gopsutil/cpu"
	_ "io"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

/*
func windisk {
	h := syscall.MustLoadDLL("kernel32.dll")
	c := h.MustFindProc("GetDiskFreeSpaceExW")

	var freeBytes int64

	_, _, err := c.Call(uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(wd))),
		uintptr(unsafe.Pointer(&freeBytes)), nil, nil)
}
*/
func disk() {
	var stat syscall.Statfs_t
	wd, _ := os.Getwd()
	syscall.Statfs(wd, &stat)
	// byte to gb
	// divide the digital storage value by 1e+9
	fmt.Printf("disk %d GB\n", (stat.Blocks*uint64(stat.Bsize))/1e9)
}

func getWinGPU() string {
	Info := exec.Command("cmd", "/C", "wmic path win32_VideoController get name")
	Info.SysProcAttr = &syscall.SysProcAttr{}
	History, _ := Info.Output()

	return strings.Replace(string(History), "Name", "", -1)
}

func getPosix() {
	// $ glxinfo | grep "OpenGL renderer string"
	findStr := "OpenGL renderer string"
	results, err := exec.Command("glxinfo").Output()
	if err != nil {
		fmt.Println(err)
	}
	for _, line := range strings.Split(string(results), "\n") {
		if !strings.Contains(line, findStr) {
			continue
		}
		fmt.Println(strings.TrimLeft(line, findStr))
		break
	}
}

/*
	c2 := exec.Command("grep", findStr)
	r, w := io.Pipe()
	c1.Stdout = w
	c2.Stdin = r
	var b2 bytes.Buffer
	c2.Stdout = &b2
	c1.Start()
	c2.Start()
	c1.Wait()
	w.Close()
	c2.Wait()
	io.Copy(os.Stdout, &b2)
	fmt.Println("g", b2.String())
	fmt.Println("f", findStr)
	fmt.Println("r", strings.TrimPrefix(b2.String(), findStr))
	return strings.TrimPrefix(b2.String(), findStr)
*/
func main() {
	// CPU
	is, err := cpu.Info()
	if err != nil {
		fmt.Printf("Error getting CPU info: %v", err)
	}
	for _, i := range is {
		fmt.Println(i.ModelName)
	}
	// memory
	fmt.Printf("Mem: %dG\n", memory.TotalMemory()/1e9)

	// disk
	disk()

	// gpu
	// window
	// mac
	getPosix()
	// linux
}
