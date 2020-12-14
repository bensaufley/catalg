package log

import (
	"os"

	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetOutput(os.Stdout)
}

type Entry = logrus.Entry
type Fields = logrus.Fields

var WithField = logrus.WithField
var WithFields = logrus.WithFields
var WithError = logrus.WithError
var WithTime = logrus.WithTime
var WithContext = logrus.WithContext

var Trace = logrus.Trace
var Tracef = logrus.Tracef

var Debug = logrus.Debug
var Debugf = logrus.Debugf

var Info = logrus.Info
var Infof = logrus.Infof

var Warn = logrus.Warn
var Warnf = logrus.Warnf

var Error = logrus.Error
var Errorf = logrus.Errorf

var Fatal = logrus.Fatal
var Fatalf = logrus.Fatalf
