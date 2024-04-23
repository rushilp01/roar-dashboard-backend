package service

import (
	"fmt"
	"roar-dashboard-backend/repository"
	"roar-dashboard-backend/utils"
	"strings"
)

type Service struct {
	repo repository.Repo
}

func (ser *Service) GetFeatureRequestsPage(token string) error {
	authTokenArr := strings.Split(token, " ")
	authToken := authTokenArr[len(authTokenArr)-1]
	fmt.Println(authToken)
	userId, err := ser.repo.GetUserIdFromToken(authToken)
	if err != nil {
		return err
	}
	fmt.Println("//////////////////", userId)
	user_data := utils.GetUserDetailsfromId(userId)
	fmt.Println(user_data)
	return nil
}

func NewService() Service {
	return Service{
		repo: repository.NewRepo(),
	}
}
