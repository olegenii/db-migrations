CREATE TABLE "warehouses" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "warehouses" ("name");

ALTER TABLE "items"
    ADD "WarehouseID" bigserial,
        ADD CONSTRAINT "fk_warehouses_items" FOREIGN KEY ("WarehouseID") 
            REFERENCES "warehouses" ("id");
