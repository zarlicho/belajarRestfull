package web

type CreateRequest struct {
	Name  string `validate:"required,min=1,max=200"`
	Kelas string `validate:"required,min=1,max=200"`
}
type RegisRequest struct {
	Name     string `validate:"required,min=1,max=200"`
	Password string `validate:"required,min=1,max=200"`
}
type LoginRequest struct {
	Name     string `validate:"required,min=1,max=200"`
	Password string `validate:"required,min=1,max=200"`
}
