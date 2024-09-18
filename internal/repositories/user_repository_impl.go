package repositories

import (
	"database/sql"
	"errors"

	"github.com/Zainal21/renco-boilerplate/internal/dtos"
	"github.com/Zainal21/renco-boilerplate/internal/entity"
	"github.com/jmoiron/sqlx"
)

type baseUserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return &baseUserRepository{db: db}
}

func (b *baseUserRepository) CreateUser(CreateUserDto *dtos.CreateUserDto) (int, error) {
	tx, err := b.db.Beginx()
	if err != nil {
		return 0, err
	}
	defer func() {
		tx.Rollback()
	}()

	query, args, err := tx.BindNamed(`
	INSERT INTO users (uid, firebase_uid, email, name, phone, created_at, updated_at)
	VALUES (:uid, :firebase_uid, :email, :name, :phone, :created_at, :updated_at)
	RETURNING id;
	`, CreateUserDto)
	if err != nil {
		return 0, err
	}
	var userID int
	err = tx.Get(&userID, query, args...)
	if err != nil {
		return 0, err
	}

	return userID, nil
}

func (b *baseUserRepository) GetUserByEmail(email string) (*entity.User, error) {
	var user entity.User

	err := b.db.Get(&user, "SELECT * FROM users WHERE email = $1;", email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return &user, nil
}

func (b *baseUserRepository) GetUserByFirebaseUID(UID string) (*entity.User, error) {
	var user entity.User

	err := b.db.Get(&user, "SELECT * FROM users WHERE firebase_uid = $1;", UID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return &user, nil
}

func (b *baseUserRepository) GetUserByUID(UID string) (*entity.User, error) {
	var user entity.User

	err := b.db.Get(&user, "SELECT * FROM users WHERE uid = $1;", UID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return &user, nil
}
