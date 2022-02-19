CREATE TABLE payment (
	id bigserial primary key,
	name varchar(50) not null,
	price bigint not null,
	date varchar(20) not null,
	type varchar(10) not null,
	comment varchar(255),
	category_id bigint,
	foreign key (category_id)
			REFERENCES category(id)
)