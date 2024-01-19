package consts

type ContextKey interface{}

var (
	ContextKeyCreateUserRequest = ContextKey("create_user_request")
	ContextKeyUpdateUserRequest = ContextKey("update_user_request")
)
