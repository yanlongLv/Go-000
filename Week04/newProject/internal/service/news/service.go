package service

import (
	"context"

	"github.com/pkg/errors"

	pb "github.com/alwayGo/newProject/api/v1/news"
	"github.com/alwayGo/newProject/internal/biz"
)

//NewsService is a new service
type newsService struct {
	b biz.BaseClient
	pb.UnimplementedNewsServer
}

func NewserviceClient(b biz.BaseClient) *newsService {
	return &newsService{b: b}
}
func (s *newsService) GetNews(ctx context.Context, in *pb.NewsRequest) (*pb.NewsReply, error){
	news, err := s.b.GetNew(in.NewId)
	if err != nil {
		return nil, errors.Wrap(err, "read failed")
	}
	return &pb.NewsReply{NewId: news.NewID, Title: news.Title, Context: news.Context}, nil
}
