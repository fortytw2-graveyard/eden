-- +goose Up
CREATE TABLE users (
    id           SERIAL PRIMARY KEY,
    username     TEXT NOT NULL UNIQUE,
    email        TEXT NOT NULL UNIQUE,
    passwordhash TEXT,
    
    banned       BOOLEAN DEFAULT false,
    admin        BOOLEAN DEFAULT false,
    confirmed    BOOLEAN DEFAULT false,

    created_at   TIMESTAMP DEFAULT now()
);

CREATE TABLE boards (
    id           SERIAL PRIMARY KEY,
    name         TEXT UNIQUE,
    creator_name TEXT,
    mod_names    TEXT[],
    summary      TEXT NOT NULL,
    deleted      BOOLEAN DEFAULT false,
    approved     BOOLEAN DEFAULT false,

    created_at TIMESTAMP DEFAULT now()
);

CREATE TABLE posts (
    id         SERIAL PRIMARY KEY,
    board_id   INT references boards (id),
    op_id      INT references users (id),
    title      TEXT NOT NULL,
    link       TEXT NOT NULL,
    body       TEXT NOT NULL,

    created_at TIMESTAMP DEFAULT now()
);

CREATE TABLE comments (
    id          SERIAL PRIMARY KEY,
    post_id     INT references posts (id),
    comment_id  INT references comments (id),

    op_id       INT references users (id),
    op_name     TEXT,
    op_admin    BOOLEAN,
    body        TEXT,

    created_at  TIMESTAMP DEFAULT now()
);

-- +goose Down
DROP TABLE users, boards, posts, comments CASCADE;
