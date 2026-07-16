package domain

type TaskStatus string

const (
	TaskStatusPublished  TaskStatus = "published"
	TaskStatusInProgress TaskStatus = "in_progress"
	TaskStatusCanceled   TaskStatus = "canceled"
	TaskStatusOnReview   TaskStatus = "on_review"
	TaskStatusCompleted  TaskStatus = "completed"
	TaskStatusRevision   TaskStatus = "revision"

	BidsStatusPending  TaskStatus = "pending"
	BidsStatusRejected TaskStatus = "rejected"
	BidsStatusAccepted TaskStatus = "accepted"

	TransactionStatusHold     TaskStatus = "hold"
	TransactionStatusPaid     TaskStatus = "paid"
	TransactionStatusReturned TaskStatus = "returned"
)
