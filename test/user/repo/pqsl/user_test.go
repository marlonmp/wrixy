package pqsl_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/marlonmp/wrixy/internal/apps/user/domain"
	"github.com/marlonmp/wrixy/internal/apps/user/port"
	"github.com/marlonmp/wrixy/internal/apps/user/repo/psql"
)

var (
	userRepo port.UserRepo

	user = domain.User{
		Nickname: "Manolo Gz!!",
		Username: "manolo214",
		Email:    "manolo214@gmial.com",
		Password: "...",
	}

	users = []domain.User{
		user,
		{
			Nickname: "Camila",
			Username: "cami2243",
			Email:    "cami2243@gmail.com",
			Password: "...",
		},
		{
			Nickname: "Juancho",
			Username: "jp45",
			Email:    "jp45@gmail.com",
			Password: "...",
		},
	}
)

func init() {

	err := client.Ping()

	if err != nil {
		panic(err)
	}

	userRepo = psql.UserRepo(client)
}

func clearDB() {
	client.Exec(`delete from "user";`)
}

func TestInsertOne(t *testing.T) {

	insertedUser, err := userRepo.InsertOne(user)

	if err != nil {
		t.Fatal("insert one err:", err)
	}

	if insertedUser.ID == uuid.Nil {
		t.Error("the user uuid is nil")
	}
	if insertedUser.Username != user.Username {
		t.Error("the username is different")
	}
	if insertedUser.Nickname != user.Nickname {
		t.Error("the nickname is different")
	}
	if insertedUser.Email != user.Email {
		t.Error("the email is different")
	}
	if insertedUser.CreatedAt.IsZero() {
		t.Error("the created_at is invalid")
	}

	clearDB()
}

func TestFindOne(t *testing.T) {
	insertedUser, _ := userRepo.InsertOne(user)

	f := port.UserFilter{
		ID:    insertedUser.ID,
		Limit: 10,
	}

	userFound, err := userRepo.FindOne(f)

	if err != nil {
		t.Fatal("find one error:", err)
	}

	if userFound.ID != insertedUser.ID {
		t.Error("the user id doesn't match")
	}

	clearDB()
}

func TestFind(t *testing.T) {

	insertedUsers := make([]domain.User, 0)

	for _, user := range users {

		u, _ := userRepo.InsertOne(user)

		insertedUsers = append(insertedUsers, u)
	}

	f := port.UserFilter{
		Limit: 100,
	}

	usersFound, err := userRepo.Find(f)

	if err != nil {
		t.Fatal("find err:", err)
	}

	if len(insertedUsers) != len(users) {
		t.Error("invalid inserted users")
	}

	if len(insertedUsers) != len(usersFound) {
		t.Error("invalid users found len")
	}

	userInList := make([]domain.User, 0)

userFoundedFor:
	for _, userFound := range usersFound {

		for _, insertedUser := range insertedUsers {

			if insertedUser.ID == userFound.ID {
				userInList = append(userInList, userFound)
				continue userFoundedFor
			}
		}
	}

	if len(userInList) != len(insertedUsers) {
		t.Error("different users found")
	}

	clearDB()
}

func TestUpdateOne(t *testing.T) {
	insertedUser, _ := userRepo.InsertOne(user)

	newNickname := user.Nickname + `_`

	userUpdate := domain.User{
		Nickname: newNickname,
	}

	updatedUser, err := userRepo.UpdateOne(insertedUser.ID, userUpdate)

	if err != nil {
		t.Fatal("update one err:", err)
	}

	if updatedUser.Nickname != newNickname {
		t.Error("different user nickname")
	}

	if updatedUser.UpdatedAt == nil || updatedUser.UpdatedAt.IsZero() {
		t.Error("invalid updated_at")
	}

	clearDB()
}

func TestDeleteOne(t *testing.T) {

	insertedUser, _ := userRepo.InsertOne(user)

	deletedUser, err := userRepo.DeleteOne(insertedUser.ID)

	if err != nil {
		t.Fatal("delete one err:", err)
	}

	if deletedUser.ID != insertedUser.ID {
		t.Error("the user id doesn't match")
	}

	if deletedUser.DeletedAt == nil || deletedUser.DeletedAt.IsZero() {
		t.Error("invalid deleted_at")
	}

	clearDB()
}
