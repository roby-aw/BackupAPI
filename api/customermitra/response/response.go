package response

import "api-redeem-point/business/customermitra"

type Login struct {
	Code     string `json:"code"`
	Messages string `json:"messages"`
	Results  customermitra.ResponseLogin
}
