CREATE TABLE "users" (
                         "username" varchar PRIMARY KEY,
                         "hashed_password" varchar NOT NULL,
                         "full_name" varchar NOT NULL,
                         "email" varchar UNIQUE NOT NULL,
                         "password_changed_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00',
                         "created_at" timestamptz NOT NULL DEFAULT (NOW())
);

ALTER TABLE "accounts" ADD FOREIGN KEY ("owner") REFERENCES "users" ("username") ON DELETE CASCADE;

ALTER TABLE "accounts" ADD CONSTRAINT "owner_currency_unique" UNIQUE ("owner", "currency");