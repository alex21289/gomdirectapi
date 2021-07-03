package gomdirectapi

import (
	"fmt"
	"net/http"

	"github.com/alex21289/merkur"
)

type Client struct {
	http merkur.Client
}

func GetClient(session *comdirectSession) (*Client, error) {

	requestInfo := fmt.Sprintf("{\"clientRequestId\":{\"sessionId\":\"%s\",\"requestId\":\"%s\"}}", session.AuthClient.SessionID, session.AuthClient.RequestID)

	headers := make(http.Header)
	headers.Set("Accept", "application/json")
	headers.Set("Authorization", "Bearer "+session.Authentication.AccessToken)
	headers.Set("x-http-request-info", requestInfo)
	headers.Set("Content-Type", "application/json")
	headers.Set("Cookie", "qSession="+session.AuthClient.QSession)

	clientBuilder := merkur.NewBuilder()
	httpClient := clientBuilder.SetHeaders(headers).Build()

	client := Client{
		http: httpClient,
	}

	return &client, nil
}
