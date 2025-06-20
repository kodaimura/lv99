package externalapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"time"

	"lv99/config"
	"lv99/internal/domain/executor"
)

type HttpCodeExecutor struct {
	httpClient *http.Client
}

func NewHttpCodeExecutor() *HttpCodeExecutor {
	return &HttpCodeExecutor{
		httpClient: &http.Client{Timeout: 20 * time.Second},
	}
}

func (api *HttpCodeExecutor) Execute(in executor.CodeExecRequest) (executor.CodeExecResponse, error) {
	jsonData, err := json.Marshal(in)
	if err != nil {
		return executor.CodeExecResponse{}, fmt.Errorf("failed to marshal request: %w", err)
	}

	req, err := http.NewRequest("POST", config.CodeExecutorHost+"/exec/python", bytes.NewBuffer(jsonData))
	if err != nil {
		return executor.CodeExecResponse{}, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := api.httpClient.Do(req)
	if err != nil {
		return executor.CodeExecResponse{}, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	var res executor.CodeExecResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return executor.CodeExecResponse{}, fmt.Errorf("failed to decode response: %w", err)
	}

	re := regexp.MustCompile(`File\s+"\/tmp\/[^\"]+\.py",\s*`)
	cleaned := re.ReplaceAllString(res.Error, "")
	res.Error = cleaned

	return res, nil
}
