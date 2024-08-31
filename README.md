# snippetbox


# Docker
docker run --name my-postgres -e POSTGRES_PASSWORD=It123456@ -d -p 5432:5432 postgres

# Table
CREATE TABLE snippets (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    created TIMESTAMP WITH TIME ZONE
    expires TIMESTAMP WITH TIME ZONE
);
