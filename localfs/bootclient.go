package localfs

import (
	"net/url"

	log "github.com/Sirupsen/logrus"
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
	mac := filter.Get("net0mac") //XXX ok for now (and probably forever.. undionly)
	asset := filter.Get("asset")
	serial := filter.Get("serial")
	uuid := filter.Get("uuid")

	log.WithFields(log.Fields{
		"mac":    mac,
		"asset":  asset,
		"serial": serial,
		"uuid":   uuid,
	}).Debug("bootClient.match")

	return false
}
