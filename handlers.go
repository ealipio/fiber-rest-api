package main

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

func createUser(c *fiber.Ctx) error {
	// add new user to users
	var userFromBody userType
	err := c.BodyParser(&userFromBody)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("bad request")
	}
	// todo ok
	users = append(users, userFromBody)
	return c.JSON(userFromBody)
}
func getAllUsers(c *fiber.Ctx) error {
	return c.JSON(users)
}
func getUser(c *fiber.Ctx) error {
	userId := c.Params("userId")
	userFound, _, err := findUserById(userId)
	if err != nil {
		return c.JSON(err.Error())
	}
	return c.JSON(userFound)
}

func update(c *fiber.Ctx) error {
	// PUT /users/:userId == > /users/3
	userId := c.Params("userId")
	var userFromBody userType

	_, index, err := findUserById(userId)

	if err != nil {
		return c.JSON(err.Error())
	}
	// todo ok
	// parse data
	err = c.BodyParser(&userFromBody)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	// data parsed
	// replace value from users
	//first: remove from users, second add to users

	// removing...
	users = append(users[:index], users[index+1:]...)
	// add new value
	users = append(users, userFromBody)

	return c.SendString("user updated completely")
}

func toggleAdmin(c *fiber.Ctx) error {
	userId := c.Params("userId")
	userFound, _, err := findUserById(userId)

	if err != nil {
		return c.JSON(err.Error())
	}
	// todo ok
	userFound.IsAdmin = !userFound.IsAdmin
	return c.SendString("is admin updated")
}

func delete(c *fiber.Ctx) error {
	userId := c.Params("userId")
	_, index, err := findUserById(userId)

	if err != nil {
		return c.JSON(err.Error())
	}
	// remove with slice operator
	users = append(users[:index], users[index+1:]...)

	return c.SendString("user removed")
}

// utility

func findUserById(userId string) (*userType, int, error) {
	// loop over users
	for index, user := range users {
		if user.Id == userId {
			return &users[index], index, nil
		}
	}
	return nil, -1, errors.New("user not found")
}
