package utilities

const (
	Administrator = 1
	Officer       = 2
)

type Role struct {
	RoleName   map[string]int
	RoleNumber map[int]string
}

var roles = Role{
	RoleName: map[string]int{
		"Administrator": 1,
		"Officer":       2,
	},
	RoleNumber: map[int]string{
		1: "Administrator",
		2: "Officer",
	},
}
