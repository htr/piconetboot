package localfs

import (
	"net/url"

	"github.com/htr/piconetboot"
)

type localFsStore struct {
	path string
}

var _ piconetboot.BootClientStore = (*localFsStore)(nil)

func NewStore(path string) (*localFsStore, error) {
	if !isDir(path) {
		// XXX
	}

	s := &localFsStore{
		path: path,
	}

	return s, nil
}

func (s *localFsStore) Find(filter url.Values) (piconetboot.BootClient, error) {
	// XXX
	return nil, nil
}
