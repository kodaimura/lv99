CREATE TABLE IF NOT EXISTS admin (
	admin_id SERIAL PRIMARY KEY,
	admin_name TEXT NOT NULL UNIQUE,
	admin_password TEXT NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

create trigger trg_admin_upd BEFORE UPDATE ON admin FOR EACH ROW
  execute procedure set_updated_at();