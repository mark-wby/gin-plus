package model

type UserModel struct {
	Id int
	Name string
}

func (this *UserModel) String() string {
	return "userModel"
}
