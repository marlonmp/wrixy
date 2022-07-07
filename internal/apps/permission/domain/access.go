package domain

type Access uint8

const (
	Private Access = iota
	Public
	Banned
	Deleted
)

var access = [...]string{
	"Privado",
	"Público",
	"Bloqueado",
	"Eliminado",
}

func NewAccess(i int) Access {
	if i < 0 || i >= len(access) {
		return Access(0)
	}
	return Access(i)
}

func (a Access) String() string {
	return access[a]
}
