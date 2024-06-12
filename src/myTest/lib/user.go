package lib

type User struct {
	FirstName, SirName string
}

func (u User) GetFullName() string {
	return u.FirstName + " " + u.SirName
}
