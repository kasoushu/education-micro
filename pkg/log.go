package pkg

import (
	"bytes"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/sirupsen/logrus"
	"os"
	"sort"
	"strings"
)

const (
	red     = 31
	yellow  = 33
	blue    = 36
	newBlue = 34
	grape   = 35
	gray    = 37
	green   = 32
)

type CustomFormatter struct{}

type CustomLogger struct {
	log *logrus.Logger
}

func NewLogger() log.Logger {
	logger := logrus.New()
	logger.Level = logrus.DebugLevel
	logger.Out = os.Stdout
	logger.Formatter = &CustomFormatter{}

	return &CustomLogger{
		log: logger,
	}
}

func (f *CustomFormatter) Format(e *logrus.Entry) ([]byte, error) {
	var b *bytes.Buffer
	if e.Buffer != nil {
		b = e.Buffer
	} else {
		b = &bytes.Buffer{}
	}
	levelColor := blue
	switch e.Level {
	case logrus.DebugLevel, logrus.TraceLevel:
		levelColor = green
	case logrus.WarnLevel:
		levelColor = yellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		levelColor = red
	case logrus.InfoLevel:
		levelColor = blue
	}
	timeStamp := e.Time.Format("2006-01-02 15:04:05")
	levelString := strings.ToUpper(e.Level.String()[0:4])
	p := fmt.Sprintf("\x1b[%dm[%s]\x1b[0m \x1b[%dm[%s]\x1b[0m", levelColor, levelString, grape, timeStamp)
	b.WriteString(p)
	b.WriteByte(' ')

	//caller := ""
	//if e.HasCaller() {
	//	funcVal := fmt.Sprintf("caller:%s() ", e.Caller.Function)
	//	//fileVal := fmt.Sprintf("%s:%d", e.Caller.File, e.Caller.Line)
	//
	//	//if f.CallerPrettyfier != nil {
	//	//	funcVal, fileVal = f.CallerPrettyfier(e.Caller)
	//	//}
	//	//if fileVal == "" {
	//	//	caller = funcVal
	//	//} else if funcVal == "" {
	//	//	caller = fileVal
	//	//} else {
	//	//	caller = fileVal + " " + funcVal
	//	//}
	//	caller = funcVal
	//}
	//b.WriteString(caller)

	data := make(map[string]interface{})
	var keys []string
	for k, v := range e.Data {
		data[k] = v
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, key := range keys {
		b.WriteString(key)
		b.WriteByte('=')
		b.WriteString(fmt.Sprintf("\"%s\" ", data[key]))
	}
	e.Message = strings.TrimSuffix(e.Message, "\n")
	//b.WriteString(e.Message)
	k := fmt.Sprintf("\x1b[%dm %s \x1b[0m", newBlue, e.Message)
	b.WriteString(k)
	b.WriteByte('\n')
	return b.Bytes(), nil
}

func (l *CustomLogger) Log(level log.Level, keyvals ...interface{}) (err error) {
	var (
		logrusLevel logrus.Level
		fields      logrus.Fields = make(map[string]interface{})
		msg         string
	)

	switch level {
	case log.LevelDebug:
		logrusLevel = logrus.DebugLevel
	case log.LevelInfo:
		logrusLevel = logrus.InfoLevel
	case log.LevelWarn:
		logrusLevel = logrus.WarnLevel
	case log.LevelError:
		logrusLevel = logrus.ErrorLevel
	default:
		logrusLevel = logrus.DebugLevel
	}

	if logrusLevel > l.log.Level {
		return
	}
	if len(keyvals) == 0 {
		return nil
	}
	if len(keyvals)%2 != 0 {
		keyvals = append(keyvals, "")
	}
	for i := 0; i < len(keyvals); i += 2 {
		key, ok := keyvals[i].(string)
		if !ok {
			continue
		}
		if key == logrus.FieldKeyMsg {
			msg, _ = keyvals[i+1].(string)
			continue
		}
		fields[key] = keyvals[i+1]
	}

	if len(fields) > 0 {
		l.log.WithFields(fields).Log(logrusLevel, msg)
	} else {
		l.log.Log(logrusLevel, msg)
	}

	return
}
