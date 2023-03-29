package model

type User struct {
	ID                  string  ` json:"id" `
	Email               string  ` json:"email" validate:"required" `
	Username            string  ` json:"username" validate:"required" `
	Password            string  ` json:"-" `
	RefreshToken        *string ` json:"-" `
	FirstName           string  ` json:"firstName" validate:"required" `
	LastName            string  ` json:"lastName" validate:"required" `
	CreatedAt           int64   ` json:"createdAt" `
	CreatedBy           string  ` json:"createdBy" validate:"required" `
	ModifiedAt          int64   ` json:"modifiedAt" `
	ModifiedBy          string  ` json:"modifiedBy" validate:"required" `
	PasswordFail        string  ` json:"passwordFail" `
	LastSuccessLogin    int64   ` json:"lastSuccessLogin" `
	LastFailLogin       int64   ` json:"latFailLogin" `
	LastChangedPassword int64   ` json:"latChangePassword" `
	SecurityCode        *string ` json:"-" `
}
