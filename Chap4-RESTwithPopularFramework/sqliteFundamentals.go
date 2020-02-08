package main

import (
	"database/sql"
	"log"
	_ "github.com/mattn/go-sqlite3"
)

type Book struct{
	id int
	name string
	author string
}

func main() {
	db, err := sql.Open("sqlite3", "./books.db")
	
	//Create table
	statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS books (id INTEGER PRIMARY KEY, isbn INTEGER, author VARCHAR(64),  name VARCHAR(64) NULL)")
	
	if err != nil{
		log.Println("Error in creating table")
		} else {
			log.Println("Successfully created a table books")
		}
		statement.Exec()
		
		//Create 1 data
		statement, _ = db.Prepare("INSERT INTO books (name, author, isbn) VALUES (?,?,?)")
		statement.Exec("A Tale of Two Cities", "Charles Dicken", 140430547)
		log.Println("Inserted the book into database!")
		//Read table 
		rows, _ := db.Query("SELECT id, name, author FROM books")
		var tempBook Book
		for rows.Next(){
			rows.Scan(&tempBook.id, &tempBook.name, &tempBook.author )
			log.Printf("ID:%d, Book:%s, Authors:%s\n", tempBook.id, tempBook.name, tempBook.author)
			log.Println("I am able to read a data")
		}
		//Update table
		statement, _ = db.Prepare("UPDATE books set name=? where id=?")
		statement.Exec("A Tale of Two Cities", 1)
		log.Println("Successfully updated a table books")
		
		
		//Delete table
		statement, _ = db.Prepare("DELETE from books WHERE id=?")
		statement.Exec(1)
		log.Println("Successfully deleted a table books")
		
	}