CREATE FUNCTION refresh_updated_at_step1() RETURNS trigger AS
$$
BEGIN
  IF NEW.updated_at = OLD.updated_at THEN
    NEW.updated_at := NULL;
  END IF;
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;
    
CREATE FUNCTION refresh_updated_at_step2() RETURNS trigger AS
$$
BEGIN
  IF NEW.updated_at IS NULL THEN
    NEW.updated_at := OLD.updated_at;
  END IF;
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE FUNCTION refresh_updated_at_step3() RETURNS trigger AS
$$
BEGIN
  IF NEW.updated_at IS NULL THEN
    NEW.updated_at := CURRENT_TIMESTAMP;
  END IF;
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TABLE IF NOT EXISTS "roles" (
    "id"            UUID NOT NULL,
    "role_name"     VARCHAR(255) NOT NULL,
    "created_at"    TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at"    TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY("id")
);

CREATE TRIGGER refresh_roles_updated_at_step1
  BEFORE UPDATE ON roles FOR EACH ROW
  EXECUTE PROCEDURE refresh_updated_at_step1();
CREATE TRIGGER refresh_roles_updated_at_step2
  BEFORE UPDATE OF updated_at ON roles FOR EACH ROW
  EXECUTE PROCEDURE refresh_updated_at_step2();
CREATE TRIGGER refresh_roles_updated_at_step3
  BEFORE UPDATE ON roles FOR EACH ROW
  EXECUTE PROCEDURE refresh_updated_at_step3();

CREATE TABLE IF NOT EXISTS "users" (
    "id"         UUID NOT NULL,
    "role_id"    UUID NOT NULL,
    "name"       VARCHAR(255) NOT NULL,
    "is_active"  BOOLEAN DEFAULT 'true' NOT NULL,
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY("id"),
    FOREIGN KEY("role_id") REFERENCES "roles"("id")
);

CREATE TRIGGER refresh_users_updated_at_step1
  BEFORE UPDATE ON users FOR EACH ROW
  EXECUTE PROCEDURE refresh_updated_at_step1();
CREATE TRIGGER refresh_users_updated_at_step2
  BEFORE UPDATE OF updated_at ON users FOR EACH ROW
  EXECUTE PROCEDURE refresh_updated_at_step2();
CREATE TRIGGER refresh_users_updated_at_step3
  BEFORE UPDATE ON users FOR EACH ROW
  EXECUTE PROCEDURE refresh_updated_at_step3();

CREATE TABLE IF NOT EXISTS "todos" (
    "id"                UUID NOT NULL,
    "user_id"           UUID NOT NULL,
    "name"              VARCHAR(255) NOT NULL,
    "target_date"       TIMESTAMPTZ NOT NULL,
    "done_date"         TIMESTAMPTZ NOT NULL,
    "is_completed"      BOOLEAN DEFAULT 'false' NOT NULL,
    "created_at"        TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at"        TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY ("id"),
    FOREIGN KEY("user_id") REFERENCES "users"("id") 
);

CREATE TRIGGER refresh_todos_updated_at_step1
  BEFORE UPDATE ON todos FOR EACH ROW
  EXECUTE PROCEDURE refresh_updated_at_step1();
CREATE TRIGGER refresh_todos_updated_at_step2
  BEFORE UPDATE OF updated_at ON todos FOR EACH ROW
  EXECUTE PROCEDURE refresh_updated_at_step2();
CREATE TRIGGER refresh_todos_updated_at_step3
  BEFORE UPDATE ON todos FOR EACH ROW
  EXECUTE PROCEDURE refresh_updated_at_step3();

CREATE TABLE IF NOT EXISTS "calendar" (
    "id"            UUID NOT NULL,
    "calendar_name" VARCHAR(255) NOT NULL,
    "created_at"    TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at"    TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY("id")
);

CREATE TRIGGER refresh_calendar_updated_at_step1
  BEFORE UPDATE ON calendar FOR EACH ROW
  EXECUTE PROCEDURE refresh_updated_at_step1();
CREATE TRIGGER refresh_calendar_updated_at_step2
  BEFORE UPDATE OF updated_at ON calendar FOR EACH ROW
  EXECUTE PROCEDURE refresh_updated_at_step2();
CREATE TRIGGER refresh_calendar_updated_at_step3
  BEFORE UPDATE ON calendar FOR EACH ROW
  EXECUTE PROCEDURE refresh_updated_at_step3();

