package company

const (
	CreateCompanyQuery   = `INSERT INTO public.company (company_name, company_balance) VALUES($1, $2) RETURNING company_id`
	GetCompanyProfile    = `SELECT company_id, company_name, company_balance FROM public.company WHERE company_id = $1`
	ViewAllCompanies     = `SELECT company_id, company_name, company_balance FROM public.company`
	UpdateCompanyBalance = `UPDATE public.company SET company_balance = $1 WHERE company_id = $2 RETURNING company_id, company_name, company_balance`
)
