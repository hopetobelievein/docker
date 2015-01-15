// +build !windows

package aufs

import (
	"fmt"
	"path"
	"syscall"

	"github.com/docker/docker/daemon/graphdriver"
)

func checkFSCompatibility(root string) error {
	rootdir := path.Dir(root)

	var buf syscall.Statfs_t
	if err := syscall.Statfs(rootdir, &buf); err != nil {
		return fmt.Errorf("Couldn't stat the root directory: %s", err)
	}

	for _, magic := range incompatibleFsMagic {
		if graphdriver.FsMagic(buf.Type) == magic {
			return graphdriver.ErrIncompatibleFS
		}
	}

	return nil
}
