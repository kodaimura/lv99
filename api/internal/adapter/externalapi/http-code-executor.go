package externalapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"

	"lv99/config"
	"lv99/internal/core"
)

type HttpCodeExecutor struct {
	httpClient *http.Client
}

func NewHttpCodeExecutor() *HttpCodeExecutor {
	return &HttpCodeExecutor{
		httpClient: &http.Client{Timeout: 20 * time.Second},
	}
}

func (api *HttpCodeExecutor) Execute(in core.CodeExecRequest) (core.CodeExecResponse, error) {
	executorURL, err := buildExecutorURL()
	if err != nil {
		return core.CodeExecResponse{}, err
	}

	jsonData, err := json.Marshal(in)
	if err != nil {
		return core.CodeExecResponse{}, fmt.Errorf("failed to marshal request: %w", err)
	}

	req, err := http.NewRequest("POST", executorURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return core.CodeExecResponse{}, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := api.httpClient.Do(req)
	if err != nil {
		return core.CodeExecResponse{}, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		return core.CodeExecResponse{}, fmt.Errorf("code executor returned status %d", resp.StatusCode)
	}

	var res core.CodeExecResponse
	if err := json.NewDecoder(io.LimitReader(resp.Body, 1<<20)).Decode(&res); err != nil {
		return core.CodeExecResponse{}, fmt.Errorf("failed to decode response: %w", err)
	}

	re := regexp.MustCompile(`File\s+"\/tmp\/[^\"]+\.py",\s*`)
	cleaned := re.ReplaceAllString(res.Error, "")
	res.Error = cleaned

	return res, nil
}

func buildExecutorURL() (string, error) {
	host := strings.TrimRight(config.CodeExecutorHost, "/")
	if host == "" {
		return "", fmt.Errorf("code executor host is not configured")
	}

	parsed, err := url.Parse(host)
	if err != nil {
		return "", fmt.Errorf("invalid code executor host: %w", err)
	}
	if parsed.Scheme != "http" && parsed.Scheme != "https" {
		return "", fmt.Errorf("invalid code executor scheme: %s", parsed.Scheme)
	}
	if parsed.Host == "" {
		return "", fmt.Errorf("invalid code executor host")
	}

	return host + "/exec/python", nil
}
