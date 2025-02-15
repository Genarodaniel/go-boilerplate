BEGIN;

CREATE TABLE IF NOT EXISTS tab_user (
    id uuid DEFAULT gen_random_uuid(),
    name VARCHAR(250),
    email VARCHAR(250) NOT NULL,
    active SMALLINT NOT NULL DEFAULT 1,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);


CREATE INDEX idx_user_email ON tab_user(email);


CREATE TRIGGER trigger_tab_user_updated_at
BEFORE UPDATE ON tab_user
FOR EACH ROW EXECUTE FUNCTION update_timestamp();




COMMIT;