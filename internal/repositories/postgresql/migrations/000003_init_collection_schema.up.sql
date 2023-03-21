CREATE TABLE "categories"
(
    "id"   SERIAL PRIMARY KEY,
    "name" VARCHAR UNIQUE NOT NULL
);

CREATE TABLE "collections"
(
    "token"       CHAR(42) NOT NULL,

    "owner"       CHAR(42) NOT NULL,
    "name"        VARCHAR  NOT NULL,
    "description" VARCHAR  NOT NULL,
    "metadata"    jsonb,
    "category"    SERIAL   NOT NULL,

    "created_at"  TIMESTAMP DEFAULT current_timestamp,

    PRIMARY KEY ("token"),
    FOREIGN KEY ("category") REFERENCES "categories" ("id")
);

INSERT INTO categories("name")
VALUES ('Photography'),
       ('Gaming'),
       ('Art'),
       ('Memberships');
