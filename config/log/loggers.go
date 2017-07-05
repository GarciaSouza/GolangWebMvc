package log

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

//Trace Trace logger
var Trace *log.Logger

//Info Info logger
var Info *log.Logger

//Warning Warning logger
var Warning *log.Logger

//Error Error logger
var Error *log.Logger

func init() {
	InitLog(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)
}

//InitLog Initialize the logger
func InitLog(
	traceHandle io.Writer,
	infoHandle io.Writer,
	warningHandle io.Writer,
	errorHandle io.Writer) {

	Trace = log.New(traceHandle,
		"TRACE: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Info = log.New(infoHandle,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Warning = log.New(warningHandle,
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(errorHandle,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}
