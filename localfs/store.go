package localfs

import "github.com/htr/piconetboot"

type localFsStore struct {
	path string
}

var _ piconetboot.BootClientStore = (*localFsStore)(nil)
