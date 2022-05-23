CREATE TABLE "items" (
  "id" bigserial PRIMARY KEY,
  "title" varchar NOT NULL,
  "price" decimal(12,2) NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "items" ("title");