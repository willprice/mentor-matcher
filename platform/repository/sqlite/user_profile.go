package sqlite

import (
	"database/sql"
	"errors"

	_ "github.com/mattn/go-sqlite3"
	"github.com/willprice/mentor-matcher/domain"
)

type SQLiteUserProfileRepository struct {
	db *sql.DB
}

// NewSQLiteUserProfileRepository initializes the SQLite repository
func NewSQLiteUserProfileRepository(databasePath string) (*SQLiteUserProfileRepository, error) {
	db, err := sql.Open("sqlite3", databasePath)
	if err != nil {
		return nil, err
	}

	// Check if the table exists
	row := db.QueryRow("SELECT name FROM sqlite_master WHERE type='table' AND name='user_profiles';")
	var tableName string
	err = row.Scan(&tableName)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("table 'user_profiles' does not exist")
		}
		return nil, err
	}

	return &SQLiteUserProfileRepository{db: db}, nil
}

func (r *SQLiteUserProfileRepository) CreateUserProfile(profile *domain.UserProfile) (int, error) {
	result, err := r.db.Exec(
		"INSERT INTO user_profiles (first_name, last_name, username, location, email) VALUES (?, ?, ?, ?, ?)",
		profile.FirstName, profile.LastName, profile.Username, profile.Location, profile.Email,
	)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (r *SQLiteUserProfileRepository) GetUserProfileByEmail(email string) (*domain.UserProfile, error) {
	row := r.db.QueryRow(
		"SELECT id, first_name, last_name, username, location, email FROM user_profiles WHERE email = ?",
		email,
	)

	var profile domain.UserProfile
	err := row.Scan(&profile.ID, &profile.FirstName, &profile.LastName, &profile.Username, &profile.Location, &profile.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &profile, nil
}
