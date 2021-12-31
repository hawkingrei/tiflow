drop database if exists `extend_column1`;
create database `extend_column1`;
use `extend_column1`;
create table t1 (c1 int, c2 int, c3 int, primary key(c1));
insert into t1 values(1, 1, 1);
insert into t1 values(2, 2, 2);
create table y1 (c1 int, c2 int, c3 int, gc3 int GENERATED ALWAYS AS (c1 + 1) VIRTUAL, primary key(c1));
insert into y1 (c1, c2, c3) values(1, 1, 1);
insert into y1 (c1, c2, c3) values(2, 2, 2);
create table y2 (c1 int, c2 int, c3 int, gc3 int GENERATED ALWAYS AS (c1 + 1) VIRTUAL, primary key(c1));
insert into y2 (c1, c2, c3) values(1, 1, 1);
insert into y2 (c1, c2, c3) values(2, 2, 2);