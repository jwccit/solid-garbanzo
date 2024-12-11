package services

import "flyte/goapi/internal/api/models"

func GetUsers() ([]models.User, error) {
	var users []models.User
	users = append(users, models.User{ID: 1, Name: "John Doe", Email: "john.doe@example.com"})
	users = append(users, models.User{ID: 2, Name: "Jane Doe", Email: "jane.doe@example.com"})

	return users, nil
}
