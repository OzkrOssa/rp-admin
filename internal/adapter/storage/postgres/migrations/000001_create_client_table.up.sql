CREATE TABLE "client" (
  "id" INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "create_at" timestamp DEFAULT (now()),
  "update_at" timestamp,
  "delete_at" timestamp,
  "first_name" varchar NOT NULL,
  "last_name" varchar,
  "client_type" int NOT NULL,
  "document" varchar NOT NULL,
  "email" varchar NOT NULL,
  "profession" varchar,
  "address" varchar NOT NULL,
  "precint" varchar NOT NULL,
  "municipality" int NOT NULL,
  "department" int NOT NULL,
  "status" int NOT NULL DEFAULT 1,
  "latitude" varchar,
  "longitude" varchar,
  "nap" varchar,
  "access_point" varchar
);

CREATE TABLE "client_type" (
  "id" INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "name" varchar NOT NULL
);

CREATE TABLE "phone_numbers" (
  "id" INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "client_id" int,
  "number" varchar NOT NULL,
  "create_at" timestamp DEFAULT (now()),
  "update_at" timestamp,
  "delete_at" timestamp
);

CREATE TABLE "client_status" (
  "id" INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "status" varchar NOT NULL,
  "create_at" timestamp DEFAULT (now()),
  "update_at" timestamp,
  "delete_at" timestamp
);

CREATE TABLE "municipality" (
  "code" int PRIMARY KEY,
  "name" varchar NOT NULL,
  "create_at" timestamp DEFAULT (now()),
  "update_at" timestamp,
  "delete_at" timestamp
);

CREATE TABLE "department" (
  "code" int PRIMARY KEY,
  "name" varchar NOT NULL,
  "create_at" timestamp DEFAULT (now()),
  "update_at" timestamp,
  "delete_at" timestamp
);

ALTER TABLE "client" ADD FOREIGN KEY ("client_type") REFERENCES "client_type" ("id");

ALTER TABLE "client" ADD FOREIGN KEY ("municipality") REFERENCES "municipality" ("code");

ALTER TABLE "client" ADD FOREIGN KEY ("department") REFERENCES "department" ("code");

ALTER TABLE "client" ADD FOREIGN KEY ("status") REFERENCES "client_status" ("id");

ALTER TABLE "phone_numbers" ADD FOREIGN KEY ("client_id") REFERENCES "client" ("id");