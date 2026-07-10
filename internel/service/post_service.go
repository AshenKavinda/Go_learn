package service

import (
	"context"

	"github.com/ashenkavinda/go_social_app/internel/dto/request"
	"github.com/ashenkavinda/go_social_app/internel/dto/response"
	appError "github.com/ashenkavinda/go_social_app/internel/error"
	"github.com/ashenkavinda/go_social_app/internel/models"
	repositoryInterfaces "github.com/ashenkavinda/go_social_app/internel/repository/interfaces"
)

type PostService struct {
	PostRepository repositoryInterfaces.PostRepository
}

func NewPostService(repository repositoryInterfaces.PostRepository) PostService {
	return PostService{PostRepository: repository}
}

func (s *PostService) Create(ctx context.Context, req *request.CreatePost) (*response.MessageResponce, error) {
	post := &models.Post{
		Content: req.Content,
		Title:   req.Title,
		Tags:    req.Tags,
		UserID:  req.UserID,
	}

	if _, err := s.PostRepository.Create(ctx, post); err != nil {
		return nil, err
	}

	return &response.MessageResponce{Message: "user created successfuly"}, nil

}

func (s *PostService) GetAll(ctx context.Context) (*[]models.Post, error) {
	posts, err := s.PostRepository.GetAll(ctx)
	if err != nil {
		return nil, appError.Internel(err)
	}

	return posts, nil
}

func (s *PostService) GetByID(ctx context.Context, id int64) (*models.Post, error) {
	post, err := s.PostRepository.GetByID(ctx, id)
	if err != nil {
		return nil, appError.NotFound("Record Not found.")
	}
	return post, nil
}

func (s *PostService) UpdateByID(ctx context.Context, id int64, data *request.UpdatePost) (*models.Post, error) {
	post := &models.Post{
		ID:      id,
		Title:   data.Title,
		Content: data.Content,
		Tags:    data.Tags,
	}

	ra, err := s.PostRepository.Update(ctx, post)
	if err != nil {
		return nil, appError.Internel(err)
	}

	if ra <= 0 {
		return nil, appError.NotFound("Record Not found.")
	}

	post, _ = s.GetByID(ctx, id)

	return post, nil

}

func (s *PostService) Delete(ctx context.Context, id int64) error {
	ra, err := s.PostRepository.Delete(ctx, id)
	if err != nil {
		return appError.Internel(err)
	}
	if ra <= 0 {
		return appError.BadRequest("Invalied id")
	}
	return nil
}
