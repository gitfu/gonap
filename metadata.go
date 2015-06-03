package gonap

// MetaData is a struct for metadata returned in a PBResp.Body
type MetaData struct {
	LastModified   string `json:"lastModifiedDate,omitempty"`
	LastModifiedBy string `json:"lastModifiedBy,omitempty"`
	Created        string `json:"createdDate,omitempty"` //"2014-02-01T11:12:12Z",
	CreatedBy      string `json:"createdBy,omitempty"`
	State          string `json:"state,omitempty"`
	Etag           string `json:"etag,omitempty"`
}
