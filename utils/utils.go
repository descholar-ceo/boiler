package utils

import (
	"fmt"
	"io"
	"os"
)

// AskRubocop is a function which asks a user if they will use rubocop and then returns the answer of the user
func AskRubocop() string {
	var isRubocop string
	fmt.Println("\nWill you use Rubocop as a linter? Enter y for yes or any other key for no")
	fmt.Scan(&isRubocop)
	return isRubocop
}

// AskGithub is a function which asks a user if they will use github
func AskGithub() string {
	var isGithub string
	fmt.Println("\nWill you use github as a collaboration tool? Enter y for yes or any other key for no")
	fmt.Scan(&isGithub)
	return isGithub
}

// AskDatabase is a function which asks a user which database they will use and returns the answer
func AskDatabase() string {
	var database string
	var tmpDb int
	fmt.Println("\nSelect Enter the number corresponding to the database you want to use: ")
	fmt.Println("\n1.sqlite3\n2.postgresql\n3.mysql\n4.oracle\n5.frontbase\n6.db2(ibm_db)\n7.sqlserver\n8.jdbcmysql\n9.jdbcpostgresql\n10.jdbcsqlite3\n11.jdbc")
	fmt.Scan(&tmpDb)
	switch tmpDb {
	case 1:
		database = "sqlite3"
	case 2:
		database = "postgresql"
	case 3:
		database = "mysql"
	case 4:
		database = "oracle"
	case 5:
		database = "frontbase"
	case 6:
		database = "ibm_db"
	case 7:
		database = "sqlserver"
	case 8:
		database = "jdbcmysql"
	case 9:
		database = "jdbcpostgresql"
	case 10:
		database = "jdbcsqlite3"
	case 11:
		database = "jdbc"
	default:
		fmt.Println("The database you choose is not supported by rails yet! So, I will create for you a rails app with a default database which is sqlite3")
		database = "sqlite3"
	}

	return database
}

// Copy function to copy files
func Copy(src, dst string) {
	source, _ := os.Open(src)

	defer source.Close()

	destination, _ := os.Create(dst)

	defer destination.Close()
	io.Copy(destination, source)
}

// WriteToFile is a function which used to write to file
func WriteToFile(file, stringToWrite string) {
	mFile, _ := os.OpenFile(file, os.O_APPEND|os.O_WRONLY, 0644)
	fmt.Fprintln(mFile, stringToWrite)
	mFile.Close()
}
