package globaltypes

type AuthResponse struct {
	ID int `json:"id"`
	Role string `json:"role"`
	Token string `json:"token"`
}
