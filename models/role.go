package models

func GetRoleName(roleID int) string {
	switch roleID {
	case 1:
		return "user"
	default:
		return "admin"
	}
}
