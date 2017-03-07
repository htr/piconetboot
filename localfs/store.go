package localfs

import (
	"net/url"
	"sync"

	"github.com/htr/piconetboot"
)

type localFsStore struct {
	path         string
	clientsCache []bootClient
	mu           *sync.RWMutex
}

var _ piconetboot.BootClientStore = (*localFsStore)(nil)

func NewStore(path string) (*localFsStore, error) {
	if !isDir(path) {
		// XXX
	}

	s := &localFsStore{
		path:         path,
		clientsCache: []bootClient{},
		mu:           new(sync.RWMutex),
	}

	return s, nil
}

func (s *localFsStore) Find(filter url.Values) (piconetboot.BootClient, error) {
	// XXX
	return nil, nil
}
