package queries

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx"
)

var SQLSession *sqlSession

type sqlSession struct {
	connection *pgx.Conn
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
