package models

var dropBrackets = `DROP TABLE IF EXISTS brackets`
var dropParticipants = `DROP TABLE IF EXISTS participants`
var dropSets = `DROP TABLE IF EXISTS sets`

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
			REFERENCES brackets(id) ON DELETE CASCADE
	)`

var createSets = `CREATE TABLE IF NOT EXISTS sets (
	id INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
	bracket_id INT,
	participantA_id INT,
	participantB_id INT,
	CONSTRAINT fk_bracket
		FOREIGN KEY(bracket_id)
			REFERENCES brackets(id) ON DELETE CASCADE,
	CONSTRAINT fk_participantA
		FOREIGN KEY(participantA_id)
			REFERENCES participants(id),
	CONSTRAINT fk_participantB
		FOREIGN KEY(participantB_id)
			REFERENCES participants(id)
	)`

var SchemaDropQueries = []string{
	dropParticipants,
	dropSets,
	dropBrackets,
}

var SchemaCreateQueries = []string{
	createBrackets,
	createParticipants,
	createSets,
}
