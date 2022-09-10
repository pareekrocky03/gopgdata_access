/* we are going to install a third party package named lib/pq.
Run the code below in terminal in the directory where go mod file is available.
To install the package.
go get -u github.com/lib/pq
*/
package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

/*
We are importing the database/sql package because we will be using it to connect
to our database, and we are importing the fmt package because we will want to use it
to construct a connection string that has all of the information required to connect to our database.
*/
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "anuj"
	dbname   = "iTrade"
)

func main() {
	//create a string titles psqlInfo which contains all of the information required to connect to our Postgres database.

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	/*
	   The sql.Open() function takes two arguments - a driver name, and a string that tells that driver how to connect to our database - and then returns a pointer to a sql.DB and an error.

	   If the error is not nil we are going to go ahead and panic because this means that we did something wrong. Most likely, we didn’t import the github.com/lib/pq package.
	*/
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//By calling db.Ping() we force our code to actually open up a connection to the database which will validate whether or not our connection string was 100% correct.
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
}
