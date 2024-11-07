package request

type NewCarRequest struct {
	Brand  string `json:"brand" validate:"required"`
	Model  string `json:"model" validate:"required"`
	Color  string `json:"color" validate:"required"`
	Amount string `json:"amount" validate:"required"`
	IsUsed *bool  `json:"isUsed" validate:"required"`
	Km     *int   `json:"km" validate:"required"`
}
