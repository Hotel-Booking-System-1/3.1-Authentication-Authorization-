package constant

const (
	// Generic errors
	DefaultErrorMsg = "Something went wrong, please try again later or contact support."

	// User errors
	FailedToCreateUser = "Failed to create user."
	InvalidCredentials = "Invalid email or password. Please try again."
	NotFoundUser       = "User not found."

	// Guest errors
	FailedToCreateGuest = "Failed to create guest."
	InvalidGuestLogin   = "Email or password is wrong. Please try again later."
	NotFoundGuest       = "Guest not found."
)
