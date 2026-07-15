CREATE TABLE IF NOT EXISTS users (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    user_name VARCHAR(30) NOT NULL,
    balance NUMERIC(10, 2) NOT NULL DEFAULT 0.00,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at TIMESTAMPTZ NULL
);

CREATE TABLE IF NOT EXISTS deleted_users (
    id BIGINT PRIMARY KEY,
    user_name VARCHAR(30) NOT NULL,
    balance NUMERIC(10, 2) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL,
    deleted_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE OR REPLACE FUNCTION fnc_update_at() RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = now();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trg_users_update BEFORE UPDATE ON users
FOR EACH ROW EXECUTE FUNCTION fnc_update_at();

CREATE OR REPLACE FUNCTION fnc_archive_user_on_delete() RETURNS TRIGGER AS $$
BEGIN
    IF OLD.deleted_at IS NULL AND NEW.deleted_at IS NOT NULL THEN
        INSERT INTO deleted_users (id, user_name, balance, created_at, deleted_at)
        VALUES (OLD.id, OLD.user_name, OLD.balance, OLD.created_at, NEW.deleted_at);
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trg_users_archive AFTER UPDATE OF deleted_at ON users
FOR EACH ROW EXECUTE FUNCTION fnc_archive_user_on_delete();

CREATE TABLE IF NOT EXISTS customers (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    user_id BIGINT NOT NULL UNIQUE REFERENCES users (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS executors (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    user_id BIGINT NOT NULL UNIQUE REFERENCES users (id) ON DELETE CASCADE
);

CREATE OR REPLACE FUNCTION fnc_check_task_status_update() RETURNS TRIGGER AS $$
BEGIN
    IF OLD.status = NEW.status THEN
        RETURN NEW;
    END IF;
    IF NOT (
        (OLD.status = 'published' AND NEW.status IN ('in_progress', 'canceled')) OR
        (OLD.status = 'in_progress' AND NEW.status IN ('on_review', 'canceled', 'revision')) OR
        (OLD.status = 'revision' AND NEW.status = 'on_review') OR
        (OLD.status = 'on_review' AND NEW.status IN ('completed', 'revision'))
    ) THEN
        RAISE EXCEPTION 'Неверная последовательность статусов, статус % не может быть после %', NEW.status, OLD.status;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TABLE IF NOT EXISTS tasks (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    customer_id BIGINT NOT NULL REFERENCES customers (id),
    executor_id BIGINT REFERENCES executors (id),
    title VARCHAR(40) NOT NULL,
    description TEXT NOT NULL,
    status TEXT NOT NULL CHECK (
        status IN (
            'published',
            'in_progress',
            'canceled',
            'on_review',
            'completed',
            'revision'
        )
    ),
    accepted_bid_id BIGINT UNIQUE, 
    deadline TIMESTAMPTZ NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TRIGGER trg_tasks_update BEFORE UPDATE ON tasks
FOR EACH ROW EXECUTE FUNCTION fnc_update_at();

CREATE TRIGGER trg_tasks_status_check BEFORE UPDATE OF status ON tasks
FOR EACH ROW EXECUTE FUNCTION fnc_check_task_status_update();

CREATE TABLE IF NOT EXISTS bids (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    task_id BIGINT NOT NULL REFERENCES tasks (id) ON DELETE CASCADE,
    executor_id BIGINT NOT NULL REFERENCES executors (id),
    message TEXT NOT NULL,
    status TEXT NOT NULL CHECK (status IN ('pending', 'rejected', 'accepted')),
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TRIGGER trg_bids_update BEFORE UPDATE ON bids
FOR EACH ROW EXECUTE FUNCTION fnc_update_at();

CREATE TABLE IF NOT EXISTS transactions (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    task_id BIGINT NOT NULL REFERENCES tasks (id),
    from_customer BIGINT NOT NULL REFERENCES customers (id),
    to_executor BIGINT NOT NULL REFERENCES executors (id),
    amount NUMERIC(10, 2) NOT NULL,
    status TEXT NOT NULL CHECK (status IN ('hold', 'paid', 'returned')),
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TRIGGER trg_transaction_update BEFORE UPDATE ON transactions
FOR EACH ROW EXECUTE FUNCTION fnc_update_at();

CREATE TABLE IF NOT EXISTS user_balance (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES users (id),
    amount NUMERIC(10, 2) NOT NULL,
    operation_type TEXT NOT NULL CHECK (operation_type IN ('deposit', 'withdraw', 'payment', 'hold', 'refund')), -- noqa: LT05
    task_id BIGINT REFERENCES tasks (id),
    description TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);


ALTER TABLE tasks ADD CONSTRAINT fk_tasks_bid FOREIGN KEY (accepted_bid_id) REFERENCES bids (id) ON DELETE SET NULL NOT VALID;
ALTER TABLE tasks VALIDATE CONSTRAINT fk_tasks_bid;

CREATE INDEX CONCURRENTLY IF NOT EXISTS idx_transaction_task ON transactions (task_id);
CREATE INDEX CONCURRENTLY IF NOT EXISTS idx_transaction_status ON transactions (status);
CREATE INDEX CONCURRENTLY IF NOT EXISTS idx_bids_task ON bids (task_id);
CREATE INDEX CONCURRENTLY IF NOT EXISTS idx_bids_status ON bids (status);
CREATE INDEX CONCURRENTLY IF NOT EXISTS idx_tasks_cr_id ON tasks (customer_id);
CREATE INDEX CONCURRENTLY IF NOT EXISTS idx_tasks_er_id ON tasks (executor_id);
CREATE INDEX CONCURRENTLY IF NOT EXISTS idx_tasks_status ON tasks (status);
CREATE INDEX CONCURRENTLY IF NOT EXISTS idx_user_deleted ON users (deleted_at) WHERE deleted_at IS NULL;
