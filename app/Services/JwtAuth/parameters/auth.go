package parameters

import ()

type TokenAuthentication struct {
	Token string `json:"access_token" form:"token"`
	// Refresh_token string `json:"access_token" form:"token"`
	
	Token_type string `json:"token_type" default:"Bearer"`
	Expires_in int `json:"expires_in"  default:"1295999"`
}
