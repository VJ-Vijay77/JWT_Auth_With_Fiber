package schema


var User = `CREATE TABLE users(
	Id serial primary key,
	name text,
	email text unique,
	password text);`