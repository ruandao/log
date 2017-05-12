package log

import (
	"fmt"
	"os"
	"sync"
)

const (
	All	   = iota
	LevelDebug
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
	Off
)

const maxChannelSize int = 1000

type Logger struct {
	level  int
	ch     chan *LogNode

	asyncToFile   bool // 是否异步存入文件
	filename      string	 // 存储日志的文件,默认为 ./output.log
	file          *os.File
	fileOnce      sync.Once
	enableFileLog bool	// 是否记录日志到日志文件

	outputToConsole bool // 是否输出到标准输出

}

func New() *Logger {
	l := &Logger{
		level:All,
		ch:make(chan *LogNode, maxChannelSize),

		asyncToFile:true,
		filename:"./output.log",
		enableFileLog:true,

		outputToConsole:true,
	}
	go l.run()

	return l
}

func (l *Logger) SetEnableFileLog(enable bool) {
	l.enableFileLog = enable
}
func (l *Logger) SetOutputToConsole(outputToConsole bool) {
	l.outputToConsole = outputToConsole
}
func (l *Logger) SetFileName(filename string) {
	l.filename = filename
}
func (l *Logger) SetAsyncToFile(asyncToFile bool) {
	l.asyncToFile = asyncToFile
}
func (l *Logger) SetLogLevel(level int) {
	l.level = level
}

func (l *Logger) createFile() {
	if l.file != nil || l.filename == "" {
		return 
	}
	file, err := os.OpenFile(l.filename, os.O_APPEND|os.O_CREATE, 0664)
	if err != nil {
		fmt.Printf("创建日志文件失败 %s", err)
		return
	}
	l.file = file
}

func (l *Logger) _output(logNode *LogNode)  {
	if l.outputToConsole {
		fmt.Print(logNode)
	}

	if !l.enableFileLog {
		return
	}
	l.fileOnce.Do(l.createFile)
	if l.file == nil {
		return
	}
	_, err := l.file.WriteString(logNode.String())
	if err != nil {
		fmt.Printf("写入日志文件失败 %s", err)
	}
}

func (l *Logger) run() {
	for ln := range l.ch {
		l._output(ln)
	}
}

func (l *Logger) output(logNode *LogNode) {
	if l.level >= logNode.level {
		if l.asyncToFile {
			l.ch <- logNode
		} else {
			l._output(logNode)
		}
	}
}
func (l *Logger) Debug(v ...interface{})  {
	logNode := newLogNode(LevelDebug, v...)
	l.output(logNode)
}

func (l *Logger) Info(v ...interface{})  {
	logNode := newLogNode(LevelInfo, v...)
	l.output(logNode)
}

func (l *Logger) Warnning(v ...interface{})  {
	logNode := newLogNode(LevelWarn, v...)
	l.output(logNode)
}

func (l *Logger) Error(v ...interface{})  {
	logNode := newLogNode(LevelError, v...)
	l.output(logNode)
}

func (l *Logger) Fatal(v ...interface{})  {
	logNode := newLogNode(LevelFatal, v...)
	l.output(logNode)
}