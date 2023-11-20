package config

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DB returns the database object
func DB() *gorm.DB {

	// host := os.Getenv("DB_HOST")
	// port := os.Getenv("DB_PORT")
	// dbname := os.Getenv("DB_NAME")
	// username := os.Getenv("DB_USER")
	// password := os.Getenv("DB_PASS")

	// psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
	// 	"password=%s dbname=%s sslmode=disable",
	// 	host, port, username, password, dbname)

	PSQL_connStr := "postgresql://arjnep000%40gmail.com:RA2wF6iqJNxa@ep-rapid-sky-61800086.us-east-2.aws.neon.tech/cnfs?sslmode=verify-full"

	db, err := gorm.Open(postgres.Open(PSQL_connStr), &gorm.Config{
		AllowGlobalUpdate: false,
		Logger:            logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		log.Fatal("Database error: ", err)
		panic(err)
	}
	log.Println("Database Connected")

	return db

}
