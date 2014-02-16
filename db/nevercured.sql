DROP TABLE IF EXISTS jft;

CREATE TABLE jft (
	id INTEGER PRIMARY KEY,
	month INTEGER,
	day INTEGER,
	title TEXT,
	excerpt TEXT,
	source TEXT,
	content TEXT,
	summary TEXT
);
