package log

import (
	"os"
	"github.com/op/go-logging"
)

var Log = logging.MustGetLogger("logger")

// Example format string. Everything except the message has a custom color
// which is dependent on the log level. Many fields have a custom output
// formatting too, eg. the time returns the hour down to the milli second.
var format = logging.MustStringFormatter(
	`%{color}%{time:15:04:05.000} %{shortfile} %{shortfunc} ▶ %{level} %{color:reset} %{message}`,
)

var NoColorFormat = logging.MustStringFormatter(
	`%{time:15:04:05.000} %{shortfile} %{shortfunc} ▶ %{level} %{message}`,
)
func Init(level int) {

	var stdoutLevel logging.LeveledBackend
	var fileoutLevel logging.LeveledBackend

	f, err := os.OpenFile("golconda.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
	    Log.Fatal(err)
	}
	// First one output to stdout second to a file.
	// 0 --> no date
	// 1 --> print date
	stdout := logging.NewLogBackend(os.Stdout, "", 0)
	fileout := logging.NewLogBackend(f, "", 1)

	stdoutFormatter := logging.NewBackendFormatter(stdout,format)
	fileoutFormatter := logging.NewBackendFormatter(fileout,NoColorFormat)
	switch level {
		case 1:
		stdoutLevel = logging.AddModuleLevel(stdoutFormatter)
		fileoutLevel = logging.AddModuleLevel(fileoutFormatter)
		stdoutLevel.SetLevel(logging.INFO,"")
		fileoutLevel.SetLevel(logging.INFO,"")
		case 2:
		stdoutLevel = logging.AddModuleLevel(stdoutFormatter)
		fileoutLevel = logging.AddModuleLevel(fileoutFormatter)
		stdoutLevel.SetLevel(logging.DEBUG,"")
		fileoutLevel.SetLevel(logging.DEBUG,"")
		default:
		stdoutLevel = logging.AddModuleLevel(stdoutFormatter)
		fileoutLevel = logging.AddModuleLevel(fileoutFormatter)
		stdoutLevel.SetLevel(logging.ERROR,"")
		fileoutLevel.SetLevel(logging.ERROR,"")
	}

	// Set the backends to be used.
	logging.SetBackend(fileoutLevel, stdoutLevel)

//	Log.Debugf("debug %s","test  debug")
//	Log.Info("info")
//	log.Notice("notice")
//	log.Warning("warning")
//	log.Error("err")
//	log.Critical("crit")
}

