CREATE TABLE IF NOT EXISTS users (
  id GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  name VARCHAR(30) NOT NULL,
  balance BIGINT NOT NULL DEFAULT 0 CHECK(balance >= 0),
  created_at TIMESTAMPTZ NOT NULL DEFAULT now()
)

CREATE TABLE IF NOT EXISTS customers (
  id GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  user_id BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  role TEXT NOT NULL DEFAULT 'customer'
)

CREATE TABLE IF NOT EXISTS executors (
  id GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  user_id BIGINT NOT NULL REFERENCES users(id)ON DELETE CASCADE,
  role TEXT NOT NULL DEFAULT 'executor'
)


CREATE TABLE IF NOT EXISTS tasks (
  id GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  customer_id BIGINT NOT NULL REFERENCES customers(id) ON DELETE CASCADE,
  executor_id BIGINT NOT NULL REFERENCES executors(id) ON DELETE SET NULL,
  title VARCHAR(40) NOT NULL,
  description TEXT NOT NULL,
  status TEXT NOT NULL CHECK (status IN ('published', 'in_progress', 'canceled', 'on_review', 'completed', 'revision'))
  deadline DATE NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
)

CREATE TABLE IF NOT EXISTS bids (
  id GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  task_id BIGINT NOT NULL REFERENCES tasks(id) ON DELETE CASCADE,
  executor_id BIGINT NOT NULL REFERENCES executors(id) ON DELETE CASCADE,
  message TEXT NOT NULL,
  status TEXT NOT NULL CHECK (status IN ('pending', 'rejected', 'accepted'))
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
)

CREATE TABLE transactions (
  id GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  task_id BIGINT NOT NULL REFERENCES tasks(id),
  from_customer BIGINT NOT NULL REFERENCES customers(id),
  to_executor BIGINT NOT NULL REFERENCES executors(id),
  amount BIGINT NOT NULL CHECK(amount >= 0),
  status TEXT NOT NULL CHECK ( status IN ('hold', 'paid', 'returned'))
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
)

CREATE FUNCTION trg_update_at()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = now()
  RETURN NEW;
END;
$$
LANGUAGE plpgsql



CREATE TRIGGER trg_tasks_update
BEFORE UPDATE ON tasks
FOR EACH ROW EXECUTE FUNCTION trg_update_at()


CREATE TRIGGER trg_bids_update
BEFORE UPDATE ON bids
FOR EACH ROW EXECUTE FUNCTION trg_update_at()


CREATE TRIGGER trg_transaction_update
BEFORE UPDATE ON transaction
FOR EACH ROW EXECUTE FUNCTION trg_update_at()


CREATE INDEX idx_transaction_task ON transactions(task_id)

CREATE INDEX idx_transaction_status ON transactions(status)

CREATE INDEX idx_bids_task ON bids(task_id)

CREATE INDEX idx_bids_status ON bids(status)

CREATE INDEX idx_tasks_cr_id ON tasks(customer_id)

CREATE INDEX idx_tasks_er_id ON tasks(executor_id)

CREATE INDEX idx_tasks_status ON tasks(status)