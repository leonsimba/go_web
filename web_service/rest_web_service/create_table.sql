drop database gwp;
create database gwp;
use gwp;

drop user gwp@localhost;
create user 'gwp'@'localhost' identified by 'qwe123';
grant all on *.* to gwp@localhost identified by 'qwe123';

create table posts (
  id      serial primary key,
  content text,
  author  varchar(255)
);
