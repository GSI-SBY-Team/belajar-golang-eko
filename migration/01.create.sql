
create table user_eko (
	id varchar(36) primary key, 
	username varchar(100) unique,
	password varchar(50)
);
insert into user_eko values('123', 'eko', '123');
