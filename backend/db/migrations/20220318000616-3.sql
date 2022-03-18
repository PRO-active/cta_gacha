
-- +migrate Up
create table if not exists Items (id char primary key, gacha_id char,index gcid_index(gacha_id), foreign key fk_gacha(gacha_id) references Gachas(id), user_id char,index usid_index(user_id), foreign key fk_user(user_id) references Users(id), name char, img_path char, rarity ENUM('N','R','SR'));
-- +migrate Down
drop table if exists Items;