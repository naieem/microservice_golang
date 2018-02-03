package app

import (
  "database/sql"
  _ "github.com/go-sql-driver/mysql" // mysql
  "log"
)

type Instruction struct {
	Title string 
	Description string 
}
var dbmap = initDb()
   
func initDb() *sql.DB {
//   db, err := sql.Open("mysql", "root:@/test")
//   checkErr(err, "sql.Open failed")

//   dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
// //   dbmap.AddTableWithName(Instruction{}, "Instruction").SetKeys(true, "Id")

//   err = dbmap.CreateTablesIfNotExists()
//   checkErr(err, "Create table failed")
//   log.Println(dbmap)
//   return dbmap
		dbDriver := "mysql"   // Database driver
		dbUser := "root"      // Mysql username
		dbPass := "" // Mysql password
		dbName := "test"   // Mysql schema

		// Realize the connection with mysql driver
		db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)

		// If error stop the application
		if err != nil {
			panic(err.Error())
		}

		// Return db object to be used by another functions
		return db
}

func checkErr(err error, msg string) {
  if err != nil {
    log.Fatalln(msg, err)
  }
}