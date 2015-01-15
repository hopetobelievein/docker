// +build !windows

package portmapper

import (
	"net"
	"os/exec"
	"strconv"
	"syscall"

	"github.com/docker/docker/pkg/reexec"
)

func NewProxyCommand(proto string, hostIP net.IP, hostPort int, containerIP net.IP, containerPort int) UserlandProxy {
	args := []string{
		userlandProxyCommandName,
		"-proto", proto,
		"-host-ip", hostIP.String(),
		"-host-port", strconv.Itoa(hostPort),
		"-container-ip", containerIP.String(),
		"-container-port", strconv.Itoa(containerPort),
	}

	return &proxyCommand{
		cmd: &exec.Cmd{
			Path: reexec.Self(),
			Args: args,
			SysProcAttr: &syscall.SysProcAttr{
				Pdeathsig: syscall.SIGTERM, // send a sigterm to the proxy if the daemon process dies
			},
		},
	}
}
