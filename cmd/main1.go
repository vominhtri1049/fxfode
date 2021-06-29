package main1

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Note struct {
	Id      int     `json:"id,omitempty" gorm:"column:id;"`
	Title   *string `json:"title" gorm:"column:title;"`
	Content string  `json:"content" gorm:"column:content;"`
}

func (Note) TableName() string {
	return "notes"
}

func main1() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := os.Getenv("DSN")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("Can not connect database")
	}

	// insert note
	// newNote := Note{Title: "Demo note", Content: "This is content of demo note"}
	// if result := db.Create(&newNote); result.Error != nil {
	// 	log.Fatalln("Insert failure")
	// }

	// var notes []Note
	// db.Where("status = ?", 1).Find(&notes)
	// fmt.Println(notes)

	// var note Note
	// if result := db.Where("id = ?", 7).First(&note); result.Error != nil {
	// 	log.Println(result.Error)
	// }
	// fmt.Println(note)

	// if result := db.Table(Note{}.TableName()).Where("id = ?", 5).Delete(nil); result.Error != nil {
	// 	log.Println(result.Error)
	// }
	// fmt.Println("Delete success!")

	// newTitle := ""
	// note.Title = &newTitle
	// if result := db.Where("id = ?", 7).Updates(&note); result.Error != nil {
	// 	log.Println(result.Error)
	// }
	// fmt.Println("Update success")
}
