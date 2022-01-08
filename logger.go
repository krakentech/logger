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

func Debug(msg string) {
	printLog(LoglvlDebug, msg)
}

func DebugWithValues(msg string, vals map[string]interface{}) {
	printLog(LoglvlDebug, msg)
	for k, v := range vals {
		printKeyVal(LoglvlDebug, k, v)
	}
}

func Info(msg string) {
	printLog(LoglvlInfo, msg)
}

func InfoWithValues(msg string, vals map[string]interface{}) {
	printLog(LoglvlInfo, msg)
	for k, v := range vals {
		printKeyVal(LoglvlInfo, k, v)
	}
}

func Special(msg string) {
	printLog(LoglvlSpecial, msg)
}

func SpecialWithValues(msg string, vals map[string]interface{}) {
	printLog(LoglvlSpecial, msg)
	for k, v := range vals {
		printKeyVal(LoglvlSpecial, k, v)
	}
}

func Err(msg string) {
	printLog(LoglvlErr, msg)
}

func ErrWithError(msg string, err error) {
	printLog(LoglvlErr, msg)
	printSubMsg(LoglvlErr, err.Error())
}

func ErrWithValues(msg string, vals map[string]interface{}) {
	printLog(LoglvlErr, msg)
	for k, v := range vals {
		printKeyVal(LoglvlErr, k, v)
	}
}

func Fatal(msg string, code int) {
	printLog(LoglvlFatal, msg)
	os.Exit(code)
}

func FatalWithError(msg string, err error, code int) {
	printLog(LoglvlFatal, msg)
	printSubMsg(LoglvlFatal, err.Error())
	os.Exit(code)
}

func printLog(llvl loglvl, msg string) {
	if llvl >= LOG_LVL {
		msgSlug := colors[llvl].Sprint(fmt.Sprintf("%s%s%s", SLUG_START, slugs[llvl], SLUG_END))
		fmt.Fprintf(OUT, "%s %s %s\n", time.Now().Format(TIME_FORMAT), msgSlug, msg)
	}
}

func printSubMsg(llvl loglvl, msg string) {
	if llvl >= LOG_LVL {
		fmt.Fprintf(OUT, "%s %s\n", subColors[llvl].Sprint("-"), colors[llvl].Sprint(msg))
	}
}

func printKeyVal(llvl loglvl, key, value interface{}) {
	if llvl >= LOG_LVL {
		slug := subColors[llvl].Sprintf("%s %s:", subColors[llvl].Sprint("-"), colorKey.Sprint(key))
		fmt.Fprintf(OUT, "%s %v\n", slug, value)
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
