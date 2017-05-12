package log

var log = New()

func Debug(v ...interface{})  {
	log.Debug(v...)
}

func Info(v ...interface{})  {
	log.Info(v...)
}

func Warnning(v ...interface{})  {
	log.Warnning(v...)
}

func Error(v ...interface{})  {
	log.Error(v...)
}

func Fatal(v ...interface{})  {
	log.Fatal(v...)
}