package service

import (
	"errors"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"synapsis-challenge/internal/consts"
	"synapsis-challenge/internal/entities"
	"synapsis-challenge/internal/middlewares/jwt"
	"synapsis-challenge/internal/repositories"
)

// Service is an interface from which api module can access our repository of all our models
type AuthService interface {
	Login(user *entities.User) (userData *entities.User, token *string, err error)
	Register(user *entities.User) (userData *entities.User, token *string, err error)
}

type authService struct {
	jwtMdwr  jwt.AuthMiddleware
	userRepo repositories.UserRepository
}

// NewService is used to create a single instance of the service
func NewAuthService(jwtMdwr jwt.AuthMiddleware, userRepo repositories.UserRepository) AuthService {
	return &authService{
		jwtMdwr:  jwtMdwr,
		userRepo: userRepo,
	}
}

func (s *authService) Login(user *entities.User) (*entities.User, *string, error) {
	userData, err := s.userRepo.FindOne(user.Username)
	if err != nil {
		return nil, nil, errors.New(consts.InternalServerError)
	}
	if userData == nil {
		return nil, nil, errors.New("username or password is wrong!")
	}

	err = bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(user.Password))
	if err != nil {
		return nil, nil, errors.New("username or password is wrong!")
	}

	token, err := s.jwtMdwr.GenerateToken(userData)

	if err != nil {
		return nil, nil, errors.New(consts.InternalServerError)
	}

	return userData, token, nil
}

func (s *authService) Register(user *entities.User) (*entities.User, *string, error) {

	checkUser, err := s.userRepo.FindOne(user.Username)

	if err != nil {
		return nil, nil, errors.New(consts.InternalServerError)
	}

	if checkUser != nil {
		return nil, nil, errors.New("username already taken!")
	}

	saltByte, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)

	if err != nil {
		return nil, nil, errors.New(consts.InternalServerError)
	}

	//set to salted password
	user.Password = string(saltByte)

	user.ID = uuid.New().String()

	err = s.userRepo.CreateUser(user)

	if err != nil {
		return nil, nil, errors.New(consts.InternalServerError)
	}

	newUser, err := s.userRepo.FindOne(user.Username)

	if err != nil {
		return nil, nil, errors.New(consts.InternalServerError)
	}

	token, err := s.jwtMdwr.GenerateToken(newUser)

	if err != nil {
		return nil, nil, errors.New(consts.InternalServerError)
	}

	return newUser, token, nil
}
