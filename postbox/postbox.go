package postbox

type PostboxApiResponse struct {
	Paging Paging     `json:"paging"`
	Values []Document `json:"values"`
}
type Paging struct {
	Index   int `json:"index"`
	Matches int `json:"matches"`
}
type Documentmetadata struct {
	Archived          bool `json:"archived"`
	Alreadyread       bool `json:"alreadyRead"`
	Predocumentexists bool `json:"predocumentExists"`
}
type Document struct {
	Documentid       string           `json:"documentId"`
	Name             string           `json:"name"`
	Datecreation     string           `json:"dateCreation"`
	Mimetype         string           `json:"mimeType"`
	Deletable        bool             `json:"deletable"`
	Advertisement    bool             `json:"advertisement"`
	Documentmetadata Documentmetadata `json:"documentMetaData"`
}
