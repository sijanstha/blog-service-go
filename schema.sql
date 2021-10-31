drop database if exists db_blog;
create database db_blog;
use db_blog;

create table tbl_role(
    id varchar(100) not null primary key,
    role_name varchar(50) not null,
    display_name varchar(50) not null,
    active boolean default true,
    deleted boolean default false,
    created_at varchar(100) not null,
    updated_at varchar(100) not null,
    deleted_at varchar(100)
);

create table tbl_posts(
    id varchar(100) not null primary key,
    title varchar(200) not null,
    description longtext,
    active boolean default true,
    deleted boolean default false,
    created_at varchar(100) not null,
    updated_at varchar(100) not null,
    deleted_at varchar(100)
);

create table tbl_comments(
    id varchar(100) not null primary key,
    review varchar(300) not null,
    fk_post_id varchar(100) not null,
    active boolean default true,
    deleted boolean default false,
    created_at varchar(100) not null,
    updated_at varchar(100) not null,
    deleted_at varchar(100),
    foreign key (fk_post_id) references tbl_posts(id)
);

create table tbl_users(
    id varchar(100) not null primary key,
    first_name varchar(50) not null,
    last_name varchar(50) not null,
    email varchar(100) unique not null,
    password_hash varchar(255) not null,
    profile_picture_url varchar(255),
    fk_role_id varchar(100) not null,
    active boolean default true,
    deleted boolean default false,
    created_at varchar(100) not null,
    updated_at varchar(100) not null,
    deleted_at varchar(100),
    foreign key (fk_role_id) references tbl_role(id)
);

