package fiber

//
// locals constants is a group of variables that represents
// the key on a context fiber value, e.p.
//
//	c.Locals(LocalKey, "some string")
//
//	value := c.Locals(LocalKey) // "some string"
//
const (
	LocalsUserFilter = "userFilter"
	LocalsUserDomain = "userDomain"
)
