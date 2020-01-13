package logging

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

// Level level of log
type Level int

var (
	//F log file object
	F *os.File

	//DefaultPrefix default prefix
	DefaultPrefix = ""
	//DefaultCallerDepth default caller depth
	DefaultCallerDepth = 2

	logger     *log.Logger
	logPrefix  = ""
	levelFlags = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
)

const (
	//DEBUG debug level
	DEBUG Level = iota
	//INFO info level
	INFO
	//WARNING waring level
	WARNING
	//ERROR error level
	ERROR
	//FATAL fatal level
	FATAL
)

func init() {
	filePath := getLogFileFullPath()
	F = openLogFile(filePath)

	logOut := io.MultiWriter(F, os.Stdout)
	logger = log.New(logOut, DefaultPrefix, log.LstdFlags)
}

//Debug debug level
func Debug(v ...interface{}) {
	setPrefix(DEBUG)
	logger.Println(v...)
}

//Info info level
func Info(v ...interface{}) {
	setPrefix(INFO)
	logger.Println(v...)
}

//Warn warn level
func Warn(v ...interface{}) {
	setPrefix(WARNING)
	logger.Println(v...)
}

//Error error level
func Error(v ...interface{}) {
	setPrefix(ERROR)
	logger.Println(v...)
}

//Fatal fatal level
func Fatal(v ...interface{}) {
	setPrefix(FATAL)
	logger.Println(v...)
}

func setPrefix(level Level) {
	_, file, line, ok := runtime.Caller(DefaultCallerDepth)
	if ok {
		logPrefix = fmt.Sprintf("[%s][%s:%d]", levelFlags[level], filepath.Base(file), line)
	} else {
		logPrefix = fmt.Sprintf("[%s]", levelFlags[level])
	}

	logger.SetPrefix(logPrefix)
}
