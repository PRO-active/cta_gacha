
-- +migrate Up
ALTER TABLE users MODIFY COLUMN name varchar(255);
ALTER TABLE users MODIFY COLUMN password varchar(255);
ALTER TABLE users MODIFY COLUMN email varchar(255);
ALTER TABLE users MODIFY COLUMN hash varchar(255);
ALTER TABLE gachas MODIFY COLUMN name varchar(255);
ALTER TABLE gachas MODIFY COLUMN description varchar(255);
ALTER TABLE items MODIFY COLUMN name varchar(255);
ALTER TABLE items MODIFY COLUMN img_path varchar(255);

-- +migrate Down
