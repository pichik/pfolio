package misc

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

type LogFormatter struct{}

var RequestLog = log.New()
var ErrorLog = log.New()
var DebugLog = log.New()

func ImportLogs() {
	fh, err := os.OpenFile(LogsDir+"errors.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0777)

	ErrorLog.SetOutput(fh)
	ErrorLog.SetFormatter(&LogFormatter{})

	fh, err = os.OpenFile(LogsDir+"requests.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777)
	if err != nil {
		ErrorLog.Printf("%s", err)
	}
	RequestLog.SetOutput(fh)
	RequestLog.SetFormatter(&LogFormatter{})

	fh, err = os.OpenFile(LogsDir+"debug.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0777)
	if err != nil {
		ErrorLog.Printf("%s", err)
	}
	DebugLog.SetOutput(fh)
	DebugLog.SetFormatter(&LogFormatter{})

}

func (f *LogFormatter) Format(entry *log.Entry) ([]byte, error) {
	return []byte(fmt.Sprintf("[%s] %s\n",
		entry.Time.Format("02.01.2006 | 15:04:05"),
		entry.Message,
	)), nil
}
