package seed

import (
	// "log"
	// "fmt"
	// "github.com/jinzhu/gorm"
	. "rest-api-gin/models"
)

var categories = []Categories{
	Categories{
		Name: "Makanan",
	},
	Categories{
		Name: "Minuman",
	},
	Categories{
		Name: "Perabotan Rumah Tangga",
	},
}

/*func Load(db *gorm.DB){
	err := db.Debug().DropTableIfExists(&Categories{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&Categories{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	// err = db.Debug().Model(&models.Categories{}).AddForeignKey("author_id", "users(id)", "cascade", "cascade").Error
	// if err != nil {
	// 	log.Fatalf("attaching foreign key error: %v", err)
	// }

	for i, _ := range categories {
		fmt.Println(&categories[i])
		// err = db.Debug().Model(&Categories{}).Create(&categories[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
	// 	posts[i].AuthorID = users[i].ID

	// 	err = db.Debug().Model(&models.Post{}).Create(&posts[i]).Error
	// 	if err != nil {
	// 		log.Fatalf("cannot seed posts table: %v", err)
	// 	}
	}
}*/