CREATE DATABASE IF NOT EXISTS devbook;
USE devbook;

DROP TABLE IF EXISTS users;

CREATE TABLE users(
    id int auto_incremet primary key,
    name varchar(50) not null,
    nickName varchar(50) not null unique,
    email varchar(50) not null unique,
    password varchar(50) not null unique,
    createdAt timestamp default current_timestamp()
) ENGINE=INNODB;