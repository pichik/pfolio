package datacenter

import (
	"crypto/sha256"
	"fmt"
	"time"

	"github.com/pichik/webwatcher/src/misc"
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
	Method      string
	Screenshot  string
	HASH        string
	Collection  string
	Visible     int
	Timestamp   int64
}

var DeepCollectorPath string = "/collector"
var DeepCollectorRef string = misc.Config.Host + DeepCollectorPath

func CreateHash(secret string) string {
	currentTime := time.Now().UnixNano() / int64(time.Millisecond)
	hash := sha256.Sum256([]byte(fmt.Sprintf("%s%d", secret, currentTime)))
	return fmt.Sprintf("%x", hash)
}

func GetCollection(timestamp string) []*Data {
	return readData(timestamp)
}

func AddToCollection(data *Data) {
	saveNewData(data)
}

func RemoveFromCollection(hash string) {
	removeData(hash)
}
