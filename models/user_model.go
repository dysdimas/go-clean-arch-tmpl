package models

import (
	"database/sql"
	"errors"
	"golang.org/x/crypto/bcrypt"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type UserModel struct {
	db *sql.DB
}

func NewUserModel(db *sql.DB) *UserModel {
	return &UserModel{
		db: db,
	}
}

func (m *UserModel) CreateUser(user *User) error {
	// Encrypt the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Insert the user into the database
	query := "INSERT INTO users (username, password, role) VALUES (?, ?, ?)"
	result, err := m.db.Exec(query, user.Username, string(hashedPassword), user.Role)
	if err != nil {
		return err
	}

	// Check the number of affected rows
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected != 1 {
		return errors.New("user not created")
	}

	return nil
}

func (m *UserModel) FetchAllUser() ([]map[string]interface{}, error) {
	// Execute the query to fetch the user list
	rows, err := m.db.Query("SELECT username, role FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Get the column names from the result set
	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	// Create a slice to store the results
	results := make([]map[string]interface{}, 0)

	// Create a slice to hold the column values for the current row
	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	// Iterate over the rows
	for rows.Next() {
		// Scan the row into the scanArgs slice
		err := rows.Scan(scanArgs...)
		if err != nil {
			return nil, err
		}

		// Create a map to hold the column names and values for the current row
		rowData := make(map[string]interface{})

		// Map the column names to the column values in the rowData map
		for i, column := range columns {
			rowData[column] = string(values[i])
		}

		// Append the rowData map to the results slice
		results = append(results, rowData)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

func (m *UserModel) GetUserByUsername(username string) (*User, error) {
	// Query the database to fetch the user by username
	query := "SELECT id, username, password, role FROM users WHERE username = ?"
	row := m.db.QueryRow(query, username)

	user := &User{}
	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Role)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return user, nil
}

func (m *UserModel) DeleteUserByUsername(username string) (string, error) {
	query := "DELETE FROM users WHERE username = ?"
	_, err := m.db.Exec(query, username)
	if err != nil {
		return "", err
	}
	return "", err
}

func (m *UserModel) GetUserByUsernameAndPassword(username, password string) (*User, error) {
	// Query the database to fetch the user by username
	query := "SELECT id, username, password, role FROM users WHERE username = ?"
	row := m.db.QueryRow(query, username)

	user := &User{}
	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Role)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	// Compare the provided password with the stored hashed password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}
