CREATE TABLE IF NOT EXISTS "users" (
    "id"         UUID NOT NULL,
    "role_id"    UUID NOT NULL,
    "name"       VARCHAR(255) NOT NULL,
    "is_active"  BOOLEAN DEFAULT 'true' NOT NULL,
    PRIMARY KEY("id")
    FOREIGN KEY("role_id") REFERENCES "roles"("id")
);

CREATE TABLE IF NOT EXISTS "todos" (
    "id"                UUID NOT NULL,
    "user_id"           UUID NOT NULL,
    "name"              VARCHAR(255) NOT NULL,
    "target_date"       TIMESTAMP NOT NULL,
    "done_date"         TIMESTAMP NOT NULL,
    "is_completed"      BOOLEAN DEFAULT 'false' NOT NULL,
    PRIMARY KEY ("id"),
    FOREIGN KEY("user_id") REFERENCES "users"("id") 
);

CREATE TABLE IF NOT EXISTS "calendar" (
    "id"            UUID NOT NULL,
    "calendar_name" VARCHAR(255) NOT NULL,
    "created_at"    TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at"    TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY("id")
)

CREATE TABLE IF NOT EXISTS "events" (
    "id"            UUID NOT NULL,
    "user_id"       UUID NOT NULL,
    "tag_id"        UUID NOT NULL,
    "title"         VARCHAR(255) NOT NULL,
    "start_date"    TIMESTAMP NOT NULL,
    "end_date"      TIMESTAMP NOT NULL,
    "created_at"    TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at"    TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY("id"),
    FOREIGN KEY("user_id") REFERENCES "users"("id")
);

CREATE TABLE IF NOT EXISTS "roles" (
    "id"            UUID NOT NULL,
    "role_name"     VARCHAR(255) NOT NULL,
    "created_at"    TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at"    TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY("id")
);

CREATE TABLE IF NOT EXISTS "tags" (
    "id"            UUID NOT NULL,
    "user_id"       UUID NOT NULL,
    "tag_name"      VARCHAR(255) NOT NULL,
    "tag_color"     VARCHAR(255) NOT NULL,
    "created_at"    TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at"    TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY("id")
);

CREATE TABLE IF NOT EXISTS "user_calendar" (
    "id"            UUID NOT NULL,
    "user_id"       UUID NOT NULL,
    "calendar_id"   UUID NOT NULL,
    "access_level"  enum('read', 'write', 'admin') NOT NULL,
    "created_at"    TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at"    TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY("id"),
    FOREIGN KEY("user_id") REFERENCES "users"("id"),
    FOREIGN KEY("calendar_id") REFERENCES "calendar"("id")
);

CREATE TABLE IF NOT EXISTS "calendar_event" (
    "id"            UUID NOT NULL,
    "calendar_id"   UUID NOT NULL,
    "event_id"      UUID NOT NULL,
    "created_at"    TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at"    TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY("id"),
    FOREIGN KEY("calendar_id") REFERENCES "calendar"("id"),
    FOREIGN KEY("event_id") REFERENCES "events"("id")
);