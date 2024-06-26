-- Create "users" table
CREATE TABLE "users" ("id" bigint NOT NULL GENERATED BY DEFAULT AS IDENTITY, "login" character varying NOT NULL, "password" character varying NOT NULL, "first_name" character varying NULL, "last_name" character varying NULL, "role" character varying NOT NULL, "phonenumber" character varying NULL, "email" character varying NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NULL, "token" character varying NOT NULL DEFAULT '', "refresh_token" character varying NOT NULL DEFAULT '', PRIMARY KEY ("id"));
-- Create index "users_email_key" to table: "users"
CREATE UNIQUE INDEX "users_email_key" ON "users" ("email");
-- Create index "users_login_key" to table: "users"
CREATE UNIQUE INDEX "users_login_key" ON "users" ("login");
-- Create index "users_phonenumber_key" to table: "users"
CREATE UNIQUE INDEX "users_phonenumber_key" ON "users" ("phonenumber");
