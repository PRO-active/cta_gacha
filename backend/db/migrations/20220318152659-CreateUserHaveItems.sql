
-- +migrate Up
create table if not exists userhaveitems (
  id varchar(20) NOT NULL primary key,
  user_id char(3),
  index usid_index(user_id),
  foreign key fk_user(user_id) references users(id),
  item_id char(3),index itid_index(item_id),
  foreign key fk_item(item_id) references items(id)
  );
-- +migrate Down
drop table if exists userhaveitems;
