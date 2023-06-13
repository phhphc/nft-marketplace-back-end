CREATE TABLE IF NOT EXISTS "users" (
  public_address VARCHAR(42) NOT NULL,
  nonce VARCHAR NOT NULL,
  PRIMARY KEY (public_address)
);

CREATE TABLE IF NOT EXISTS "roles" (
  id SERIAL PRIMARY KEY,
  name VARCHAR NOT NULL
);

INSERT INTO "roles" (name) VALUES ('admin');
INSERT INTO "roles" (name) VALUES ('moderator');
INSERT INTO "roles" (name) VALUES ('user');

CREATE TABLE IF NOT EXISTS "user_roles" (
  address VARCHAR(42) NOT NULL,
  role_id INTEGER NOT NULL,
  PRIMARY KEY (address, role_id),
  FOREIGN KEY (address) REFERENCES "users" (public_address),
  FOREIGN KEY (role_id) REFERENCES "roles" (id)
);