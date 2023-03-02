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
}

var collection []*Data

var DeepCollectorPath string = "/collector"
var DeepCollectorRef string = misc.Config.Host + DeepCollectorPath

func CreateHash(secret string) string {
	currentTime := time.Now().String()
	hash := sha256.Sum256([]byte(secret + currentTime))
	return fmt.Sprintf("%x", hash)
}

func GetCollection() []*Data {
	return collection
}

func AddToCollection(data *Data) {
	collection = append([]*Data{data}, collection...)
}

func RemoveFromCollection(hash string) {
	for s, cd := range collection {
		if cd.HASH == hash {
			collection = append(collection[:s], collection[s+1:]...)
		}
	}
}

func ClearCollection() {
	collection = []*Data{}
}
