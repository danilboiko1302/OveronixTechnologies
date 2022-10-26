package queries

import (
	"bytes"
	"database/sql"
	"errors"
	"log"
	"strconv"
	"strings"
	"user/app/api/dto"
	model "user/app/db/models"
	voc "user/app/vocabulary"

	"github.com/jackc/pgx"
)

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

func (s *sqlSession) UpdateUser(data map[string]string, id string) (*model.User, error) {
	row, err := buildPreparedStatementForUpdate(s, data, id)

	if err != nil {
		return nil, err
	}

	var user model.User

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

func buildPreparedStatementForUpdate(s *sqlSession, data map[string]string, id string) (*pgx.Row, error) {
	var sqlRequest bytes.Buffer
	var values []string = make([]string, 0, len(data))

	sqlRequest.WriteString("UPDATE users SET")

	for key, val := range data {
		sqlRequest.WriteString(" ")
		sqlRequest.WriteString(key)
		sqlRequest.WriteString(" = $")

		values = append(values, val)

		sqlRequest.WriteString(strconv.Itoa(len(values)))

		if len(values) != len(data) {
			sqlRequest.WriteString(",")
		}
	}

	sqlRequest.UnreadRune()

	sqlRequest.WriteString(" WHERE id = $")
	sqlRequest.WriteString(strconv.Itoa(len(values) + 1))
	sqlRequest.WriteString(" RETURNING id, login, password, \"firstName\", \"lastName\", to_char(birthday, 'YYYY-MM-DD');")

	return insertData(s, sqlRequest.String(), values, id)
}

func insertData(s *sqlSession, sqlRequest string, values []string, id string) (*pgx.Row, error) {
	if len(values) == 1 {
		return s.connection.QueryRow(sqlRequest,
			values[0],
			id,
		), nil
	}

	if len(values) == 2 {
		return s.connection.QueryRow(sqlRequest,
			values[0],
			values[1],
			id,
		), nil
	}

	if len(values) == 3 {
		return s.connection.QueryRow(sqlRequest,
			values[0],
			values[1],
			values[2],
			id,
		), nil
	}

	if len(values) == 4 {
		return s.connection.QueryRow(sqlRequest,
			values[0],
			values[1],
			values[2],
			values[3],
			id,
		), nil
	}

	return nil, errors.New(voc.TOO_MUCH_DATA_FOR_UPDATE)
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
