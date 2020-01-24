
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

create table users(
    id int primary key auto_increment,
    name varchar(255) not null,
    token varchar(255) not null unique
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

