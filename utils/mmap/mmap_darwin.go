// +build darwin

package mmap

import (
	"os"
)

func Mmap(fd *os.File, writable bool, size int64) ([]byte, error) {
	return mmap(fd, writable, size)
}

// Munmap unmaps a previously mapped slice.
func Munmap(b []byte) error {
	return munmap(b)
}

// Madvise uses the madvise system call to give advise about the use of memory
// when using a slice that is memory-mapped to a file. Set the readahead flag to
// false if page references are expected in random order.
func Madvise(b []byte, readahead bool) error {
	return madvise(b, readahead)
}

// Msync would call sync on the mmapped data.
func Msync(b []byte) error {
	return msync(b)
}