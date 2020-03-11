package main    
import (
  "log"
  "github.com/JJs-fun-stuff/Chap7-PostGreSQL/models"
)


func main() {
  db, err := models.InitDB()
  if err != nil {
    log.Println(db)
  }
}