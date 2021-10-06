package loan

type ProccesReservation struct {
	Id       int    `json:"id"`
	Status   int    `json:"status"`
	LoanDate string `json:"loan_date"`
}
