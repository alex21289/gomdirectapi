package gomdirectapi

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/alex21289/gomdirectapi/postbox"
	"github.com/alex21289/merkur"
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

	// Get All
	matches := pb.Paging.Matches
	params := merkur.NewParams()
	params.Add("paging-count", strconv.Itoa(matches))

	response, err = c.http.GetQuery(PostboxURL, params)
	if err = handleErr(*response, err); err != nil {
		return nil, err
	}

	if err = response.UnmarshalJson(&pb); err != nil {
		return nil, err
	}

	return &pb, nil
}

func (c *Client) GetDocumentBytes(documentID string) ([]byte, error) {
	url := fmt.Sprintf(PostboxDocumentURL, documentID)
	headers := make(http.Header)
	headers.Set("Accept", "application/pdf")
	response, err := c.http.Get(url, headers)
	if err = handleErr(*response, err); err != nil {
		return nil, err
	}

	return response.Bytes(), nil
}

func (c *Client) DownloadPostboxDocument(documentID string, downloadPath string) error {
	doc, err := c.GetPostboxDocument(documentID)
	if err != nil {
		return err
	}

	if doc.Mimetype != "application/pdf" {
		return errors.New("cant download none pdf")
	}

	log.Println("Download: " + doc.Name)
	fileName := strings.ReplaceAll(doc.Name, "/", "-") + ".pdf"

	url := fmt.Sprintf(PostboxDocumentURL, documentID)
	headers := make(http.Header)
	headers.Set("Accept", "application/pdf")
	response, err := c.http.Get(url, headers)
	if err = handleErr(*response, err); err != nil {
		return err
	}
	if err := ioutil.WriteFile(filepath.Join(downloadPath, fileName), response.Bytes(), 0644); err != nil {
		return err
	}

	return nil
}

func (c *Client) GetPostboxDocument(documentID string) (*postbox.Document, error) {
	pb, err := c.GetPostbox()
	if err != nil {
		return nil, err
	}

	var document postbox.Document
	for _, d := range pb.Values {
		if d.DocumentID == documentID {
			document = d
		}
	}

	return &document, nil
}
