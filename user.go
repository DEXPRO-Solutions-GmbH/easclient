package easclient

// User is a custom type to represent one of the four EAS user instances.
type User string

const (
	// UserAdmin is the administrator which most permissions.
	UserAdmin User = "eas_administrator"
	// UserKeeper is responsible for auditing, moderating, etc.
	UserKeeper User = "eas_keeper"
	// UserUser is the default user being able to read and write.
	UserUser User = "eas_user"
	// UserGuest is a guest user.
	UserGuest User = "eas_guest"
)

type RequestHeader string

const (
	HeaderUser         RequestHeader = "x-otris-eas-user"
	HeaderUserFullname RequestHeader = "x-otris-eas-user-fullname"
	HeaderTokens       RequestHeader = "x-otris-eas-tokens"
)
