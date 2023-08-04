package services

import (
	. "clokify/models"
	"errors"
	"fmt"
	"log"
	"net/mail"
	"sync"

	. "clokify/utils"

	. "clokify/types"
)

type LoginType struct {
	Email    string
	Password string
}

type UserResult struct {
	err  error
	user User
}

type UserServiceManager struct {
	*ServiceManager
}

func (us *UserServiceManager) IsEmailExists(email string) (bool, error) {
	var user = &User{}
	res := us.Db.Model(user).Where("email = ?", email).First(user)
	if res.Error == nil {
		return true, nil
	}

	return false, errors.New("Email didnt find or have error finding")
}

func (us *UserServiceManager) CreateUser(user *User) (error, bool) {
	hashedPassword, err := GenerateHash(user.Password)
	if err == nil {
		user.Password = hashedPassword
	} else {
		return errors.New("Didnt able to hash properly"), false
	}

	if user.Email == "" || user.Password == "" {
		return errors.New("Email and Password must not be null"), false
	}

	if _, err := mail.ParseAddress(user.Email); err != nil {
		return errors.New("Email not valid"), false
	}

	if _, err := us.IsEmailExists(user.Email); err != nil {

		res := us.Db.Model(&User{}).Create(user)
		if res.Error != nil {
			return errors.New("User not saved"), false
		}

		log.Println("user saved")

		return nil, true
	} else {
		return errors.New("Email already exist"), false
	}
}

func (us *UserServiceManager) LoginUser(loginData *LoginType) (error, string, bool) {
	if loginData.Email == "" || loginData.Password == "" {
		return errors.New("Email and Password must not be null"), "", false
	}

	var user = User{}
	res := us.Db.Model(&User{}).Where("email = ?", loginData.Email).Select("email").First(&user)

	if res.Error != nil {
		return errors.New("User with that email not found"), "", false
	}

	token, err := GenerateJWTToken(user.Email)
	if err != nil {
		return err, "", false
	}

	return nil, token, false
}

func (s *UserServiceManager) DeleteUser(email string, srv *ServiceManager) (error, User) {

	taskService := &TaskServiceManager{
		ServiceManager: srv,
	}

	// getting the user if not exist
	userError, userData := s.GetUserByEmail(email)
	if userError != nil {
		return errors.New("didnt get user"), User{}
	}

	// deleting associated user_id rows from project_users table
	projectUserResult := s.Db.Raw("SELECT user_id FROM project_users user_id = ?", userData.ID)
	if projectUserResult.Error != nil {
		fmt.Println("Error executing raw query:", projectUserResult.Error)
	} else {
		result := s.Db.Exec("DELETE FROM project_users WHERE user_id = ?", userData.ID)
		if result.Error != nil {
			fmt.Println("Error executing delete raw query:", result.Error)
		}
	}

	// deleting associated task table
	taskErr, taskData := taskService.GetTask(1)
	if taskErr != nil {
		fmt.Print("User didnt assiciate with any task")
	} else {
		var task Task
		if err := s.Db.Model(&task).Where("id = ?", taskData.ID).Delete(&task); err.Error != nil {
			return errors.New("delete failed for task associated with user"), User{}
		}
	}

	// deleting user
	var user User
	if err := s.Db.Model(&user).Where("email = ?", email).Delete(&user); err.Error != nil {
		return errors.New("delete failed"), User{}
	}

	return nil, user
}

func (s *UserServiceManager) GetUser(id int) (error, User) {
	var user User
	if err := s.Db.Model(&user).Preload("Projects").Where("id = ?", id).First(&user); err.Error != nil {
		return errors.New("didnt find user"), User{}
	}

	return nil, user
}

func (s *UserServiceManager) GetUserWithChannel(
	id int,
	ch chan UserResult,
	wg *sync.WaitGroup,
) {
	defer wg.Done()
	userResult := UserResult{}
	userResult.err, userResult.user = s.GetUser(id)
	ch <- userResult
	return
}

func (s *UserServiceManager) GetUserByEmail(email string) (error, User) {
	var user User
	if err := s.Db.Model(&user).Preload("Projects").Where("email = ?", email).First(&user); err.Error != nil {
		return errors.New("didnt find user"), User{}
	}

	return nil, user
}
