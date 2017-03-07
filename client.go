package piconetboot

import "net/url"

type BootClientStore interface {
	Find(filter url.Values) (BootClient, err)
}

type BootClient interface {
	// TODO
}
