package logging

import (
	"bytes"

	"os"
	"os/exec"
	"testing"

	"github.com/gcarrenho/logging/utils"
)

func TestPanic(t *testing.T) {
	os.Setenv("ENVIRONMENT", "TEST")
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	tests := []struct {
		name string
		log  *Logging
	}{
		{
			name: "Should get media succesfully",
			log: &Logging{
				Index:    utils.PtrTo("test" + "-" + os.Getenv("ENVIRONMENT")),
				LabelApp: utils.PtrTo("test"),
				Message:  "Test",
			},
		},
	}

	for _, tt := range tests {
		out := &bytes.Buffer{}
		loggingService := NewLoggingService(out)

		loggingService.Panic(tt.log)

		if got, want := string(out.String()), `{"level":"info","index":"test-TEST","label_app":"test","message":"Test"}`+"\n"; got != want {
			t.Errorf("invalid log output:\ngot:  %v\nwant: %v", got, want)
		}

	}
}

func TestDebug(t *testing.T) {
	os.Setenv("ENVIRONMENT", "TEST")
	tests := []struct {
		name string
		log  *Logging
	}{
		{
			name: "Should log debug",
			log: &Logging{
				Index:    utils.PtrTo("test" + "-" + os.Getenv("ENVIRONMENT")),
				LabelApp: utils.PtrTo("test"),
				Message:  "Test",
			},
		},
	}
	// · Runner · //

	for _, tt := range tests {

		// Prepare
		out := &bytes.Buffer{}
		loggingService := NewLoggingService(out)

		loggingService.Debug(tt.log)

		if got, want := out.String(), `{"level":"debug","index":"test-TEST","label_app":"test","message":"Test"}`+"\n"; got != want {
			t.Errorf("invalid log output:\ngot:  %v\nwant: %v", got, want)
		}
	}
}

func TestFatal(t *testing.T) {
	os.Setenv("ENVIRONMENT", "TEST")
	tests := []struct {
		name string
		log  *Logging
	}{
		{
			name: "Should log Fatal",
			log: &Logging{
				Index:    utils.PtrTo("test" + "-" + os.Getenv("ENVIRONMENT")),
				LabelApp: utils.PtrTo("test"),
				Message:  "Test",
			},
		},
	}
	// · Runner · //

	for _, tt := range tests {

		if os.Getenv("BE_CRASHER") == "1" {
			// Prepare
			out := &bytes.Buffer{}
			loggingService := NewLoggingService(out)

			loggingService.Fatal(tt.log)

			if got, want := out.String(), `{"level":"info","index":"test-TEST","label_app":"test","message":"Test"}`+"\n"; got != want {
				t.Errorf("invalid log output:\ngot:  %v\nwant: %v", got, want)
			}
			return
		}
		cmd := exec.Command(os.Args[0], "-test.run=TestFatal")
		cmd.Env = append(os.Environ(), "BE_CRASHER=1")
		err := cmd.Run()
		if e, ok := err.(*exec.ExitError); ok && !e.Success() {
			return
		}
		t.Fatalf("process ran with err %v, want exit status 1", err)
	}
}

func TestWarn(t *testing.T) {
	os.Setenv("ENVIRONMENT", "TEST")
	tests := []struct {
		name string
		log  *Logging
	}{
		{
			name: "Should log Warn",
			log: &Logging{
				Index:    utils.PtrTo("test" + "-" + os.Getenv("ENVIRONMENT")),
				LabelApp: utils.PtrTo("test"),
				Message:  "Test",
			},
		},
	}
	// · Runner · //

	for _, tt := range tests {

		out := &bytes.Buffer{}
		loggingService := NewLoggingService(out)

		loggingService.Warn(tt.log)

		if got, want := out.String(), `{"level":"warn","index":"test-TEST","label_app":"test","message":"Test"}`+"\n"; got != want {
			t.Errorf("invalid log output:\ngot:  %v\nwant: %v", got, want)
		}
	}
}

func TestInfo(t *testing.T) {
	os.Setenv("ENVIRONMENT", "TEST")
	tests := []struct {
		name string
		log  *Logging
	}{
		{
			name: "Should log Info",
			log: &Logging{
				Index:    utils.PtrTo("test" + "-" + os.Getenv("ENVIRONMENT")),
				LabelApp: utils.PtrTo("test"),
				Message:  "Test",
			},
		},
	}
	// · Runner · //
	for _, tt := range tests {

		out := &bytes.Buffer{}
		loggingService := NewLoggingService(out)

		loggingService.Info(tt.log)

		if got, want := out.String(), `{"level":"info","index":"test-TEST","label_app":"test","message":"Test"}`+"\n"; got != want {
			t.Errorf("invalid log output:\ngot:  %v\nwant: %v", got, want)
		}
	}
}

func TestError(t *testing.T) {
	os.Setenv("ENVIRONMENT", "TEST")
	tests := []struct {
		name string
		log  *Logging
	}{
		{
			name: "Should log error",
			log: &Logging{
				Index:    utils.PtrTo("test" + "-" + os.Getenv("ENVIRONMENT")),
				LabelApp: utils.PtrTo("test"),
				Message:  "Test",
			},
		},
	}
	// · Runner · //

	for _, tt := range tests {

		out := &bytes.Buffer{}
		loggingService := NewLoggingService(out)

		loggingService.Error(tt.log)

		if got, want := out.String(), `{"level":"error","index":"test-TEST","label_app":"test","message":"Test"}`+"\n"; got != want {
			t.Errorf("invalid log output:\ngot:  %v\nwant: %v", got, want)
		}
	}
}
