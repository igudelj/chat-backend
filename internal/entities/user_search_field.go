package entities

type UserSearchField string

const (
	UserSearchFieldID       UserSearchField = "id"
	UserSearchFieldEmail    UserSearchField = "email"
	UserSearchFieldUsername UserSearchField = "username"
)

func ParseUserSearchField(s string) (UserSearchField, bool) {
	switch s {
	case string(UserSearchFieldID),
		string(UserSearchFieldEmail),
		string(UserSearchFieldUsername):
		return UserSearchField(s), true
	default:
		return "", false
	}
}
