package models

type TaskStatus string
type BidStatus string
type TransactionStatus string

const (
	TaskStatusPublished  TaskStatus = "published"
	TaskStatusInProgress TaskStatus = "in_progress"
	TaskStatusCanceled   TaskStatus = "canceled"
	TaskStatusOnReview   TaskStatus = "on_review"
	TaskStatusCompleted  TaskStatus = "completed"
	TaskStatusRevision   TaskStatus = "revision"
)

const (
	TransactionStatusHold     TransactionStatus = "hold"
	TransactionStatusPaid     TransactionStatus = "paid"
	TransactionStatusReturned TransactionStatus = "returned"
)

const (
	BidsStatusPending  BidStatus = "pending"
	BidsStatusRejected BidStatus = "rejected"
	BidsStatusAccepted BidStatus = "accepted"
)
