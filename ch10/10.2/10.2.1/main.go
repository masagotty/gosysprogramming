package main

import (
	"fmt"
	"sync"
	"syscall"
	"time"
)

type FileLock struct {
	l  sync.Mutex
	fd int
}

func NewFileLock(filename string) *FileLock {
	if filename == "" {
		panic("filename needed")
	}
	fd, err := syscall.Open(filename, syscall.O_CREAT|syscall.O_RDONLY, 0750)
	if err != nil {
		panic(err)
	}
	return &FileLock{fd: fd}
}
func (m *FileLock) Lock() {
	m.l.Lock()
	if err := syscall.Flock(m.fd, syscall.LOCK_EX); err != nil {
		panic(err)
	}
}
func (m *FileLock) UnLock() {
	if err := syscall.Flock(m.fd, syscall.LOCK_UN); err != nil {
		panic(err)
	}
}

// 10.2.3
func main() {
	l := NewFileLock("main.go")
	fmt.Println("try Locking...")
	l.Lock()
	fmt.Println("Locked!")
	time.Sleep(10 * time.Second)
	l.UnLock()
	fmt.Println("unlock")
}