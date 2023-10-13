package web

type UpdateRequest struct {
	Id    int    `validate:"required"`
	Name  string `validate:"required,min=1,max=200"`
	Kelas string `validate:"required,min=1,max=200"`
}
