package db

import (
	"../types"
	"../utils"
	"github.com/asdine/storm/q"
	"time"
)

func PutUserInDB(user *types.User) (string, error) {
	var err error
	if user.Id == "" {
		user.Id = utils.GetNewUUID()
	}
	if user.Username == "" {
		user.Username = utils.GetNewUUID()
	}
	user.CreatedAt = time.Now()
	err = MasterDB["users"].Save(user)
	return user.Id, err
}

func GetUserById(id string) (*types.User, error) {
	user := new(types.User)
	err := MasterDB["users"].One("Id", id, user)
	return user, err
}

func GetAllUsers() ([]types.User, error) {
	var users []types.User
	err := MasterDB["users"].All(&users)
	return users, err
}

func GetUsersByGroup(group string) ([]types.User, error) {
	var users []types.User
	err := MasterDB["users"].Find("Group", group, &users)
	return users, err
}


func GetUserByEmail(email string) (*types.User, error) {
	user := new(types.User)
	err := MasterDB["users"].One("Email", email, user)
	return user, err
}

func GetUserByUsername(username string) (*types.User, error) {
	user := new(types.User)
	err := MasterDB["users"].One("Username", username, user)
	return user, err
}

func GetUsersByGroup(group string) ([]types.User, error) {
	var users []types.User
	err := MasterDB["users"].Find("Group", group, &users)
	return users, err
}


func DeleteUserById(id string) (*types.User, error) {
	user, err := GetUserById(id)
	if err != nil {
		return user, err
	}
	err = MasterDB["users"].DeleteStruct(user)
	return user, err

}

func DeleteAllUsers() error {
	query := MasterDB["users"].Select(q.True())
	err := query.Delete(new(types.User))
	return err
}

func UpdateUserById(userMod *types.User) error {
	err := MasterDB["users"].Update(userMod)
	return err
}
