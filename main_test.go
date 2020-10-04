package logger

import (
	"fmt"
	"testing"
)

func TestPrintInfo(t *testing.T) {
	Emergency("emergency\n")
	Alert("alert\n")
	Critical("critical\n")
	Error("error\n")
	Warning("warning\n")
	Notice("notice\n")
	Informational("informational\n")
	fmt.Println()
}
