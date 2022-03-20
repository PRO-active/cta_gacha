
-- +migrate Up
create table if not exists Gachas (
  id char NOT NULL primary key,
  user_id char,
  index usid_index(user_id),
  foreign key fk_user(user_id) references Users(id),
  name char,
  description char,
  del tinyint
  );
-- +migrate Down
drop table if exists Gachas;