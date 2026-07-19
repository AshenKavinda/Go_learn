package service

import (
	"context"

	"github.com/ashenkavinda/go_social_app/internel/dto/request"
	"github.com/ashenkavinda/go_social_app/internel/dto/response"
	appError "github.com/ashenkavinda/go_social_app/internel/error"
	"github.com/ashenkavinda/go_social_app/internel/models"
	repositoryInterfaces "github.com/ashenkavinda/go_social_app/internel/repository/interfaces"
)

type UserService struct {
	UserRepository     repositoryInterfaces.UserRepository
	FollowerRepository repositoryInterfaces.Follower_repository
}

func NewUserService(userRepository repositoryInterfaces.UserRepository, followerRepository repositoryInterfaces.Follower_repository) UserService {
	return UserService{UserRepository: userRepository, FollowerRepository: followerRepository}
}

func (s *UserService) Create(ctx context.Context, req *request.CreateUser) (*response.MessageResponce, error) {
	user := &models.User{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}

	if _, err := s.UserRepository.Create(ctx, user); err != nil {
		return nil, err
	}

	return &response.MessageResponce{Message: "user created successfully"}, nil
}

func (s *UserService) GetAll(ctx context.Context) (*[]models.User, error) {
	users, err := s.UserRepository.GetAll(ctx)
	if err != nil {
		return nil, appError.Internel(err)
	}

	return users, nil
}

func (s *UserService) GetByID(ctx context.Context, id int64) (*models.User, error) {
	user, err := s.UserRepository.GetByID(ctx, id)
	if err != nil {
		return nil, appError.NotFound("Record not found")
	}
	return user, nil
}

func (s *UserService) UpdateByID(ctx context.Context, id int64, data *request.UpdateUser) (*models.User, error) {
	user := &models.User{
		ID:       id,
		Username: data.Username,
		Password: data.Password,
	}

	ra, err := s.UserRepository.Update(ctx, user)
	if err != nil {
		return nil, appError.Internel(err)
	}

	if ra <= 0 {
		return nil, appError.NotFound("Record not found")
	}

	user, _ = s.GetByID(ctx, id)

	return user, nil
}

func (s *UserService) Delete(ctx context.Context, id int64) error {
	ra, err := s.UserRepository.Delete(ctx, id)
	if err != nil {
		return appError.Internel(err)
	}
	if ra <= 0 {
		return appError.BadRequest("Invalid id")
	}
	return nil
}

func (s *UserService) Feed(ctx context.Context, userID int64) (*[]models.Post, error) {
	posts, err := s.FollowerRepository.Feed(ctx, userID)
	if err != nil {
		return nil, appError.Internel(err)
	}

	return posts, nil
}

func (s *UserService) Follow(ctx context.Context, follower int64, following int64) (*response.MessageResponce, error) {
	if follower == following {
		return nil, appError.BadRequest("follower and following must be different")
	}

	if _, err := s.UserRepository.GetByID(ctx, follower); err != nil {
		return nil, appError.NotFound("follower user not found")
	}

	if _, err := s.UserRepository.GetByID(ctx, following); err != nil {
		return nil, appError.NotFound("following user not found")
	}

	if err := s.FollowerRepository.Follow(ctx, follower, following); err != nil {
		return nil, appError.Internel(err)
	}

	return &response.MessageResponce{Message: "followed successfully"}, nil
}

func (s *UserService) Unfollow(ctx context.Context, follower int64, following int64) (*response.MessageResponce, error) {
	if follower == following {
		return nil, appError.BadRequest("follower and following must be different")
	}

	if _, err := s.UserRepository.GetByID(ctx, follower); err != nil {
		return nil, appError.NotFound("follower user not found")
	}

	if _, err := s.UserRepository.GetByID(ctx, following); err != nil {
		return nil, appError.NotFound("following user not found")
	}

	if err := s.FollowerRepository.Unfollow(ctx, follower, following); err != nil {
		return nil, appError.Internel(err)
	}

	return &response.MessageResponce{Message: "unfollowed successfully"}, nil
}
