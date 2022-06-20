CREATE TABLE posts(
    id SERIAL NOT NULL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    title text NOT NULL,
    body text NOT NULL
);