package logger

func Info(args ...any) {
	Blogger.Info(args...)
	defer Blogger.MustFlush()
}

func Error(args ...any) {
	Blogger.Error(args...)
	defer Blogger.MustFlush()
}

func Debug(args ...any) {
	Blogger.Debug(args...)
	defer Blogger.MustFlush()
}

func Warn(args ...any) {
	Blogger.Warn(args...)
	defer Blogger.MustFlush()
}

func Fatal(args ...any) {
	Blogger.Fatal(args...)
	defer Blogger.MustFlush()
}

func Panic(args ...any) {
	Blogger.Panic(args...)
	defer Blogger.MustFlush()
}

func Infof(format string, args ...any) {
	Blogger.Infof(format, args...)
	defer Blogger.MustFlush()
}

func Errorf(format string, args ...any) {
	Blogger.Errorf(format, args...)
	defer Blogger.MustFlush()
}

func Debugf(format string, args ...any) {
	Blogger.Debugf(format, args...)
	defer Blogger.MustFlush()
}

func Warnf(format string, args ...any) {
	Blogger.Warnf(format, args...)
	defer Blogger.MustFlush()
}

func Fatalf(format string, args ...any) {
	Blogger.Fatalf(format, args...)
	defer Blogger.MustFlush()
}

func Panicf(format string, args ...any) {
	Blogger.Panicf(format, args...)
	defer Blogger.MustFlush()
}
