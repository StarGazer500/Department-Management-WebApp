package repository

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq" // or your preferred SQL driver
	"github.com/StarGazer500/Department-Management-WebApp/internals/accounts/models"
)



// CheckIfRoleExistAndGetID checks if a role exists by name and returns its ID.
func CheckIfRoleExistAndGetID(db *sql.DB, roleName string) (int, error) {
	query := `
		SELECT id
		FROM roles
		WHERE name = $1`
	
	var id int
	err := db.QueryRow(query, roleName).Scan(&id)
	if err == sql.ErrNoRows {
		return 0, errors.New("role not found")
	}
	if err != nil {
		return 0, fmt.Errorf("failed to query role: %w", err)
	}
	return id, nil
}

func CreateRole(db *sql.DB, roleName, roleDescription string) (*models.Role, error) {
	// Check if role already exists
	var exists bool
	checkQuery := `SELECT EXISTS(SELECT 1 FROM roles WHERE name = $1)`
	err := db.QueryRow(checkQuery, roleName).Scan(&exists)
	if err != nil {
		return nil, fmt.Errorf("error checking if role exists: %v", err)
	}
	if exists {
		return nil, fmt.Errorf("role '%s' already exists", roleName)
	}

	// Insert new role and return the created role
	var role models.Role
	insertQuery := `INSERT INTO roles (name, description) VALUES ($1, $2) RETURNING id, name, description`
	err = db.QueryRow(insertQuery, roleName, roleDescription).Scan(&role.ID, &role.Name, &role.Description)
	if err != nil {
		return nil, fmt.Errorf("error creating role: %v", err)
	}
	return &role, nil
}

func CreateUserRoles(db *sql.DB, userID, roleID int) (*models.UserRole, error) {
	// Check if user-role combination already exists
	var exists bool
	checkQuery := `SELECT EXISTS(SELECT 1 FROM user_roles WHERE user_id = $1 AND role_id = $2)`
	err := db.QueryRow(checkQuery, userID, roleID).Scan(&exists)
	if err != nil {
		return nil, fmt.Errorf("error checking if user-role exists: %v", err)
	}
	if exists {
		return nil, fmt.Errorf("user-role combination already exists")
	}

	// Insert new user-role and return the created record
	var userRole models.UserRole
	insertQuery := `INSERT INTO user_roles (user_id, role_id) VALUES ($1, $2) RETURNING id, user_id, role_id`
	err = db.QueryRow(insertQuery, userID, roleID).Scan(&userRole.ID, &userRole.UserID, &userRole.RoleID)
	if err != nil {
		return nil, fmt.Errorf("error creating user-role: %v", err)
	}
	return &userRole, nil
}

func CreateUser(db *sql.DB, email, passwordHash, firstName, lastName, roleName string) (*models.Users, error) {
	// Check if user already exists
	var exists bool
	checkQuery := `SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)`
	err := db.QueryRow(checkQuery, email).Scan(&exists)
	if err != nil {
		return nil, fmt.Errorf("error checking if email exists: %v", err)
	}
	if exists {
		return nil, fmt.Errorf("user with email '%s' already exists", email)
	}

	// Check if role exists and get its ID
	roleID, err := CheckIfRoleExistAndGetID(db, roleName)
	if err != nil {
		return nil, fmt.Errorf("role validation failed: %v", err)
	}

	// Insert new user and return the created user
	var user models.Users
	insertQuery := `INSERT INTO users (email, password_hash, first_name, last_name) VALUES ($1, $2, $3, $4) RETURNING id, email, password_hash, first_name, last_name`
	err = db.QueryRow(insertQuery, email, passwordHash, firstName, lastName).Scan(&user.ID, &user.Email, &user.PasswordHash, &user.FirstName, &user.LastName)
	if err != nil {
		return nil, fmt.Errorf("error creating user: %v", err)
	}

	// Create user-role association
	userRole, err := CreateUserRoles(db, user.ID, roleID)
	if err != nil {
		return nil, fmt.Errorf("user-role could not be created: %v", err)
	}

	fmt.Printf("Created user-role association: %+v\n", userRole)
	return &user, nil
}

func GetUserWithRoles(db *sql.DB, email string) (*models.UserWithRoles, error) {
	query := `
		SELECT u.id, u.email, u.password_hash, u.first_name, u.last_name, r.id, r.name, r.description
		FROM users u
		LEFT JOIN user_roles ur ON u.id = ur.user_id
		LEFT JOIN roles r ON ur.role_id = r.id
		WHERE u.email = $1`

	rows, err := db.Query(query, email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result *models.UserWithRoles
	roleMap := make(map[int]bool) // To avoid duplicate roles

	for rows.Next() {
		var user models.Users
		var roleID sql.NullInt64
		var roleName sql.NullString
		var roleDescription sql.NullString

		err := rows.Scan(&user.ID, &user.Email, &user.PasswordHash, &user.FirstName, &user.LastName, &roleID, &roleName, &roleDescription)
		if err != nil {
			return nil, err
		}

		fmt.Println(user)

		// Initialize user if this is the first row
		if result == nil {
			result = &models.UserWithRoles{
				User:  user,
				Roles: []models.Role{},
			}
		}

		// Add role if it exists and hasn't been added yet
		if roleID.Valid && roleName.Valid && !roleMap[int(roleID.Int64)] {
			role := models.Role{
				ID:          int(roleID.Int64),
				Name:        roleName.String,
				Description: roleDescription.String,
			}
			result.Roles = append(result.Roles, role)
			roleMap[int(roleID.Int64)] = true
		}
	}

	if result == nil {
		return nil, sql.ErrNoRows
	}

	return result, nil
}