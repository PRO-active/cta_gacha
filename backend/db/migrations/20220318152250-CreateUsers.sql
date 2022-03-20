
-- +migrate Up
create table if not exists Users (
  id char NOT NULL primary key,
  name char,
  password varchar(10),
  email char,
  hash varchar(10)
  );
-- +migrate Down
drop table if exists Users;