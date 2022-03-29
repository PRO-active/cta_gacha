
-- +migrate Up
alter table gachas drop foreign key gachas_ibfk_1;
alter table items drop foreign key items_ibfk_1;
alter table items drop foreign key items_ibfk_2;
alter table userhaveitems drop foreign key userhaveitems_ibfk_1;
alter table userhaveitems drop foreign key userhaveitems_ibfk_2;

ALTER TABLE users MODIFY COLUMN id varchar(255);
ALTER TABLE gachas MODIFY COLUMN id varchar(255);
ALTER TABLE gachas MODIFY COLUMN user_id varchar(255);
ALTER TABLE items MODIFY COLUMN id varchar(255);
ALTER TABLE items MODIFY COLUMN gacha_id varchar(255);
ALTER TABLE items MODIFY COLUMN user_id varchar(255);
ALTER TABLE userhaveitems MODIFY COLUMN id varchar(255);
ALTER TABLE userhaveitems MODIFY COLUMN user_id varchar(255);
ALTER TABLE userhaveitems MODIFY COLUMN item_id varchar(255);

-- +migrate Down
ALTER TABLE users MODIFY COLUMN id varchar(20);
ALTER TABLE gachas MODIFY COLUMN id varchar(20);
ALTER TABLE gachas MODIFY COLUMN user_id char(3);
ALTER TABLE items MODIFY COLUMN id varchar(20);
ALTER TABLE items MODIFY COLUMN gacha_id char(3);
ALTER TABLE items MODIFY COLUMN user_id char(3);
ALTER TABLE userhaveitems MODIFY COLUMN id varchar(20);
ALTER TABLE userhaveitems MODIFY COLUMN user_id char(3);
ALTER TABLE userhaveitems MODIFY COLUMN item_id char(3);