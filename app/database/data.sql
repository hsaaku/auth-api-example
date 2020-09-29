CREATE TABLE "users" (
  "id" SERIAL PRIMARY KEY,
  "name" VARCHAR(30) NOT NULL,
  "display_name" VARCHAR(45) NOT NULL,
  "password" VARCHAR(255) NOT NULL
);

CREATE TABLE "roles" (
  "id" SERIAL PRIMARY KEY,
  "name" VARCHAR(45) NOT NULL
);

CREATE TABLE "user_roles" (
  "user_id" INT NOT NULL,
  "role_id" INT NOT NULL,
  PRIMARY KEY ("user_id", "role_id"),
  CONSTRAINT "fk_user_roles_users"
    FOREIGN KEY ("user_id")
      REFERENCES "users"("id")
      ON DELETE CASCADE,
  CONSTRAINT "fk_user_roles_roles"
    FOREIGN KEY ("role_id")
      REFERENCES "roles"("id")
      ON DELETE CASCADE
);

INSERT INTO "users"
  ("name", "display_name", "password")
VALUES
  ('beni39', 'Beni Suroto', 'rahasia123'),
  ('ferimahendra', 'Feri Mahendra', 'F9Hi3k8_FKi3');

INSERT INTO "roles"
  ("name")
VALUES
  ('Mahasiswa'),
  ('Dosen Tetap'),
  ('Teller');

INSERT INTO "user_roles"
  ("user_id", "role_id")
VALUES
  (1, 2),
  (1, 3),
  (2, 1);
