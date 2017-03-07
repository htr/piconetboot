package localfs

import (
	"net/url"

	"github.com/htr/piconetboot"
)

type bootClient struct {
	data struct {
		Matchers []struct {
			Serial string
			Asset  string
			UUID   string
			Mac    string
		}
		BootScript string
		Name       string
	}
}

var _ piconetboot.BootClient = (*bootClient)(nil)

func (c bootClient) BootScript() string {
	return c.data.BootScript
}

func (c bootClient) match(filter url.Values) bool {
	mac := vals.Get("net0mac") //XXX ok for now (and probably forever.. undionly)
	asset := vals.Get("asset")
	serial := vals.Get("serial")
	uuid := vals.Get("uuid")

	return false
}
