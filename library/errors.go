package library

import (
	"fmt"
)

// PermissionDeniedError occurs when a librarian attempts an operation
//   it does not have permission to carry out.
type PermissionDeniedError struct {
	// Role is the role of the librarian attempting an operation
	Role string 
}

// Error satisfies the error interface.
func (pde *PermissionDeniedError) Error() string {
	return fmt.Sprintf("The role: %s does not have permission for this operation.", pde.Role)
}