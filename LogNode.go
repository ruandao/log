package log

import "fmt"

type LogNode struct {
	msg  string
	level int
}

func newLogNode(level int, v ...interface{}) *LogNode {
	node := &LogNode{
		level:level,
		msg:fmt.Sprintln(v...),
	}
	return node
}

func (ln *LogNode) String() string {
	s := ""
	switch ln.level {
	case LevelFatal:
		s += "[F] "
	case LevelError:
		s += "[E] "
	case LevelWarn:
		s += "[W] "
	case LevelInfo:
		s += "[I] "
	case LevelDebug:
		s += "[D] "
	}
	s += ln.msg
	return s
}