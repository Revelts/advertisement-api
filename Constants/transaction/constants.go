package transaction

type TransactionType int

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
