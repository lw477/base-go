package pocketlog_test

import (
	"learngo-pockets/logger/pocketlog"
	"testing"
)

type testWriter struct {
	contents string
}

func (tw *testWriter) Write(p []byte) (n int, err error) {
	tw.contents += string(p)
	return len(p), nil
}

func TestLogger_Levels(t *testing.T) {
	const (
		debugMsg = "this is debug message"
		infoMsg  = "this is info message"
		errMsg   = "this is error message"
	)

	type tc struct {
		level   pocketlog.Level
		wantOut string
	}

	tests := map[string]tc{
		"debug level": {
			level:   pocketlog.LevelDebug,
			wantOut: debugMsg + "\n" + infoMsg + "\n" + errMsg + "\n",
		},
		"info level": {
			level:   pocketlog.LevelInfo,
			wantOut: infoMsg + "\n" + errMsg + "\n",
		},
		"error level": {
			level:   pocketlog.LevelError,
			wantOut: errMsg + "\n",
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			tw := &testWriter{}
			logger := pocketlog.New(tt.level, pocketlog.WithOutput(tw))

			logger.Debugf(debugMsg)
			logger.Infof(infoMsg)
			logger.Errorf(errMsg)

			if tw.contents != tt.wantOut {
				t.Errorf("want:\n%q,\ngot:\n%q", tt.wantOut, tw.contents)
			}
		})
	}
}

func ExampleLogger_Debugf() {
	debugLogger := pocketlog.New(pocketlog.LevelDebug)
	debugLogger.Debugf("Hello,%s", "world")
	// Output:
	// Hello,world
}
