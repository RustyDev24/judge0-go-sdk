package judge0

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Submission struct {
	Source                               string        `json:"source_code,omitempty"`
	Language                             LanguageAlias `json:"language_id,omitempty"`
	AdditionalFiles                      *string       `json:"additional_files,omitempty"`
	CompilerOptions                      *string       `json:"compiler_options,omitempty"`
	CommandLineArguments                 *string       `json:"command_line_arguments,omitempty"`
	TestInput                            string        `json:"stdin,omitempty"`
	ExpectedOutput                       *string       `json:"expected_output,omitempty"`
	CpuTimeLimit                         *float32      `json:"cpu_time_limit,omitempty"`
	CpuExtraTime                         *float32      `json:"cpu_extra_time,omitempty"`
	WallTimeLimit                        *float32      `json:"wall_time_limit,omitempty"`
	MemoryLimit                          *float32      `json:"memory_limit,omitempty"`
	StackLimit                           *int          `json:"stack_limit,omitempty"`
	MaxProcessAndOrThreads               *int          `json:"max_process_and_or_threads,omitempty"`
	EnablePerProcessAndThreadTimeLimit   bool          `json:"enable_per_process_and_thread_time_limit,omitempty"`
	EnablePerProcessAndThreadMemoryLimit bool          `json:"enable_per_process_and_thread_memory_limit,omitempty"`
	MaxFileSize                          *int          `json:"max_file_size,omitempty"`
	RedirectStderrToStdin                bool          `json:"redirect_stderr_to_stdin,omitempty"`
	EnableNetwork                        bool          `json:"enable_network,omitempty"`
	NumberOfRuns                         int           `json:"number_of_runs,omitempty"`
	CallbackURL                          string        `json:"callback_url,omitempty"`
}

func (c *Client) CreateSubmission(submission *Submission, wait bool) (string, error) {
	js, err := json.Marshal(submission)
	if err != nil {
		return "", err
	}

	res, err := c.doRequest(http.MethodPost, fmt.Sprintf(SubmissionEndpoint, wait), js)
	if err != nil {
		// TODO: Handle errors
		return "", err
	}

	return "", nil
}
