package models

var dropBrackets = `DROP TABLE IF EXISTS brackets`
var dropParticipants = `DROP TABLE IF EXISTS participants`

var createBrackets = `CREATE TABLE IF NOT EXISTS brackets (
	id INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY, 
	name TEXT
	)`

var createParticipants = `CREATE TABLE IF NOT EXISTS participants (
	id INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
	name TEXT,
	bracket_id INT,
	CONSTRAINT fk_bracket
		FOREIGN KEY(bracket_id)
			REFERENCES brackets(id)
	)`

var SchemaDropQueries = []string{
	dropParticipants,
	dropBrackets,
}

var SchemaCreateQueries = []string{
	createBrackets,
	createParticipants,
}
