package domain

type OperationType string

const (
	OperationTypeDeposit  OperationType = "deposit"
	OperationTypeWithdraw OperationType = "withdraw"
	OperationTypePayment  OperationType = "payment"
	OperationTypeHold     OperationType = "hold"
	OperationTypeRefund   OperationType = "refund"
)
