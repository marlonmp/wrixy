// TODO: add error handlers for constraints, not found and unknown error
package psql

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/marlonmp/wrixy/internal/apps/user/domain"
	"github.com/marlonmp/wrixy/internal/apps/user/port"
	"github.com/nullism/bqb"
)

const (
	// queries
	userInsertOneQuery = `insert into "user" ("nickname", "username", "email", "password") values ($1, $2, $3, $4) returning "uuid", "nickname", "username", "email", "created_at";`
	userDeleteOneQuery = `update "user" set "deleted_at" = now() where "uuid" = $1 and "deleted_at" is null returning "uuid", "nickname", "username", "email", "deleted_at";`
)

type userRepo struct {
	client *sql.DB
}

func UserRepo(client *sql.DB) port.UserRepo {
	return userRepo{client}
}

func (ur userRepo) Find(f port.UserFilter) ([]domain.User, error) {
	ctx := context.Background()

	conn, err := ur.client.Conn(ctx)

	if err != nil {
		return nil, err
	}

	defer conn.Close()

	q := bqb.Optional(`select "uuid", "nickname", "username" from "user" where`)

	if len(f.Nickname) > 0 {
		q.And(`"nickname" ilike '%' || ? || '%'`, f.Nickname)
	}

	if len(f.Username) > 0 {
		q.And(`"username" ilike '%' || ? || '%'`, f.Username)
	}

	if !f.IsDeleted {
		q.And(`"deleted_at" is null`)
	}

	q.Space(`limit ? offset ?;`, f.Limit, f.Offset)

	query, args, _ := q.ToPgsql()

	row, err := conn.QueryContext(ctx, query, args...)

	if err != nil {
		return nil, err
	}

	defer row.Close()

	users := make([]domain.User, 0)

	for row.Next() {
		u := domain.User{}

		err = row.Scan(&u.ID, &u.Nickname, &u.Username)

		if err != nil {
			return nil, err
		}

		users = append(users, u)
	}

	return users, nil
}

func (ur userRepo) FindOne(f port.UserFilter) (u domain.User, err error) {

	q := bqb.Optional(`select "uuid", "nickname", "username", "email" from "user" where`)

	if f.ID != uuid.Nil {
		q.And(`"uuid" = ?`, f.ID)
	}

	if !f.IsDeleted {
		q.And(`"deleted_at" is null`)
	}

	if len(f.Email) > 0 {
		q.And(`"email" = ?`, f.Email)
	}

	if len(f.Username) > 0 {
		q.And(`"username" = ?`, f.Username)
	}

	q.Space(`limit 1;`)

	query, args, _ := q.ToPgsql()

	err = ur.userQueryRowScanner(query, args,
		// Scan
		&u.ID, &u.Nickname, &u.Username, &u.Email,
	)

	return
}

func (ur userRepo) InsertOne(user domain.User) (u domain.User, err error) {

	args := []any{user.Nickname, user.Username, user.Email, user.Password}

	err = ur.userQueryRowScanner(userInsertOneQuery, args,
		// Scan
		&u.ID, &u.Nickname, &u.Username, &u.Email, &u.CreatedAt,
	)

	return
}

func (ur userRepo) UpdateOne(id uuid.UUID, user domain.User) (u domain.User, err error) {

	q := bqb.Optional(`update "user" set "updated_at" = now()`)

	if len(user.Nickname) > 0 {
		q.Comma(`"nickname" = ?`, user.Nickname)
	}

	if len(user.Username) > 0 {
		q.Comma(`"username" = ?`, user.Username)
	}

	if len(user.Password) > 0 {
		q.Comma(`"password" = ?`, user.Password)
	}

	q.Space(`where "uuid" = ? and "deleted_at" is null`, id)

	q.Space(`returning "uuid", "nickname", "username", "email", "updated_at";`)

	query, args, _ := q.ToPgsql()

	err = ur.userQueryRowScanner(query, args,
		// Scan
		&u.ID, &u.Nickname, &u.Email, &u.Email, &u.UpdatedAt,
	)

	return
}

func (ur userRepo) DeleteOne(id uuid.UUID) (u domain.User, err error) {
	args := []any{id}

	err = ur.userQueryRowScanner(userDeleteOneQuery, args,
		// Scan
		&u.ID, &u.Nickname, &u.Username, &u.Email, &u.DeletedAt,
	)

	return
}

func (ur userRepo) userQueryRowScanner(query string, args []any, dest ...any) (err error) {
	ctx := context.Background()

	conn, err := ur.client.Conn(ctx)

	if err != nil {
		return
	}

	defer conn.Close()

	row := conn.QueryRowContext(ctx, query, args...)

	if err = row.Err(); err != nil {
		return
	}

	err = row.Scan(dest...)

	return
}
