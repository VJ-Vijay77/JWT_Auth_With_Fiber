package schema


var User = `CREATE TABLE users(
	Id serial primary key,
	name text,
	age integer,
	role text);`