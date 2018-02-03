package app

import (
  "database/sql"
  _ "github.com/go-sql-driver/mysql" // mysql
  "gopkg.in/gorp.v1"
  "log"
)

type Instruction struct {
	Title string
	Description string
}
var dbmap = initDb()
   
func initDb() *gorp.DbMap {
	dbDriver := "mysql"   // Database driver
		dbUser := "root"      // Mysql username
		dbPass := "" // Mysql password
		dbName := "test"   // Mysql schema
		db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
		checkErr(err, "sql.Open failed")

		dbmap := &gorp.DbMap{Db: db}
		//   dbmap.AddTableWithName(Instruction{}, "Instruction").SetKeys(true, "Id")

		err = dbmap.CreateTablesIfNotExists()
		checkErr(err, "Create table failed")
		log.Println(dbmap)
		return dbmap
		

		// Realize the connection with mysql driver
		// db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)

		// // If error stop the application
		// if err != nil {
		// 	panic(err.Error())
		// }

		// // Return db object to be used by another functions
		// return db
}

func checkErr(err error, msg string) {
  if err != nil {
	panic(err.Error())
    // log.Fatalln(msg, err)
  }
}