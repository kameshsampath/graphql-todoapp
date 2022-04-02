drop table  if exists "users";
create table "users"
(
  id          INTEGER not null
    primary key,
  name        VARCHAR not null,
  gender      VARCHAR not null,
  modified_at TIMESTAMP default 'current_timestamp' not null
);


drop table if exists "todos";
create table "todos"
(
  id          INTEGER not null
    primary key,
  text        VARCHAR not null,
  done        BOOLEAN   default 'false' not null,
  user_id     INTEGER not null,
  modified_at TIMESTAMP default 'current_timestamp' not null
);
