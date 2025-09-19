package autoname

import (
	"runtime"
	"strconv"
	"strings"
)

const (
	fieldSeparator = ":"
	undefined      = "?"
)

func fileLineName(file string, line int) string {
	builder := strings.Builder{}

	builder.WriteString(file)
	builder.WriteString(fieldSeparator)
	builder.WriteString(strconv.Itoa(line))

	return builder.String()
}

func GetRuntimeFunc(skipFrames int) string {
	pc, file, line, ok := runtime.Caller(skipFrames)
	if !ok {
		return undefined
	}

	fn := runtime.FuncForPC(pc)
	if fn == nil {
		return fileLineName(file, line)
	}

	return fn.Name()
}
