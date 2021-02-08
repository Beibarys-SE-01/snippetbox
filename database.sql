CREATE TABLE snippets(
     id SERIAL PRIMARY KEY,
     title varchar(100) NOT NULL,
     content text NOT NULL,
     created date NOT NULL,
     expires date NOT NULL
);

INSERT INTO snippets (title, content, created, expires)
VALUES ('Test 1', 'Testing...',
current_date, current_date + INTERVAL '365 day');

INSERT INTO snippets (title, content, created, expires)
VALUES ('Test 2', 'Testing...',
current_date,current_date + INTERVAL '365 day');

INSERT INTO snippets (title, content, created, expires)
VALUES ('Test 3', 'Testing...',
current_date, current_date + INTERVAL '365 day');

CREATE USER snippet;
GRANT SELECT, INSERT ON snippets TO snippet;
ALTER USER snippet WITH PASSWORD 'snippet';

GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA public TO snippet;