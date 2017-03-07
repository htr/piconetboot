package localfs

import (
	"net/url"

	"github.com/htr/piconetboot"
)

type localFsStore struct {
	path string
}

var _ piconetboot.BootClientStore = (*localFsStore)(nil)

func (s *localFsStore) Find(filter url.Values) (piconetboot.BootClient, error) {
	return nil, nil
}
