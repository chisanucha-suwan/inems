package utilities

func GetRole(role interface{}, toCode bool) interface{} {
	if toCode {
		if code, ok := role.(string); ok {
			if val, exists := roles.RoleName[code]; exists {
				return val
			}
		}
		return -1
	} else {
		if name, ok := role.(int); ok {
			if val, exists := roles.RoleNumber[name]; exists {
				return val
			}
		}
		return ""
	}
}
