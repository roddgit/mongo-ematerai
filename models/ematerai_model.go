package models

// import "go.mongodb.org/mongo-driver/bson/primitive"

type Login struct {
	// Id       primitive.ObjectID `json:"id,omitempty"`
	User     string `json:"user,omitempty" validate:"required"`
	Password string `json:"password,omitempty" validate:"required"`
}

// https://mholt.github.io/json-to-go/
type AutoGenerated struct {
	Message string `json:"message"`
	Result  struct {
		Data struct {
			Login struct {
				Token string `json:"token"`
				User  struct {
					Accounts []struct {
						Companys []struct {
							Nama string `json:"nama"`
						} `json:"companys"`
						Desc      interface{} `json:"desc"`
						ID        string      `json:"id"`
						Parent    interface{} `json:"parent"`
						Statusreg string      `json:"statusreg"`
						Tipeacct  string      `json:"tipeacct"`
						Tipepay   string      `json:"tipepay"`
					} `json:"accounts"`
					Email       string `json:"email"`
					FirstLogin  bool   `json:"firstLogin"`
					FirstName   string `json:"firstName"`
					ID          string `json:"id"`
					LastName    string `json:"lastName"`
					PhoneNumber string `json:"phoneNumber"`
					Role        string `json:"role"`
					Userdetails []struct {
						Locationother []struct {
							CustID  int    `json:"custID"`
							ID      string `json:"id"`
							Loccode string `json:"loccode"`
							Loctype struct {
								Description string `json:"description"`
							} `json:"loctype"`
							Tpay string `json:"tpay"`
						} `json:"locationother"`
						Locations []struct {
							CustID  int    `json:"custID"`
							ID      string `json:"id"`
							Loccode string `json:"loccode"`
							Loctype struct {
								Description string `json:"description"`
							} `json:"loctype"`
							Tpay string `json:"tpay"`
						} `json:"locations"`
					} `json:"userdetails"`
				} `json:"user"`
			} `json:"login"`
		} `json:"data"`
	} `json:"result"`
	StatusCode string `json:"statusCode"`
	Token      string `json:"token"`
}
