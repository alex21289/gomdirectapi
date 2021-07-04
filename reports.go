package gomdirectapi

import "github.com/alex21289/gomdirectapi/report"

func (c *Client) GetReport() (*report.ReportApiResponse, error) {
	response, err := c.http.Get(ReportURL)
	if err = handleErr(*response, err); err != nil {
		return nil, err
	}

	var r report.ReportApiResponse
	if err = response.UnmarshalJson(&r); err != nil {
		return nil, err
	}

	return &r, nil
}
