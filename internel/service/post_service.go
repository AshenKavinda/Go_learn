package service

import (
	"context"

	"github.com/ashenkavinda/go_social_app/internel/dto/request"
	"github.com/ashenkavinda/go_social_app/internel/dto/response"
	"github.com/ashenkavinda/go_social_app/internel/models"
	repositoryInterfaces "github.com/ashenkavinda/go_social_app/internel/repository/interfaces"
)

type PostService struct {
	PostRepository repositoryInterfaces.PostRepository
}

func NewPostService(repository repositoryInterfaces.PostRepository) PostService {
	return PostService{PostRepository: repository}
}

func (s *PostService) Create(ctx context.Context, req *request.PostRequest) (*response.MessageResponce, error) {
	post := &models.Post{
		Content: req.Content,
		Title:   req.Title,
		Tags:    req.Tags,
		UserID:  req.UserID,
	}

	if err := s.PostRepository.Create(ctx, post); err != nil {
		return nil, err
	}

	return &response.MessageResponce{Message: "user created successfuly"}, nil

}
