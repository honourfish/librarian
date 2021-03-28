package errors

import (
	"fmt"
)

// PermissionDeniedError occurs when a librarian attempts an operation
//   it does not have permission to carry out.
type PermissionDeniedError struct {
	// Role is the role of the librarian attempting an operation
	Role string 
}

// NotEnoughCopiesError occurs when there are not enough copies.
type NotEnoughCopiesError struct {}

// NoteCheckedOutError occurs when a book is not checked out, but is thought to have been.
type NotCheckedOutError struct{}

// UserAlreadyExistsError occurs when a user exists when trying to be added.
type UserAlreadyExistsError struct{}

// UserDoesNotExistError occurs when an operation is attempted on a user that does not exist.
type UserDoesNotExistError struct{}

// Error satisfies the error interface.
func (pde *PermissionDeniedError) Error() string {
	return fmt.Sprintf("The role: %s does not have permission for this operation.", pde.Role)
}

// Error satisfies the error interface.
func (nece *NotEnoughCopiesError) Error() string {
	return fmt.Sprintf("The Book does not have enough copies.")
}

// Error satisfies the error interface.
func (ncoe *NotCheckedOutError) Error() string {
	return fmt.Sprintf("The Book is not checked out.")
}

// Error satisfies the error interface.
func (uaee *UserAlreadyExistsError) Error() string {
	return fmt.Sprintf("The User already exists.")
}

// Error satisfies the error interface.
func (udnee *UserDoesNotExistError) Error() string {
	return fmt.Sprintf("The User does not exist.")
}