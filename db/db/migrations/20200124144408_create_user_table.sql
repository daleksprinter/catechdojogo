
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

create table users(
    id int primary key auto_increment,
    name varchar(255) not null unique,
    token varchar(255) not null
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

