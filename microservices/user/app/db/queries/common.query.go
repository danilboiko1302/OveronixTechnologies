package queries

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"user/app/api/dto"
	model "user/app/db/models"
	voc "user/app/vocabulary"

	"github.com/jackc/pgx"
)

var SQLSession *sqlSession

type sqlSession struct {
	connection *pgx.Conn
}

func (s *sqlSession) GetUsers() ([]model.User, error) {
	rows, err := s.connection.Query(`SELECT id, login, password, "firstName", "lastName", to_char(birthday, 'YYYY-MM-DD') FROM users`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var result []model.User
	for rows.Next() {
		var user model.User
		err := rows.Scan(
			&user.Id,
			&user.Login,
			&user.Password,
			&user.FirstName,
			&user.LastName,
			&user.Birthday,
		)
		if err != nil {
			return nil, err
		}
		result = append(result, user)
	}

	return result, nil
}

func (s *sqlSession) DeleteUser(id string) (*model.User, error) {
	row := s.connection.QueryRow(`DELETE FROM users WHERE id = $1 RETURNING id, login, password, "firstName", "lastName", to_char(birthday, 'YYYY-MM-DD');`,
		id,
	)

	var (
		user model.User
		err  error
	)

	err = row.Scan(
		&user.Id,
		&user.Login,
		&user.Password,
		&user.FirstName,
		&user.LastName,
		&user.Birthday,
	)

	if err != nil {
		// err == sql.ErrNoRows not work ???
		if err.Error() == strings.Replace(sql.ErrNoRows.Error(), "sql: ", "", 1) {
			return nil, errors.New(voc.USER_NOT_FOUND)
		}

		return nil, err
	}

	return &user, nil
}

func (s *sqlSession) GetUser(id string) (*model.User, error) {
	row := s.connection.QueryRow(`SELECT id, login, password, "firstName", "lastName", to_char(birthday, 'YYYY-MM-DD') FROM users WHERE id = $1;`,
		id,
	)

	var user model.User

	err := row.Scan(
		&user.Id,
		&user.Login,
		&user.Password,
		&user.FirstName,
		&user.LastName,
		&user.Birthday,
	)

	if err != nil {
		// err == sql.ErrNoRows not work ???
		if err.Error() == strings.Replace(sql.ErrNoRows.Error(), "sql: ", "", 1) {
			return nil, errors.New(voc.USER_NOT_FOUND)
		}

		return nil, err
	}

	return &user, nil
}

func (s *sqlSession) CreateUser(data *dto.CreateUserDto) (*model.User, error) {
	row := s.connection.QueryRow(`INSERT INTO users (login, password, "firstName", "lastName", birthday) VALUES ($1, $2, $3, $4, $5) RETURNING id;`,
		data.Login,
		data.Password,
		data.FirstName,
		data.LastName,
		data.Birthday,
	)

	var (
		id  int
		err error
	)

	err = row.Scan(&id)

	if err != nil {
		return nil, err
	}

	return &model.User{
		Id:        id,
		Login:     data.Login,
		Password:  data.Password,
		FirstName: data.FirstName,
		LastName:  data.LastName,
		Birthday:  data.Birthday,
	}, nil
}

func Init() error {
	config, err := pgx.ParseConnectionString(fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		os.Getenv("SQL_USER"),
		os.Getenv("SQL_PASSWORD"),
		os.Getenv("SQL_HOST"),
		os.Getenv("SQL_PORT"),
		os.Getenv("SQL_DB")))

	if err != nil {
		return err
	}

	connection, err := pgx.Connect(config)

	if err != nil {
		return err
	}

	if err = connection.Ping(context.Background()); err != nil {
		return err
	}

	SQLSession = &sqlSession{
		connection: connection,
	}

	rows, err := connection.Query(`create table if not exists users (
		id serial not null,
		login varchar not null,
		password varchar not null,
		firstName varchar not null,
		lastName varchar not null,
		birthday date not null,
		PRIMARY KEY (id)
	)`)

	if err != nil {
		return err
	}

	rows.Close()

	rows, err = connection.Query("create unique index if not exists users_id_uindex on users (id)")

	if err != nil {
		return err
	}

	rows.Close()

	rows, err = connection.Query("create unique index if not exists users_login_uindex on users (login)")

	if err != nil {
		return err
	}

	rows.Close()

	return nil
}

func (s *sqlSession) Close() {
	s.connection.Close()
}
