package data

// News struct
type News struct {
	NewID   string `json: 'newId'`
	Title   string `json: 'title'`
	Context string `json: 'context'`
}

type databaseClient struct {
}

//NewDataBaseClient ...
func NewDataBaseClient() DatabaseClient {
	return &databaseClient{}
}

// DatabaseClient install
type DatabaseClient interface {
	GetNews(newID string) (*News, error)
}

// GetNew get new by newId
func (d *databaseClient) GetNews(newID string) (*News, error) {
	return &News{NewID: "123", Title: "kjfdk", Context: "nmnmnmnmnm"}, nil
}
