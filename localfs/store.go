package localfs

import (
	"errors"
	"fmt"
	"net/url"
	"os"
	"sync"

	"github.com/htr/piconetboot"
)

type localFsStore struct {
	path         string
	clientsCache []bootClient
	mu           *sync.RWMutex
}

var _ piconetboot.BootClientStore = (*localFsStore)(nil)

var errNoMatch = errors.New("unable to find a matching boot client")

func NewStore(path string) (*localFsStore, error) {
	if !isDir(path) {
		return nil, fmt.Errorf("path %s is not a directory", path)
	}

	s := &localFsStore{
		path:         path,
		clientsCache: []bootClient{},
		mu:           new(sync.RWMutex),
	}

	s.updateCache()

	return s, nil
}

func (s *localFsStore) Find(filter url.Values) (piconetboot.BootClient, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for _, client := range s.clientsCache {
		if client.match(filter) {
			return client, nil
		}
	}

	return nil, errNoMatch
}

func (s *localFsStore) updateCache() {
	// XXX

}

func isDir(path string) bool {
	fi, err := os.Stat(path)
	return err == nil && fi.IsDir()
}
