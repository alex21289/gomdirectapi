package gomdirectapi

import (
	"fmt"
	"net/http"

	"github.com/alex21289/merkur"
)

type Client struct {
	http    merkur.Client
	Session *comdirectSession
}

func GetClient(s *comdirectSession) (Client, error) {

	requestInfo := fmt.Sprintf("{\"clientRequestId\":{\"sessionId\":\"%s\",\"requestId\":\"%s\"}}", s.AuthClient.SessionID, s.AuthClient.RequestID)

	headers := make(http.Header)
	headers.Set("Accept", "application/json")
	headers.Set("Authorization", "Bearer "+s.AuthClient.AccessToken)
	headers.Set("x-http-request-info", requestInfo)
	headers.Set("Content-Type", "application/json")
	headers.Set("Cookie", "qSession="+s.AuthClient.QSession)

	clientBuilder := merkur.NewBuilder()
	httpClient := clientBuilder.SetHeaders(headers).Build()

	client := Client{
		http:    httpClient,
		Session: s,
	}

	return client, nil
}
