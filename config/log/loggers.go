package log

import (
	"io/ioutil"
	"log"
	"os"
)

var (
	ltrace   *log.Logger
	ldebug   *log.Logger
	linfo    *log.Logger
	lwarning *log.Logger
	lerror   *log.Logger
)

func init() {
	ltrace = log.New(ioutil.Discard,
		"TRACE: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	ldebug = log.New(os.Stdout,
		"DEBUG: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	linfo = log.New(os.Stdout,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	lwarning = log.New(os.Stdout,
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	lerror = log.New(os.Stderr,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

//Trace Trace
func Trace(v ...interface{}) {
	ltrace.Println(v)
}

//Tracef Format trace log, plus "\n" at the end of format
func Tracef(format string, v ...interface{}) {
	ltrace.Printf(format+"\n", v)
}

//Debug Debug
func Debug(v ...interface{}) {
	ldebug.Println(v)
}

//Debugf Format debug log, plus "\n" at the end of format
func Debugf(format string, v ...interface{}) {
	ldebug.Printf(format+"\n", v)
}

//Info Info
func Info(v ...interface{}) {
	linfo.Println(v)
}

//Infof Format info log, plus "\n" at the end of format
func Infof(format string, v ...interface{}) {
	linfo.Printf(format+"\n", v)
}

//Warning Warning
func Warning(v ...interface{}) {
	lwarning.Println(v)
}

//Warningf Format warning log, plus "\n" at the end of format
func Warningf(format string, v ...interface{}) {
	lwarning.Printf(format+"\n", v)
}

//Error Error
func Error(v ...interface{}) {
	lerror.Println(v)
}

//Errorf Format error log, plus "\n" at the end of format
func Errorf(format string, v ...interface{}) {
	lerror.Printf(format+"\n", v)
}

//Fatal Fatal
func Fatal(v ...interface{}) {
	lerror.Fatalln(v)
}

//Fatalf Format fatal log, plus "\n" at the end of format
func Fatalf(format string, v ...interface{}) {
	lerror.Printf(format+"\n", v)
}
