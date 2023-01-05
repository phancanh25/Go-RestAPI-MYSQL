package service

import (
	"context"
	"database/sql"
	"fmt"
	q "github.com/core-go/sql"
	"reflect"
	"strings"

	. "go-service/internal/filter"
	. "go-service/internal/model"
)

type UserService interface {
	All(ctx context.Context) ([]User, error)
	Load(ctx context.Context, id string) (*User, error)
	Insert(ctx context.Context, user *User) (int64, error)
	Update(ctx context.Context, user *User) (int64, error)
	Patch(ctx context.Context, user map[string]interface{}) (int64, error)
	Delete(ctx context.Context, id string) (int64, error)
	Search(ctx context.Context, filter UserFilter) (*Result, error)
}

type userService struct {
	DB *sql.DB
	BuildParam func(int) string
}

func NewUserService(db *sql.DB) UserService {
	buildParam := q.GetBuild(db)
	return &userService{DB: db, BuildParam: buildParam}
}

func (s *userService) All(ctx context.Context) ([]User, error) {
	query := "select id, username, email, phone, date_of_birth from users"
	rows, err := s.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []User
	for rows.Next() {
		var user User
		err = rows.Scan(&user.Id, &user.Username, &user.Phone, &user.Email, &user.DateOfBirth)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (s *userService) Load(ctx context.Context, id string) (*User, error) {
	query := "select id, username, email, phone, date_of_birth from users where id = ? limit 1"
	rows, err := s.DB.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var user User
		err = rows.Scan(&user.Id, &user.Username, &user.Phone, &user.Email, &user.DateOfBirth)
		if err != nil {
			return nil, err
		}
		return &user, nil
	}
	return nil, nil
}

func (s *userService) Insert(ctx context.Context, user *User) (int64, error) {
	query := "insert into users (id, username, email, phone, date_of_birth) values (?, ?, ?, ?, ?)"
	stmt, er0 := s.DB.Prepare(query)
	if er0 != nil {
		return -1, nil
	}
	res, er1 := stmt.ExecContext(ctx, user.Id, user.Username, user.Email, user.Phone, user.DateOfBirth)
	if er1 != nil {
		return -1, nil
	}
	return res.RowsAffected()
}

func (s *userService) Update(ctx context.Context, user *User) (int64, error) {
	query := "update users set username = ?, email = ?, phone = ?, date_of_birth = ? where id = ?"
	stmt, er0 := s.DB.Prepare(query)
	if er0 != nil {
		return -1, nil
	}
	res, er1 := stmt.ExecContext(ctx, user.Username, user.Email, user.Phone, user.DateOfBirth, user.Id)
	if er1 != nil {
		return -1, er1
	}
	return res.RowsAffected()
}

func (s *userService) Patch(ctx context.Context, user map[string]interface{}) (int64, error) {
	userType := reflect.TypeOf(User{})
	jsonColumnMap := q.MakeJsonColumnMap(userType)
	colMap := q.JSONToColumns(user, jsonColumnMap)
	keys, _ := q.FindPrimaryKeys(userType)
	query, args := q.BuildToPatch("users", colMap, keys, q.BuildParam)
	res, err := s.DB.ExecContext(ctx, query, args...)
	if err != nil {
		return -1, err
	}
	return res.RowsAffected()
}

func (s *userService) Delete(ctx context.Context, id string) (int64, error) {
	query := "delete from users where id = ?"
	stmt, er0 := s.DB.Prepare(query)
	if er0 != nil {
		return -1, nil
	}
	res, er1 := stmt.ExecContext(ctx, id)
	if er1 != nil {
		return -1, er1
	}
	return res.RowsAffected()
}

func (s *userService) Search(ctx context.Context, filter UserFilter) (*Result, error) {
	query, params := BuildQuery(filter, s.BuildParam)
	rows, err := s.DB.QueryContext(ctx, query, params...)
	if err != nil {
		return nil, err
	}
	var users []User
	for rows.Next() {
		user := User{}
		err := rows.Scan(&user.Id, &user.Username, &user.Email, &user.Phone, &user.DateOfBirth)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	query, params = BuildCount(filter, s.BuildParam)
	rows, err = s.DB.QueryContext(ctx, query, params...)
	if err != nil {
		return nil, err
	}
	var total int64
	for rows.Next() {
		err := rows.Scan(&total)
		if err != nil {
			return nil, err
		}
	}
	return &Result{List: users, Total: total}, nil
}

func BuildCount(filter UserFilter, buildParam func(int) string) (string, []interface{}) {
	query := "select count(*) from users"
	where, params := BuildFilter(filter, buildParam)
	if len(where) > 0 {
		query = query + " where " + where
	}
	return query, params
}
func BuildQuery(filter UserFilter, buildParam func(int) string) (string, []interface{}) {
	query := "select * from users"
	where, params := BuildFilter(filter, buildParam)
	if len(where) > 0 {
		query = query + " where " + where
	}
	if filter.PageSize > 0 {
		query = query + fmt.Sprintf(` limit %d`, filter.PageSize)
		if filter.PageIndex > 0 {
			pageIndex := (filter.PageIndex - 1) * filter.PageSize
			query = query + fmt.Sprintf(` offset %d`, pageIndex)
		}
	}
	return query, params
}
func BuildFilter(filter UserFilter, buildParam func(int) string) (string, []interface{}) {
	var condition []string
	var params []interface{}
	i := 1

	if len(filter.Id) > 0 {
		params = append(params, filter.Id)
		condition = append(condition, fmt.Sprintf(`id = %s`, buildParam(i)))
		i++
	}
	if len(filter.Email) > 0 {
		q := "%" + filter.Email + "%"
		params = append(params, q)
		condition = append(condition, fmt.Sprintf(`email like %s`, buildParam(i)))
		i++
	}
	if len(filter.Username) > 0 {
		q := "%" + filter.Username + "%"
		params = append(params, q)
		condition = append(condition, fmt.Sprintf(`username like %s`, buildParam(i)))
		i++
	}
	if len(filter.Phone) > 0 {
		q := "%" + filter.Phone + "%"
		params = append(params, q)
		condition = append(condition, fmt.Sprintf(`phone like %s`, buildParam(i)))
		i++
	}

	if len(condition) > 0 {
		return strings.Join(condition, " and "), params
	} else {
		return "", params
	}
}
