package users_transport_http

import "github.com/Vadick-do/todo_app/internal/core/domain"

type UserDTOResponce struct {
	ID          int     `json:"id"`
	Version     int     `json:"version"`
	FullName    string  `json:"full_name"`
	PhoneNumber *string `json:"phone_number"`
}

func userDTOFromDomain(user domain.User) UserDTOResponce {
	return UserDTOResponce{
		ID:          user.ID,
		Version:     user.Version,
		FullName:    user.FullName,
		PhoneNumber: user.PhoneNumber,
	}
}

func usersDTOFromDomains(users []domain.User) []UserDTOResponce {
	userDTO := make([]UserDTOResponce, len(users))

	for i, user := range users {
		userDTO[i] = userDTOFromDomain(user)
	}

	return userDTO
}
