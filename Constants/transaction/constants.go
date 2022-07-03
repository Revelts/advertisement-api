package transaction

type TransactionType int

const (
	CreateTransaction = `INSERT INTO public.transaction 
		(advertisement_id, company_id, start_advertising, end_advertising, 
		advertisement_location, fee) 
	VALUES($1, $2, $3, $4, $5, $6) RETURNING advertisement_id, company_id, TO_CHAR(start_advertising, 'YYYY-MM-DD HH24:MI:SS'), TO_CHAR(end_advertising, 'YYYY-MM-DD HH24:MI:SS'), 
		advertisement_location, fee`
)

const (
	Pending = iota
	Started
	Expired
)

func (constantStatus TransactionType) Transaction() string {
	statusString := map[TransactionType]string{
		Pending: "Pending",
		Started: "Started",
		Expired: "Expired",
	}
	result, exist := statusString[constantStatus]
	if !exist {
		return "Data not found"
	}
	return result
}
