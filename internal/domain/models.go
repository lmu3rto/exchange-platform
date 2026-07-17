package domain

import "time"

type User struct {
	ID        int64     `json:"id"`
	Name      string    `json:"user_name"`
	Balance   float64   `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

type DeletedUser struct {
	ID        int64     `json:"id"`
	Name      string    `json:"user_name"`
	Balance   float64   `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

type Customer struct {
	ID     int64 `json:"id"`
	UserID int64 `json:"user_id"`
}

type Executor struct {
	ID     int64 `json:"id"`
	UserID int64 `json:"user_id"`
}

type Task struct {
	ID            int64      `json:"id"`
	CustomerID    int64      `json:"customer_id"`
	ExecutorID    *int64     `json:"executor_id,omitempty"`
	Title         string     `json:"title"`
	Description   string     `json:"description"`
	Status        TaskStatus `json:"status"`
	AcceptedBidID int64      `json:"accepted_bid_id"`
	Deadline      time.Time  `json:"deadline"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
}

type Bids struct {
	ID         int64      `json:"id"`
	TaskID     int64      `json:"task_id"`
	ExecutorID int64      `json:"executor_id"`
	Message    string     `json:"message"`
	Status     TaskStatus `json:"status"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
}

type Transaction struct {
	ID             int64      `json:"id"`
	TaskID         int64      `json:"task_id"`
	FromCustomerID int64      `json:"from_customer"`
	ToExecutorID   int64      `json:"to_executor"`
	Amount         float64    `json:"amount"`
	Status         TaskStatus `json:"status"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
}

type WalletTransactions struct {
	ID          int64         `json:"id"`
	UserID      int64         `json:"user_id"`
	Amount      float64       `json:"amount"`
	Operation   OperationType `json:"operation_type"`
	TaskID      int64         `json:"task_id"`
	Description string        `json:"description"`
	CreatedAt   time.Time     `json:"created_at"`
}
