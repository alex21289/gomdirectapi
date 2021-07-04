package gomdirectapi

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/alex21289/merkur/mcore"
)

type sessionError struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
	Summary          string `json:"summary"`
}

type validationError struct {
	Code     string              `json:"code"`
	Messages []validationMessage `json:"messages"`
}

type validationMessage struct {
	Serverity string            `json:"serverity"`
	Key       string            `json:"key"`
	Message   string            `json:"message"`
	Args      map[string]string `json:"args"`
	Origin    []string          `json:"origin"`
}

type comError struct {
	Status     string
	StatusCode int
	Error      string
	Message    string
}

// Obsolet
func newComRestError(statusCode int, err sessionError) *comError {
	return &comError{
		Status:     http.StatusText(statusCode),
		StatusCode: statusCode,
		Error:      err.Error,
		Message:    err.ErrorDescription,
	}
}

// Handles error and  response error
func handleErr(response mcore.Response, err error) error {
	if err != nil {
		return err
	}
	if response.StatusCode >= 400 {
		var restErr sessionError
		var validationErr validationError
		var errMsg string
		if err := response.UnmarshalJson(&restErr); err != nil {
			errMsg = fmt.Sprintf("%d %s: %s", response.StatusCode, http.StatusText(response.StatusCode), response.String())
		} else if restErr.Error != "" {
			errMsg = fmt.Sprintf("%d %s: %s - %s", response.StatusCode, http.StatusText(response.StatusCode), restErr.Error, restErr.ErrorDescription)
		} else if restErr.Summary != "" {
			errMsg = fmt.Sprintf("%d %s: %s", response.StatusCode, http.StatusText(response.StatusCode), restErr.Summary)
		} else if err := response.UnmarshalJson(&validationErr); err == nil {
			errMsg = fmt.Sprintf("%d %s: %s - %s", response.StatusCode, http.StatusText(response.StatusCode), validationErr.Code, validationErr.Messages[0].Message)
		} else {
			errMsg = fmt.Sprintf("%d %s: %s", response.StatusCode, http.StatusText(response.StatusCode), response.String())
		}
		return errors.New(errMsg)
	}
	return nil
}
