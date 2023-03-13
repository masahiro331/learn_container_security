package main 

import (
	"unsafe"
	"log"
	"os"
	"fmt"
	"syscall"
	"time"
)

const (
	SYS_FCNTL = syscall.SYS_FCNTL
	SYS_MEMFD_CREATE = 319
	MFD_CLOEXEC       uint = 1
	MFD_ALLOW_SEALING uint = 2

	f_LINUX_SPECIFIC_BASE int = 1024

	F_ADD_SEALS int = (f_LINUX_SPECIFIC_BASE + 9)
	F_GET_SEALS int = (f_LINUX_SPECIFIC_BASE + 10)

	F_SEAL_SEAL   int = 0x0001
	F_SEAL_SHRINK int = 0x0002
	F_SEAL_GROW   int = 0x0004
	F_SEAL_WRITE  int = 0x0008
)

func MemfdCreate(name string, flags uint) (fd uintptr, err error) {
	var _p0 *byte
	_p0, err = syscall.BytePtrFromString(name)
	if err != nil {
		return
	}
	fd, _, e1 := syscall.Syscall(SYS_MEMFD_CREATE, uintptr(unsafe.Pointer(_p0)), uintptr(flags), uintptr(0))
	if e1 != 0 {
		err = os.NewSyscallError("memfd_create", e1)
	}
	return
}

func main() {
	fd, err := MemfdCreate("hello", 1)
	file := os.NewFile(fd, "")

	buf, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	_, err = file.Write(buf)
	if err != nil {
		panic(err)
	}
	fname := fmt.Sprintf("/proc/%d/fd/%d", os.Getpid(), fd)
	log.Println(fname, "wait exec 20 second...")
	time.Sleep(20*time.Second)

	err = syscall.Exec(fname, []string{}, []string{})
	if err != nil {
		panic(err)
	}
}

