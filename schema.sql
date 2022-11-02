create table snippets (
    id integer not null primary key auto_increment,
    title varchar(100) not null,
    content text not null,
    created datetime not null,
    expires datetime not null
);

create index idx_snippets_created on snippets(created);

-- use the source command in mysql cli to run these statements