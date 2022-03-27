
-- +migrate Up
create table if not exists gachas (
  id char(3) NOT NULL primary key,
  user_id char(3),
  index usid_index(user_id),
  foreign key fk_user(user_id) references users(id),
  name varchar(10),
  description char(100),
  hidden tinyint
  );
-- +migrate Down
drop table if exists gachas;