// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"
)

// DeleteLibrarianUsernameUserUserURL generates an URL for the delete librarian username user user operation
type DeleteLibrarianUsernameUserUserURL struct {
	User     string
	Username string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *DeleteLibrarianUsernameUserUserURL) WithBasePath(bp string) *DeleteLibrarianUsernameUserUserURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *DeleteLibrarianUsernameUserUserURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *DeleteLibrarianUsernameUserUserURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/librarian/{username}/user/{user}"

	user := o.User
	if user != "" {
		_path = strings.Replace(_path, "{user}", user, -1)
	} else {
		return nil, errors.New("user is required on DeleteLibrarianUsernameUserUserURL")
	}

	username := o.Username
	if username != "" {
		_path = strings.Replace(_path, "{username}", username, -1)
	} else {
		return nil, errors.New("username is required on DeleteLibrarianUsernameUserUserURL")
	}

	_basePath := o._basePath
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *DeleteLibrarianUsernameUserUserURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *DeleteLibrarianUsernameUserUserURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *DeleteLibrarianUsernameUserUserURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on DeleteLibrarianUsernameUserUserURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on DeleteLibrarianUsernameUserUserURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *DeleteLibrarianUsernameUserUserURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
