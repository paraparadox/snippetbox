create database snippetbox;

use snippetbox;

create table snippets
(
    id      integer      not null primary key auto_increment,
    title   varchar(100) not null,
    content text         not null,
    created datetime     not null,
    expires datetime     not null
);

create index idx_snippets_created on snippets (created);

create table sessions
(
    token  char(43) primary key,
    data   blob         not null,
    expiry timestamp(6) not null
);

create table users
(
    id integer not null primary key auto_increment,
    name varchar(255) not null,
    email varchar(255) not null,
    hashed_password char(60) not null,
    created datetime not null
);

alter table users add constraint users_uc_email unique (email);

-- use the source command in mysql cli to run these statements