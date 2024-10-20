package sqlite

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/Andrew-Savin-msk/sso/internal/domain/models"
	"github.com/Andrew-Savin-msk/sso/internal/store"
	"github.com/mattn/go-sqlite3"
)

type Storage struct {
	db *sql.DB
}

func New(path string) (*Storage, error) {
	const op = "sqlite.New"

	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Storage{db: db}, nil
}

func (s *Storage) Save(ctx context.Context, email string, passwdHash []byte) (int64, error) {
	const op = "sqlite.Save"

	stmt, err := s.db.Prepare("INSERT INTO users(email, passwd_hash) VALUES($1, $2);")
	if err != nil {
		return -1, fmt.Errorf("%s: %w", op, err)
	}

	res, err := stmt.ExecContext(ctx, email, passwdHash)
	if err != nil {
		var sqliteErr sqlite3.Error

		if errors.As(err, &sqliteErr) && sqliteErr.ExtendedCode == sqlite3.ErrConstraintUnique {
			return -1, fmt.Errorf("%s: %w", op, store.ErrUserExists)
		}
		return -1, fmt.Errorf("%s: %w", op, err)
	}

	return res.LastInsertId()
}

func (s *Storage) Get(ctx context.Context, email string) (*models.User, error) {
	const op = "sqlite.Get"

	stmt, err := s.db.Prepare("SELECT id, email, passwd_hash FROM users WHERE email = $1;")
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	row := stmt.QueryRowContext(ctx, email)

	u := &models.User{}
	err = row.Scan(&u.ID, &u.Email, &u.PasswdHash)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("%s: %w", op, store.ErrUserNotFound)
		}
		return nil, err
	}

	return u, nil
}

func (s *Storage) IsAdmin(ctx context.Context, userId int64) (bool, error) {
	const op = "sqlite.IsAdmin"

	stmt, err := s.db.Prepare("SELECT is_admin FROM users WHERE id = $1;")
	if err != nil {
		return false, fmt.Errorf("%s: %w", op, err)
	}

	row := stmt.QueryRowContext(ctx, userId)

	var isAdmin bool
	err = row.Scan(&isAdmin)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, fmt.Errorf("%s: %w", op, store.ErrUserNotFound)
		}
		return false, err
	}

	return isAdmin, nil
}

func (s *Storage) App(ctx context.Context, appId int) (*models.App, error) {
	const op = "sqlite.App"

	stmt, err := s.db.Prepare("SELECT id, name, secret FROM apps WHERE id = $1;")
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	row := stmt.QueryRowContext(ctx, appId)

	a := &models.App{}
	err = row.Scan(&a.ID, &a.Name, &a.Secret)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("%s: %w", op, store.ErrUserNotFound)
		}
		return nil, err
	}

	return a, nil
}
