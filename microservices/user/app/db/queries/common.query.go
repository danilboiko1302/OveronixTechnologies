package queries

import (
	"database/sql"
	"fmt"
	"os"
	"user/app/api/dto"
	model "user/app/db/models"

	_ "github.com/lib/pq"
)

var SQLSession *sqlSession

type sqlSession struct {
	connection *sql.DB
}

func (s *sqlSession) CreateUser(data *dto.CreateUserDto) (*model.User, error) {
	stmt, err := s.connection.Prepare("INSERT INTO users (login, password, firstName, lastName, birthday) VALUES (? ? ? ? ?);")
	if err != nil {
		return nil, err
	}

	defer stmt.Close()
	result, err := stmt.Exec(
		data.Login,
		data.Password,
		data.FirstName,
		data.LastName,
		data.Birthday,
	)

	if err != nil {
		fmt.Println("4")
		return nil, err
	}

	fmt.Println(result)

	return &model.User{}, nil
}

func Init() error {

	connection, err := sql.Open("postgres",
		fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			os.Getenv("SQL_HOST"),
			os.Getenv("SQL_PORT"),
			os.Getenv("SQL_USER"),
			os.Getenv("SQL_PASSWORD"),
			os.Getenv("SQL_DB")))

	if err != nil {
		return err
	}

	if err = connection.Ping(); err != nil {
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
