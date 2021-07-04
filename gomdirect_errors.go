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
	Messages []map[string]string `json:"messages"`
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
		var errMsg string
		if err := response.UnmarshalJson(&restErr); err != nil {
			errMsg = fmt.Sprintf("%d %s: %s", response.StatusCode, http.StatusText(response.StatusCode), response.String())
		}
		if restErr.Error != "" {
			errMsg = fmt.Sprintf("%d %s: %s - %s", response.StatusCode, http.StatusText(response.StatusCode), restErr.Error, restErr.ErrorDescription)
		}
		if restErr.Summary != "" {
			errMsg = fmt.Sprintf("%d %s: %s", response.StatusCode, http.StatusText(response.StatusCode), restErr.Summary)
		}
		return errors.New(errMsg)
	}
	return nil
}
