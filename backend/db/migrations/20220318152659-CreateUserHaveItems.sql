
-- +migrate Up
create table if not exists UserHaveItems (
  id char NOT NULL primary key,
  user_id char,
  index usid_index(user_id),
  foreign key fk_user(user_id) references Users(id),
  item_id char,index itid_index(item_id),
  foreign key fk_item(item_id) references Items(id)
  );
-- +migrate Down
drop table if exists UserHaveItems;