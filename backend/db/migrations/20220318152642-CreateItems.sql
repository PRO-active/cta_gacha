
-- +migrate Up
create table if not exists items (
  id varchar(20) NOT NULL primary key,
  gacha_id char(3),
  index gcid_index(gacha_id),
  foreign key fk_gacha(gacha_id) references gachas(id),
  user_id char(3),index usid_index(user_id),
  foreign key fk_user(user_id) references users(id),
  name varchar(10),
  img_path char(20),
  rarity ENUM('N','R','SR')
  );
-- +migrate Down
drop table if exists items;
