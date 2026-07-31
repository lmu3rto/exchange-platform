DROP TRIGGER IF EXISTS trg_users_update ON users;
DROP TRIGGER IF EXISTS trg_users_archive ON users;
DROP TRIGGER IF EXISTS trg_tasks_update ON tasks;
DROP TRIGGER IF EXISTS trg_tasks_status_check ON tasks;
DROP TRIGGER IF EXISTS trg_bids_update ON bids;
DROP TRIGGER IF EXISTS trg_transaction_update ON transactions;

DROP FUNCTION IF EXISTS fnc_update_at() CASCADE;
DROP FUNCTION IF EXISTS fnc_archive_user_on_delete() CASCADE;
DROP FUNCTION IF EXISTS fnc_check_task_status_update() CASCADE;


DROP TABLE IF EXISTS users;

DROP TABLE IF EXISTS deleted_users;

DROP TABLE IF EXISTS customers;

DROP TABLE IF EXISTS executors;

DROP TABLE IF EXISTS tasks;

DROP TABLE IF EXISTS bids;

DROP TABLE IF EXISTS transactions;

DROP TABLE IF EXISTS wallet_transactions;