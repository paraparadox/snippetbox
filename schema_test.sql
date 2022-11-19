create database test_snippetbox character set utf8mb4 collate utf8mb4_unicode_ci;

create user 'test_web'@'localhost';
grant create,drop,alter,index,select,insert,update,delete on test_snippetbox.* to 'test_web'@'localhost';
alter user 'test_web'@'localhost' identified by 'pass';
