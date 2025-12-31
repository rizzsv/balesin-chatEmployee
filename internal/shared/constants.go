package shared

const (
	// User roles
	RoleAdmin    = "admin"
	RoleEmployee = "employee"
	RoleUser     = "user"

	// JWT context keys
	ContextKeyUserID = "user_id"
	ContextKeyRole   = "role"

	// WebSocket message types
	MessageTypeText  = "text"
	MessageTypeImage = "image"
	MessageTypeFile  = "file"
)
