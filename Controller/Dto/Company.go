package Dto

type CreateCompany struct {
	Name    string  `json:"name"`
	Balance float64 `json:"balance"`
}

type UpdatedCompanyBalance struct {
	Id         int     `json:"id"`
	Name       string  `json:"name"`
	OldBalance float64 `json:"old_balance"`
	NewBalance float64 `json:"new_balance"`
}

type GetCompanyProfile struct {
	Id      int     `json:"id"`
	Name    string  `json:"name"`
	Balance float64 `json:"balance"`
}

type ViewAllCompany struct {
	Id      int     `json:"id"`
	Name    string  `json:"name"`
	Balance float64 `json:"balance"`
}
