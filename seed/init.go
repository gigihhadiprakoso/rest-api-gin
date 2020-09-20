package seed

import ( 
	"log"
	"io/ioutil"
	. "rest-api-gin/models"
	"github.com/jinzhu/gorm"
)

var files, err = ioutil.ReadDir("./seed")

func Load(db *gorm.DB){	
	/*		Create Table Brands			*/
	err = db.Debug().DropTableIfExists(&Brands{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&Brands{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	/*		Create Table Categories		*/
	err = db.Debug().DropTableIfExists(&Categories{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&Categories{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	/*		Create Table Units			*/
	err = db.Debug().DropTableIfExists(&Units{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&Units{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	/*		Create Table Unit Types			*/
	err = db.Debug().DropTableIfExists(&UnitTypes{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&UnitTypes{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	/*		Create Table Companies			*/
	err = db.Debug().DropTableIfExists(&Companies{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&Companies{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	/*		Create Table Products		*/
	err = db.Debug().DropTableIfExists(&Products{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&Products{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	/*		Create Table Tokens			*/
	err = db.Debug().DropTableIfExists(&Tokens{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&Tokens{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	/*		Create Table Users			*/
	err = db.Debug().DropTableIfExists(&Users{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&Users{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	/*		Create Table Warehouses			*/
	err = db.Debug().DropTableIfExists(&Warehouses{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&Warehouses{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	/*		Seeding			*/	
	for i, _ := range categories {
		err = db.Debug().Model(&Categories{}).Create(&categories[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
	}

	for i, _ := range unit_types {
		err = db.Debug().Model(&UnitTypes{}).Create(&unit_types[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
	}

	for i, _ := range units {
		err = db.Debug().Model(&Units{}).Create(&units[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
	}

	/*err = db.Debug().Model(&Products{}).AddForeignKey("category_id", "categories(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}*/

    /*if err != nil {
        log.Fatal(err)
	}*/

	/*predata["categories"]=categories
	predata["products"]=products

	tables["categories"]=&Categories{}
	tables["products"]=&Products{}*/


	// for _, f := range files {
	// 	variable := strings.Replace(f.Name(),".go","",-1)
	// 	if variable != "init" {

			// json_data := reflect.ValueOf(predata[variable])
			// fmt.Println(reflect.TypeOf(json_data))
			
			/*err := db.Debug().DropTableIfExists(tables[variable]).Error
			if err != nil {
				log.Fatalf("cannot drop table: %v", err)
			}
			err = db.Debug().AutoMigrate(tables[variable]).Error
			if err != nil {
				log.Fatalf("cannot migrate table: %v", err)
			}*/

			/*for i:=0; i < json_data.Len(); i++ {
				// fmt.Println(json_tables.Index(i))
				err = db.Debug().Model(tables[variable]).Create(json_data.Index(i)).Error
				if err != nil {
					log.Fatalf("cannot seed users table: %v", err)
				}
			}*/
	// 	}
	// }
}
