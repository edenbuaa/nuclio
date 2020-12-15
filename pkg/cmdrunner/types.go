/*
Copyright 2017 The Nuclio Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cmdrunner

type CaptureOutputMode int

const (
	CaptureOutputModeCombined CaptureOutputMode = iota
	CaptureOutputModeStdout
)

// RunOptions specifies runOptions to CmdRunner.Run
type RunOptions struct {
	WorkingDir        *string
	Stdin             *string
	Env               map[string]string
	LogRedactions     []string
	CaptureOutputMode CaptureOutputMode
}

// RunResult holds command execution returned values
type RunResult struct {
	Output   string
	Stderr   string
	ExitCode int
}

// CmdRunner specifies the interface to an underlying command runner
type CmdRunner interface {

	// Run runs a command, given runOptions
	Run(runOptions *RunOptions, format string, vars ...interface{}) (RunResult, error)
}