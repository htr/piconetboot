package piconetboot

import "net/url"

type BootClientStore interface {
	Find(filter url.Values) (BootClient, error)
}

type BootClient interface {
	BootScript() string
}
