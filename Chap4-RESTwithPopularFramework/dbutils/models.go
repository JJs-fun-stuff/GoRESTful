package dbutils

const train = `CREATE TABLE IF NOT EXISTS train(
	ID integer primary key autoincrement,
	DRIVER_NAME varchar(64) null,
	OPERATING_STATUS BOOLEAN
)`


const station = `CREATE TABLE IF NOT EXISTS station(
	ID integer primary key autoincrement,
	NAME VARCHAR(64) NULL,
	OPENING_TIME  TIME null,
	CLOSING_TIME TIME null
)`


const schedule = `CREATE TABLE IF NOT EXISTS schedule(
	ID integer primary key autoincrement,
	TRAIN_ID INT,
	STATION_ID INT,
	ARRIVAL_TIME TIME,
	FOREIGN KEY (TRAIN_ID) REFERENCES train(ID),
	FOREIGN KEY (STATION_ID) REFERENCES station(ID)
)`
