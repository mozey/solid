package main

import (
	"fmt"
	"os"
)

// Example demonstrating the Single Responsibility Principle (SRP):

// User struct handles user data.
// It only knows about user info
type User struct {
	ID       int
	Username string
	Email    string
}

// UserValidator handles user data validation.
// It only knows how to validate User data
type UserValidator struct {
}

// Validate the User
func (v *UserValidator) Validate(user *User) error {
	if user.Username == "" {
		return fmt.Errorf("username cannot be empty")
	}
	if user.Email == "" {
		return fmt.Errorf("email cannot be empty")
	}
	// ... more validation rules ...
	return nil
}

// UserSerializer handles serializing User data to a string.
// It only knows how to format User data.
type UserSerializer struct {
}

func (s *UserSerializer) Serialize(user *User) string {
	return fmt.Sprintf(
		"ID: %d, Username: %s, Email: %s", user.ID, user.Username, user.Email)
}

// UserRepository handles storing and retrieving User data.
// It only knows how to persist User data
type UserRepository struct {
	users  map[int]*User
	nextID int
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		users:  make(map[int]*User),
		nextID: 1,
	}
}

// Save method is a simplified example,
// would interact with a database or other storage in a real application
func (r *UserRepository) Save(user *User) error {
	if _, exists := r.users[user.ID]; exists {
		return fmt.Errorf("user with ID %d already exists", user.ID)
	}
	user.ID = r.nextID
	r.nextID++
	r.users[user.ID] = user
	return nil
}

func (r *UserRepository) Get(id int) (*User, error) {
	user, exists := r.users[id]
	if !exists {
		return nil, fmt.Errorf("user with ID %d not found", id)
	}
	return user, nil
}

func main() {
	repo := NewUserRepository()
	validator := &UserValidator{}
	serializer := &UserSerializer{}

	newUser := &User{Username: "john_doe", Email: "john.doe@example.com"}

	// Validate the user
	if err := validator.Validate(newUser); err != nil {
		fmt.Println("Validation error:", err)
		os.Exit(1)
	}

	// Save the user
	if err := repo.Save(newUser); err != nil {
		fmt.Println("Error saving user:", err)
		os.Exit(1)
	}

	retrievedUser, err := repo.Get(newUser.ID)
	if err != nil {
		fmt.Println("Error retrieving user:", err)
		os.Exit(1)
	}

	// Serialize the user data
	serializedUser := serializer.Serialize(retrievedUser)
	fmt.Println("Serialized user:", serializedUser)

	// Try to save a user with an existing ID
	// (demonstrates the repository's responsibility)
	existingUser := &User{
		ID:       newUser.ID, // same ID that was previously saved
		Username: "another_user",
		Email:    "another@example.com"}
	if err := repo.Save(existingUser); err != nil {
		fmt.Println("Error saving existing user:", err)
		// Expecting an error here
	}

	os.Exit(0)

}
