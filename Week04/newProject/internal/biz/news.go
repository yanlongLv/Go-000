package biz

import "github.com/alwayGo/newProject/internal/data"

// BaseClient is new  interface
type BaseClient interface {
	GetNew(newID string) (*data.News, error)
}

// Client is news
type client struct {
	cli data.DatabaseClient
}

//NewClient ..
func NewClient() BaseClient {
	return &client{}
}

// GetNew is return
func (c *client) GetNew(newID string) (*data.News, error) {
	return c.cli.GetNews(newID)
}