package data

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"
)

var home, _ = os.UserHomeDir()

var TemplateDir string
var AssetsDir string
var LogsDir string
var DataDir string
var CertDir string

type ConfigData struct {
	Hostname  string
	Host      string
	Token     string
	Directory string
}

var Config ConfigData

func init() {
	file, err := ioutil.ReadFile(home + "/.pfconfig")

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(file, &Config)
	if err != nil {
		panic(err)
	}

	Config.Host = "https://" + Config.Hostname
	Config.Directory = strings.ReplaceAll(Config.Directory, "$HOME", home)

	TemplateDir = Config.Directory + "/templates/"
	AssetsDir = Config.Directory + "/assets/"
	LogsDir = Config.Directory + "/logs/"
	DataDir = Config.Directory + "/data/"
	CertDir = Config.Directory + "/certs/"

}
