ALTER TABLE IF EXISTS "accounts" DROP CONSTRAINT "owner_currency_unique";

ALTER TABLE IF EXISTS "accounts" DROP CONSTRAINT "accounts_owner_fkey";

DROP TABLE IF EXISTS "users";