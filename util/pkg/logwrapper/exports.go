package logwrapper

func Fatal(args ...interface{})  {
	Log.Fatal(args...)
}

func Fatalf(format string,args ...interface{})  {
	Log.Fatalf(format,args...)
}
