package logger

import (
	"github.com/sirupsen/logrus"
	"io"
	"mall-go/pkg/utils"
	"os"
)

var (
	Formatter = map[string]logrus.Formatter{
		"json": &logrus.JSONFormatter{TimestampFormat: "2006-01-02 15:04:05"},
		"text": &logrus.TextFormatter{TimestampFormat: "2006-01-02 15:04:05"},
		//"test": &inLog.TestFormatter{TimestampFormat: "2006-01-02 15:04:05"},
	}
	logDir string
	//ApiLog    = apiLog()

)

func setFileOut() (f io.Writer, err error) {
	err = os.MkdirAll(logDir, 777)
	if err != nil {
		return
	}
	logFile := logDir + "app.log"
	f, err = os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)
	if err != nil {
		return
	}

	return
}

type Logger struct {
	*logrus.Logger
}

func NewLogger(format, lvl string, reportCaller bool, savePath string) (l *Logger, err error) {
	lrus := logrus.StandardLogger()

	lrus.SetFormatter(Formatter[format])
	level, e := logrus.ParseLevel(lvl)
	if e != nil {
		return
	}
	lrus.SetLevel(level)
	lrus.SetReportCaller(reportCaller)
	logDir = utils.LogDir(savePath)
	l = &Logger{
		lrus,
	}
	l.setOutput()
	return
}
func (l *Logger) setOutput() {
	writer, e := setFileOut()
	if e != nil {

	}
	l.SetOutput(io.MultiWriter(l.Out, writer))
}