CREATE TABLE IF NOT EXISTS "events" (
    "id"            UUID NOT NULL,
    "user_id"       UUID NOT NULL,
    "tag_id"        UUID NOT NULL,
    "title"         VARCHAR(255) NOT NULL,
    "start_date"    TIMESTAMPTZ NOT NULL,
    "end_date"      TIMESTAMPTZ NOT NULL,
    "created_at"    TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at"    TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY("id"),
    FOREIGN KEY("user_id") REFERENCES "users"("id")
);

CREATE TRIGGER refresh_events_updated_at_step1
  BEFORE UPDATE ON events FOR EACH ROW
  EXECUTE PROCEDURE refresh_updated_at_step1();
CREATE TRIGGER refresh_events_updated_at_step2
  BEFORE UPDATE OF updated_at ON events FOR EACH ROW
  EXECUTE PROCEDURE refresh_updated_at_step2();
CREATE TRIGGER refresh_events_updated_at_step3
  BEFORE UPDATE ON events FOR EACH ROW
  EXECUTE PROCEDURE refresh_updated_at_step3();

CREATE TABLE IF NOT EXISTS "tags" (
    "id"            UUID NOT NULL,
    "user_id"       UUID NOT NULL,
    "tag_name"      VARCHAR(255) NOT NULL,
    "tag_color"     VARCHAR(255) NOT NULL,
    "created_at"    TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at"    TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY("id")
);

CREATE TRIGGER refresh_tags_updated_at_step1
  BEFORE UPDATE ON tags FOR EACH ROW
  EXECUTE PROCEDURE refresh_updated_at_step1();
CREATE TRIGGER refresh_tags_updated_at_step2
  BEFORE UPDATE OF updated_at ON tags FOR EACH ROW
  EXECUTE PROCEDURE refresh_updated_at_step2();
CREATE TRIGGER refresh_tags_updated_at_step3
  BEFORE UPDATE ON tags FOR EACH ROW
  EXECUTE PROCEDURE refresh_updated_at_step3();

CREATE TYPE access_level AS ENUM ('read', 'write', 'admin');

CREATE TABLE IF NOT EXISTS "user_calendar" (
    "id"            UUID NOT NULL,
    "user_id"       UUID NOT NULL,
    "calendar_id"   UUID NOT NULL,
    "access_level"  access_level NOT NULL,
    "created_at"    TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at"    TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY("id"),
    FOREIGN KEY("user_id") REFERENCES "users"("id"),
    FOREIGN KEY("calendar_id") REFERENCES "calendar"("id")
);

CREATE TRIGGER refresh_user_calendar_updated_at_step1
  BEFORE UPDATE ON user_calendar FOR EACH ROW
  EXECUTE PROCEDURE refresh_updated_at_step1();
CREATE TRIGGER refresh_user_calendar_updated_at_step2
  BEFORE UPDATE OF updated_at ON user_calendar FOR EACH ROW
  EXECUTE PROCEDURE refresh_updated_at_step2();
CREATE TRIGGER refresh_user_calendar_updated_at_step3
  BEFORE UPDATE ON user_calendar FOR EACH ROW
  EXECUTE PROCEDURE refresh_updated_at_step3();

CREATE TABLE IF NOT EXISTS "calendar_event" (
    "id"            UUID NOT NULL,
    "calendar_id"   UUID NOT NULL,
    "event_id"      UUID NOT NULL,
    "created_at"    TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at"    TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY("id"),
    FOREIGN KEY("calendar_id") REFERENCES "calendar"("id"),
    FOREIGN KEY("event_id") REFERENCES "events"("id")
);

CREATE TRIGGER refresh_calendar_event_updated_at_step1
  BEFORE UPDATE ON calendar_event FOR EACH ROW
  EXECUTE PROCEDURE refresh_updated_at_step1();
CREATE TRIGGER refresh_calendar_event_updated_at_step2
  BEFORE UPDATE OF updated_at ON calendar_event FOR EACH ROW
  EXECUTE PROCEDURE refresh_updated_at_step2();
CREATE TRIGGER refresh_calendar_event_updated_at_step3
  BEFORE UPDATE ON calendar_event FOR EACH ROW
  EXECUTE PROCEDURE refresh_updated_at_step3();