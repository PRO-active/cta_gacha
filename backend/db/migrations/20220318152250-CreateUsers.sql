
-- +migrate Up
create table if not exists users (
  id char(3) NOT NULL primary key,
  name varchar(5),
  password varchar(10),
  email varchar(20),
  hash char(128)
  );
-- +migrate Down
drop table if exists users;