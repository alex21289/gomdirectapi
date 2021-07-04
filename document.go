package gomdirectapi

import (
	"fmt"

	"github.com/alex21289/gomdirectapi/postbox"
)

func (c *Client) GetPostbox() (*postbox.PostboxApiResponse, error) {
	response, err := c.http.Get(PostboxURL)
	if err = handleErr(*response, err); err != nil {
		return nil, err
	}

	var pb postbox.PostboxApiResponse
	if err = response.UnmarshalJson(&pb); err != nil {
		return nil, err
	}

	return &pb, nil
}

func (c *Client) GetPostboxDocument(documentID string) (*postbox.Document, error) {
	url := fmt.Sprintf(PostboxDocumentURL, documentID)
	response, err := c.http.Get(url)
	if err = handleErr(*response, err); err != nil {
		return nil, err
	}

	var d postbox.Document
	if err = response.UnmarshalJson(&d); err != nil {
		return nil, err
	}

	return &d, nil
}
