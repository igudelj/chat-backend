package entities

type UserSearchField string

const (
	UserKeycloakFieldID     UserSearchField = "keycloak_id"
	UserSearchFieldID       UserSearchField = "id"
	UserSearchFieldEmail    UserSearchField = "email"
	UserSearchFieldUsername UserSearchField = "username"
)

func ParseUserSearchField(s string) (UserSearchField, bool) {
	switch s {
	case string(UserSearchFieldID),
		string(UserKeycloakFieldID),
		string(UserSearchFieldEmail),
		string(UserSearchFieldUsername):
		return UserSearchField(s), true
	default:
		return "", false
	}
}
