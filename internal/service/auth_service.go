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
	//find the user
	userData, err := s.userRepo.FindOne(user.Username)
	if err != nil {
		return nil, nil, errors.New(consts.InternalServerError)
	}

	//if user not in db then throw error
	if userData == nil {
		return nil, nil, errors.New("username or password is wrong!")
	}

	//compare password from request and form db if error means password is not matched
	err = bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(user.Password))
	if err != nil {
		return nil, nil, errors.New("username or password is wrong!")
	}

	//generate token based from user that already fetch
	token, err := s.jwtMdwr.GenerateToken(userData)

	if err != nil {
		return nil, nil, errors.New(consts.InternalServerError)
	}

	return userData, token, nil
}

func (s *authService) Register(user *entities.User) (*entities.User, *string, error) {

	//checking user
	checkUser, err := s.userRepo.FindOne(user.Username)

	if err != nil {
		return nil, nil, errors.New(consts.InternalServerError)
	}

	//if user already in db that means username already been taken
	if checkUser != nil {
		return nil, nil, errors.New("username already taken!")
	}

	//generate salt password
	saltByte, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)

	if err != nil {
		return nil, nil, errors.New(consts.InternalServerError)
	}

	//set to salted password
	user.Password = string(saltByte)

	//set id for new user
	user.ID = uuid.New().String()

	//save new user
	err = s.userRepo.CreateUser(user)

	if err != nil {
		return nil, nil, errors.New(consts.InternalServerError)
	}

	//get one for the response
	newUser, err := s.userRepo.FindOne(user.Username)

	if err != nil {
		return nil, nil, errors.New(consts.InternalServerError)
	}

	//generate token based from user that already fetch
	token, err := s.jwtMdwr.GenerateToken(newUser)

	if err != nil {
		return nil, nil, errors.New(consts.InternalServerError)
	}

	return newUser, token, nil
}
