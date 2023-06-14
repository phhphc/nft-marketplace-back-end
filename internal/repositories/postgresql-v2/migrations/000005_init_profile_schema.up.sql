CREATE TABLE "profiles" (
    "address" CHAR(42) NOT NULL,
    "username" VARCHAR(255),
    "metadata" JSONB,
    "signature" VARCHAR NOT NULL,

    PRIMARY KEY ("address")
);