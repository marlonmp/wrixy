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
)

func init() {

	err := client.Ping()

	if err != nil {
		panic(err)
	}

	userRepo = psql.UserRepo(client)
}

func TestInsertAndDeleteOne(t *testing.T) {

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

	t.Logf("inserted user: %#v\n", insertedUser)

	deletedUser, err := userRepo.DeleteOne(insertedUser.ID)

	if err != nil {
		t.Fatal("delete one err:", err)
	}

	if deletedUser.ID != insertedUser.ID {
		t.Error("the user id doesn't match")
	}

	t.Logf("deleted user: %#v\n", deletedUser)
}

func TestFindOne(t *testing.T) {
	insertedUser, _ := userRepo.InsertOne(user)

	f := port.UserFilter{
		ID:    insertedUser.ID,
		Limit: 10,
	}

	foundUser, err := userRepo.FindOne(f)

	if err != nil {
		t.Fatal("find one error:", err)
	}

	if foundUser.ID != insertedUser.ID {
		t.Error("the user id doesn't match")
	}

	userRepo.DeleteOne(insertedUser.ID)
}
