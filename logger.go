package logger

import (
	"fmt"

	"io"
	"os"
	"time"

	"github.com/fatih/color"
)

// Logger Settings
var OUT io.Writer = os.Stdout
var TIME_FORMAT string = time.RFC3339
var LOG_LVL loglvl = LoglvlDebug
var SLUG_START string = "["
var SLUG_END string = "]"
var slugs []string = []string{"debug", "info", "special", "error", "fatal"}

var subColors []*color.Color = []*color.Color{
	color.New(color.FgHiBlue),
	color.New(color.FgHiCyan),
	color.New(color.FgHiGreen),
	color.New(color.FgHiRed),
	color.New(color.FgHiMagenta),
}

var colors []*color.Color = []*color.Color{
	color.New(color.FgBlue),
	color.New(color.FgCyan),
	color.New(color.FgGreen),
	color.New(color.FgRed),
	color.New(color.FgMagenta),
}

var colorKey = color.New(color.FgHiWhite)

type loglvl int

const (
	LoglvlDebug   loglvl = 0
	LoglvlInfo    loglvl = 1
	LoglvlSpecial loglvl = 2
	LoglvlErr     loglvl = 3
	LoglvlFatal   loglvl = 4
)

func SetSpecialSlug(slug string) {
	slugs[LoglvlSpecial] = slug
}

func Debug(format string, a ...interface{}) {
	printLog(LoglvlDebug, format, a...)
}

func DebugWithValues(values map[string]interface{}, format string, a ...interface{}) {
	printLog(LoglvlDebug, format, a...)
	for k, v := range values {
		printKeyVal(LoglvlDebug, k, v)
	}
}

func Info(format string, a ...interface{}) {
	printLog(LoglvlInfo, format, a...)
}

func InfoWithValues(values map[string]interface{}, format string, a ...interface{}) {
	printLog(LoglvlInfo, format, a...)
	for k, v := range values {
		printKeyVal(LoglvlInfo, k, v)
	}
}

func Special(format string, a ...interface{}) {
	printLog(LoglvlSpecial, format, a...)
}

func SpecialWithValues(values map[string]interface{}, format string, a ...interface{}) {
	printLog(LoglvlSpecial, format, a...)
	for k, v := range values {
		printKeyVal(LoglvlSpecial, k, v)
	}
}

func Err(format string, a ...interface{}) {
	printLog(LoglvlErr, format, a...)
}

func ErrWithError(err error, format string, a ...interface{}) {
	printLog(LoglvlErr, format, a...)
	printSubMsg(LoglvlErr, err.Error())
}

func ErrWithValues(values map[string]interface{}, format string, a ...interface{}) {
	printLog(LoglvlErr, format, a...)
	for k, v := range values {
		printKeyVal(LoglvlErr, k, v)
	}
}

func Fatal(code int, format string, a ...interface{}) {
	printLog(LoglvlFatal, format, a...)
	os.Exit(code)
}

func FatalWithError(code int, err error, format string, a ...interface{}) {
	printLog(LoglvlFatal, format, a...)
	printSubMsg(LoglvlFatal, err.Error())
	os.Exit(code)
}

func printLog(llvl loglvl, format string, a ...interface{}) {
	if llvl >= LOG_LVL {
		msgSlug := colors[llvl].Sprint(fmt.Sprintf("%s%s%s", SLUG_START, slugs[llvl], SLUG_END))
		_, err := fmt.Fprintf(OUT, "%s %s %s\n", time.Now().Format(TIME_FORMAT), msgSlug, fmt.Sprintf(format, a...))
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}

func printSubMsg(llvl loglvl, msg string) {
	if llvl >= LOG_LVL {
		_, err := fmt.Fprintf(OUT, "%s %s\n", subColors[llvl].Sprint("-"), colors[llvl].Sprint(msg))
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}

func printKeyVal(llvl loglvl, key, value interface{}) {
	if llvl >= LOG_LVL {
		slug := subColors[llvl].Sprintf("%s %s:", subColors[llvl].Sprint("-"), colorKey.Sprint(key))
		_, err := fmt.Fprintf(OUT, "- %s: %v\n", slug, value)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}

func SeeColors() {
	char := "█"

	color.New(color.FgHiBlack).Print(char)
	color.New(color.FgHiBlue).Print(char)
	color.New(color.FgHiCyan).Print(char)
	color.New(color.FgHiGreen).Print(char)
	color.New(color.FgHiMagenta).Print(char)
	color.New(color.FgHiRed).Print(char)
	color.New(color.FgHiWhite).Print(char)
	color.New(color.FgHiYellow).Print(char)
	fmt.Println("")
	color.New(color.FgBlack).Print(char)
	color.New(color.FgBlue).Print(char)
	color.New(color.FgCyan).Print(char)
	color.New(color.FgGreen).Print(char)
	color.New(color.FgMagenta).Print(char)
	color.New(color.FgRed).Print(char)
	color.New(color.FgWhite).Print(char)
	color.New(color.FgYellow).Print(char)
	fmt.Println("")
}
