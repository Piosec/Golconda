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
	    log.Fatal(err)
	}
	// For demo purposes, create two backend for os.Stderr.
	stdout := logging.NewLogBackend(os.Stderr, "", 0)
	fileout := logging.NewLogBackend(f, "", 1)

	// For messages written to backend2 we want to add some additional
	// information to the output, including the used log level and the name of
	// the function.
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

//	log.Debugf("debug %s","test  debug")
//	log.Info("info")
//	log.Notice("notice")
//	log.Warning("warning")
//	log.Error("err")
//	log.Critical("crit")
}

