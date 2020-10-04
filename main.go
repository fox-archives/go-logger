package logger

import (
	"fmt"
	"os"
)

func isColor() bool {
	_, isNoColor := os.LookupEnv("NO_COLOR")
	if isNoColor || os.Getenv("TERM") == "dumb" {
		return false
	}

	fi, err := os.Stdout.Stat()
	if fi.Mode()&os.ModeCharDevice != 0 && err == nil {
		return true
	}

	return false
}

// Emergency log
func Emergency(text string, args ...interface{}) {
	text = fmt.Sprintf(text, args...)

	if isColor() {
		fmt.Printf("\033[1;35mEMERG ▶\033[0;35m %s\033[m", text)
		return
	}

	fmt.Printf("EMERG ▶ %s", text)
}

// Alert log
func Alert(text string, args ...interface{}) {
	text = fmt.Sprintf(text, args...)

	if isColor() {
		fmt.Printf("\033[1;35mALERT ▶\033[0;35m %s\033[m", text)
		return
	}

	fmt.Printf("ALERT ▶ %s", text)
}

// Critical log
func Critical(text string, args ...interface{}) {
	text = fmt.Sprintf(text, args...)

	if isColor() {
		fmt.Printf("\033[1;31mCRITL ▶\033[0;31m %s\033[m", text)
		return
	}

	fmt.Printf("CRITL ▶ %s", text)
}

// Error log
func Error(text string, args ...interface{}) {
	text = fmt.Sprintf(text, args...)

	if isColor() {
		fmt.Printf("\033[1;31mERROR ▶\033[0;31m %s\033[m", text)
		return
	}

	fmt.Printf("ERROR ▶ %s", text)
}

// Warning log
func Warning(text string, args ...interface{}) {
	text = fmt.Sprintf(text, args...)
	if isColor() {
		fmt.Printf("\033[1;33mWARNG ▶\033[0;33m %s\033[m", text)
		return
	}

	fmt.Printf("WARNG ▶ %s", text)
}

// Notice log
func Notice(text string, args ...interface{}) {
	text = fmt.Sprintf(text, args...)
	if isColor() {
		fmt.Printf("\033[1;94mNOTCE ▶\033[0;94m %s\033[m", text)
		return
	}

	fmt.Printf("NOTCE ▶ %s", text)
}

// Informational log
func Informational(text string, args ...interface{}) {
	text = fmt.Sprintf(text, args...)

	if isColor() {
		fmt.Printf("\033[1;32mINFRM ▶\033[0;32m %s\033[m", text)
		return
	}

	fmt.Printf("INFRM ▶ %s", text)
}

// Debug log
func Debug(text string, args ...interface{}) {
	text = fmt.Sprintf(text, args...)

	if isColor() {
		fmt.Printf("\033[1;36mDEBUG ▶ \033[0;36m%s\033[m", text)
		return
	}

	fmt.Printf("DEBUG ▶ %s", text)
}
