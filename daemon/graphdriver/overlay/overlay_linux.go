// +build linux

package overlay

import (
	"path"
	"syscall"

	log "github.com/Sirupsen/logrus"
	"github.com/docker/docker/daemon/graphdriver"
)

func checkFSCompatibility(home string) error {
	var buf syscall.Statfs_t
	if err := syscall.Statfs(path.Dir(home), &buf); err != nil {
		return err
	}

	switch graphdriver.FsMagic(buf.Type) {
	case graphdriver.FsMagicBtrfs:
		log.Error("'overlay' is not supported over btrfs.")
		return graphdriver.ErrIncompatibleFS
	case graphdriver.FsMagicAufs:
		log.Error("'overlay' is not supported over aufs.")
		return graphdriver.ErrIncompatibleFS
	}

	return nil
}
