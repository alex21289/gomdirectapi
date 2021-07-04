package gomdirectapi

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/alex21289/merkur"
)

type Client struct {
	http    merkur.Client
	session *ComdirectSession
}

func GetClient(s *ComdirectSession) (Client, error) {

	requestInfo := fmt.Sprintf("{\"clientRequestId\":{\"sessionId\":\"%s\",\"requestId\":\"%s\"}}", s.Session.SessionID, s.Session.RequestID)

	headers := make(http.Header)
	headers.Set("Accept", "application/json")
	headers.Set("Authorization", "Bearer "+s.Session.AccessToken)
	headers.Set("x-http-request-info", requestInfo)
	headers.Set("Content-Type", "application/json")
	headers.Set("Cookie", "qSession="+s.Session.QSession)

	clientBuilder := merkur.NewBuilder()
	httpClient := clientBuilder.SetHeaders(headers).Build()

	client := Client{
		http:    httpClient,
		session: s,
	}

	return client, nil
}

func (c *Client) Close() error {

	response, err := c.http.Delete(RevokeURL)
	if err != nil {
		return err
	}
	if response.StatusCode >= 400 {
		return errors.New(response.String() + " StatusCode: " + response.Status)
	}

	return nil
}
