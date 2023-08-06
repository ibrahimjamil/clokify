package cruds

import (
	. "clokify/models"
	. "clokify/services"
	"log"

	. "clokify/types"

	"gorm.io/gorm"
)

func UserCrud(db *gorm.DB, srvMan *ServiceManager) {
	userService := &UserServiceManager{
		ServiceManager: srvMan,
	}

	// create user manually
	user := &User{
		ID:       1,
		Name:     "Ibrahim",
		Email:    "ibrahimjamil090@gmail.com",
		Password: "Admin@123",
		Type:     "DEVELOPER",
	}
	err, userRes := userService.CreateUser(user)
	if err == nil {
		log.Println("user created successfully", userRes)
	} else {
		log.Println(err)
	}

	// login user to get token
	userData := &LoginType{
		Email:    "ibrahimjamil090@gmail.com",
		Password: "Admin@123",
	}
	err, loginData, bool := userService.LoginUser(userData)
	if err == nil {
		log.Println("user loged in successfully", loginData)
	} else {
		log.Println(err, loginData, bool)
	}

	// get User
	err, getUser := userService.GetUser(user.ID)
	if err == nil {
		log.Println("user fetched successfully", getUser)
	} else {
		log.Println(err, getUser)
	}

	// delete user
	// err, deletedUser := userService.DeleteUser("ibrahimjamil090@gmail.com", srvMan)
	// if err == nil {
	// 	log.Println("user delete successfully", deletedUser)
	// } else {
	// 	log.Println(err, deletedUser)
	// }
}
