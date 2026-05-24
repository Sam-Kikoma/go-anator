package pocketlog

type Logger struct{
	threshold Level
}
func New(threshold Level) *Logger{
	return &Logger{
		threshold: threshold,
	}
}

// Debugf formats and prints a message if the log level is debug or higher
func (l *Logger) Debugf(fromat string, args ...any){

}

// Infof formats and prints a message if the log level is info or higher
func (l *Logger) Infof(format string, args ...any){

}

func (l *Logger) Errorf(format string, args ...any){

}
