// +build windows

package pty

func ioctl(fd, cmd, ptr uintptr) error {
	return nil
}
