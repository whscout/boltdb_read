package main

import (
	"fmt"
	"unsafe"
)

const maxAllocSize = 0x7FFFFFFF

type pgid uint64

// branchPageElement represents a node on a branch page.
type branchPageElement struct {
	pos   uint32 //该元信息和真实key之间的偏移量
	ksize uint32
	pgid  pgid
}

// key returns a byte slice of the node key.
func (n *branchPageElement) key() []byte {
	buf := (*[maxAllocSize]byte)(unsafe.Pointer(n))
	fmt.Println(len(buf), cap(buf))
	// pos~ksize
	return (*[maxAllocSize]byte)(unsafe.Pointer(&buf[n.pos]))[:n.ksize]
}

func main() {
}
