DROP TABLE IF EXISTS users;
CREATE TABLE IF NOT EXISTS users
(
    id   SERIAL PRIMARY KEY,
    name text NOT NULL,
    sex varchar(6) NOT NULL,
    modifiedAt timestamp default current_timestamp
);
DROP TABLE IF EXISTS todos;
CREATE TABLE IF NOT EXISTS todos
(
    id     SERIAL PRIMARY KEY,
    text   text NOT NULL,
    done   bool default false,
    userId int  not null,
    modifiedAt timestamp default current_timestamp,
    foreign key (userId) references users(id)
);

