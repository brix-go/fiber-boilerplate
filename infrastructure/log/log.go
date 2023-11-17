package infrastructure

import (
	"github.com/sirupsen/logrus"
	"sync"
)

type LogCustom struct {
	Logrus *logrus.Logger
}

var logInstance *LogCustom
var once sync.Once

func NewLogCustom() *LogCustom {
	var log *logrus.Logger
	log = logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{PrettyPrint: true})
	//log.SetReportCaller(true)
	//log.Out = os.Stdout

	once.Do(func() {
		logInstance = &LogCustom{
			log,
		}
	})
	return logInstance
}
