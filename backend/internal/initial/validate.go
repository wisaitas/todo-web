package initial

import (
	"github.com/wisaitas/todo-web/internal/validates"
)

func initializeValidates() *Validates {
	return &Validates{
		UserValidate:        *validates.NewUserValidate(),
		AuthValidate:        *validates.NewAuthValidate(),
		ProvinceValidate:    *validates.NewProvinceValidate(),
		DistrictValidate:    *validates.NewDistrictValidate(),
		SubDistrictValidate: *validates.NewSubDistrictValidate(),
		TodoValidate:        *validates.NewTodoValidate(),
	}
}

type Validates struct {
	UserValidate        validates.UserValidate
	AuthValidate        validates.AuthValidate
	ProvinceValidate    validates.ProvinceValidate
	DistrictValidate    validates.DistrictValidate
	SubDistrictValidate validates.SubDistrictValidate
	TodoValidate        validates.TodoValidate
}
