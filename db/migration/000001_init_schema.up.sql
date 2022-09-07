CREATE TABLE "users" (
                        "id" bigserial PRIMARY KEY NOT NULL,
                        "email" varchar UNIQUE NOT NULL,
                        "email_confirmation" varchar,
                        "password" varchar(60),
                        "created_at" timestamp NOT NULL DEFAULT (now()),
                        "updated_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "products" (
                            "id" varchar(7) PRIMARY KEY NOT NULL,
                            "estilo" varchar NOT NULL,
                            "piel" varchar NOT NULL,
                            "suela" varchar NOT NULL,
                            "ca" bool DEFAULT 'false',
                            "cantidad" integer DEFAULT 0,
                            "description" varchar,
                            "created_at" timestamp NOT NULL DEFAULT (now()),
                            "updated_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "sizes" (
                         "id" varchar PRIMARY KEY NOT NULL,
                         "24" int DEFAULT 0,
                         "24.5" int DEFAULT 0,
                         "25" int DEFAULT 0,
                         "25.5" int DEFAULT 0,
                         "26" int DEFAULT 0,
                         "26.5" int DEFAULT 0,
                         "27" int DEFAULT 0,
                         "27.5" int DEFAULT 0,
                         "28" int DEFAULT 0,
                         "28.5" int DEFAULT 0,
                         "29" int DEFAULT 0,
                         "29.5" int DEFAULT 0,
                         "30" int DEFAULT 0,
                         "30.5" int DEFAULT 0,
                         "31" int DEFAULT 0,
                         "created_at" timestamp NOT NULL DEFAULT (now()),
                         "updated_at" timestamp NOT NULL DEFAULT (now())
);

ALTER TABLE "sizes" ADD FOREIGN KEY ("id") REFERENCES "products" ("id");

CREATE INDEX sizes_id_index ON sizes ("id")