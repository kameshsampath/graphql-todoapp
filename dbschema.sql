DROP TABLE IF EXISTS users CASCADE;
CREATE TABLE IF NOT EXISTS users
(
    id   SERIAL PRIMARY KEY,
    name text NOT NULL,
    sex varchar(6) NOT NULL,
    modifiedAt timestamp default now()
);
DROP TABLE IF EXISTS todos;
CREATE TABLE IF NOT EXISTS todos
(
    id     SERIAL PRIMARY KEY,
    text   text NOT NULL,
    done   bool default false,
    userId int  not null,
    modifiedAt timestamp default now()
);

ALTER TABLE todos DROP CONSTRAINT  IF EXISTS  fk_userid;
ALTER TABLE todos ADD CONSTRAINT fk_userid FOREIGN KEY (userId) REFERENCES users (id);

