package users_transport_http

import "github.com/Vadick-do/todo_app/internal/core/domain"

type UserDTOResponce struct {
	ID          int     `json:"id"            example:"10"`
	Version     int     `json:"version"       example:"3"`
	FullName    string  `json:"full_name"     example:"Ivan Ivanov"`
	PhoneNumber *string `json:"phone_number"  example:"+79998887766"`
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
