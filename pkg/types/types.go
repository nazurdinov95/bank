package types

// Money represents a monetary amount in minimum units (cents, kopecks, diramas, etc.).
type Money int64

// Category передставляет собой категорию, в которой был совершён платёж (авто, аптеки, рестораны и т.д.).
type Category string

//Status представляет собой статус платежа.
type Status string

//Передопределённые статусы платежей.
const (
	StatusOk Status = "OK"
	StatusFail Status = "FAIL"
	StatusInProgress Status = "INPROGRESS"
)

// Payment presents information about the payment.
type Payment struct {
	ID 			int
	Amount 		Money
	Category 	Category
	Status		Status
}

type PaymentSource struct {
	Type 		string // 'card'
	Number 		string // номер вида '5058 xxxx xxxx 8888'
	Balance 	Money // баланс в дирамах
}
