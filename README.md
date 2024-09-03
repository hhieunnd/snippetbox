# snippetbox

## Docker

```bash
docker run --name my-postgres -e POSTGRES_PASSWORD=It123456@ -d -p 5432:5432 postgres


CREATE TABLE snippets (
  id SERIAL PRIMARY KEY,
  title VARCHAR(255) NOT NULL,
  content TEXT NOT NULL,
  created TIMESTAMP WITH TIME ZONE,
  expires TIMESTAMP WITH TIME ZONE
);


CREATE TABLE sessions (
    token CHAR(43) PRIMARY KEY,
    data BYTEA NOT NULL,
    expiry TIMESTAMP(6) NOT NULL
);

CREATE INDEX sessions_expiry_idx ON sessions (expiry);

CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  email VARCHAR(255) NOT NULL,
  hashed_password CHAR(60) NOT NULL,
  created TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE users ADD CONSTRAINT users_uc_email UNIQUE (email);
