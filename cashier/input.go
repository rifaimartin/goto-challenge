package cashier

type CasherInput struct {
	Name        string `json:"name" binding:"required"`
}

//LoginInput is struct
type LoginInput struct {
	PassCode    string `json:"passcode" form:"passcode" binding:"required"`
}
