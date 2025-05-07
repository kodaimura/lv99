package externalapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"lv99/config"
	"lv99/internal/service"
)

type HttpCodeExecutor struct {
	httpClient *http.Client
}

func NewHttpCodeExecutor() *HttpCodeExecutor {
	return &HttpCodeExecutor{
		httpClient: &http.Client{Timeout: 20 * time.Second},
	}
}

func (api *HttpCodeExecutor) Execute(in service.CodeExecRequest) (service.CodeExecResponse, error) {
	jsonData, err := json.Marshal(in)
	if err != nil {
		return service.CodeExecResponse{}, fmt.Errorf("failed to marshal request: %w", err)
	}

	req, err := http.NewRequest("POST", config.CodeExecutorHost+"/exec/python", bytes.NewBuffer(jsonData))
	if err != nil {
		return service.CodeExecResponse{}, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := api.httpClient.Do(req)
	if err != nil {
		return service.CodeExecResponse{}, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	var res service.CodeExecResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return service.CodeExecResponse{}, fmt.Errorf("failed to decode response: %w", err)
	}

	return res, nil
}
