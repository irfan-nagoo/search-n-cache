
CREATE table sc_article (
	id int primary key auto_increment,
	title varchar(200),
	description varchar(500),
	author varchar(100),
	article_category varchar(20),
	article_type varchar(20),
	content longblob,
	tags varchar(20),
	created_at timestamp,
	created_by varchar(50),
	updated_at timestamp,
	updated_by varchar(50)
);
