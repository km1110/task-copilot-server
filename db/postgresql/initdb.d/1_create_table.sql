CREATE TABLE IF NOT EXISTS "users" (
    "id"      UUID NOT NULL,
    "name"    VARCHAR(255) NOT NULL,
    "role_id" UUID NOT NULL,
    "status"  BOOLEAN DEFAULT 'true' NOT NULL,
    PRIMARY KEY("id")
);

CREATE TABLE IF NOT EXISTS "todos" (
    "id"          UUID NOT NULL,
    "user_id"     UUID NOT NULL,
    "name"        VARCHAR(255) NOT NULL,
    "target_date" TIMESTAMP NOT NULL,
    "done_date"   TIMESTAMP,
    "status"      BOOLEAN DEFAULT 'false' NOT NULL,
    PRIMARY KEY ("id"),
    FOREIGN KEY("user_id") REFERENCES "users"("id") 
);
