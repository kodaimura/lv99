CREATE TABLE IF NOT EXISTS chat (
  id BIGSERIAL PRIMARY KEY,
  from_id INTEGER NOT NULL,
  to_id INTEGER NOT NULL,
  message TEXT NOT NULL,
  is_read BOOLEAN NOT NULL DEFAULT false,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deleted_at TIMESTAMP NULL
);