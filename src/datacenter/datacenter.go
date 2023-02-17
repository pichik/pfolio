package datacenter

import (
	"crypto/sha256"
	"fmt"
	"time"

	"github.com/pichik/WebWatcher/src/misc"
)

type Data struct {
	Location    string
	Cookies     string
	Referrer    string
	UserAgent   string
	BrowserTime string
	Origin      string
	DOM         string
	IP          string
	HASH        string
	Method      string
	Screenshot  string
}

type DataCollection struct {
	DeepData []*Data
	Request  []*Data
}

var Collection DataCollection
var DeepCollectorPath string = "/collector"
var DeepCollectorRef string = misc.Config.Host + DeepCollectorPath

func CreateHash(secret string) string {
	currentTime := time.Now().String()
	hash := sha256.Sum256([]byte(secret + currentTime))
	return fmt.Sprintf("%x", hash)
}
