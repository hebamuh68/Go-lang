package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type Student struct {
	id        int
	firstname string
	lastname  string
}

func main() {

	//================================== Connect to the MySQL database:
	db, err := sql.Open("mysql", "heba:heba333@tcp(localhost:3306)/grade_2")
	if err != nil {
		fmt.Printf("error")
	}
	defer db.Close()


	//================================== handles incoming requests to the root path ("/")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		
		
	//================================== Query the database for student records:
		entry, err := db.Query("select * from `grade_2`.`students`;")
		if err != nil {
			panic(err.Error())
		}
		defer entry.Close()


	//================================== Parse the student records and accumulate them into a string
		students := []Student{}
		for entry.Next() {
			var s Student
			err := entry.Scan(&s.id, &s.firstname, &s.lastname)
			if err != nil {
				panic(err)
			}
			students = append(students, s)
			
		}

		var sb strings.Builder
		for i, s := range students {
			sb.WriteString(fmt.Sprintf("%d %s %s\n", i+1, s.firstname, s.lastname))
		}

		// Convert the accumulated string in the StringBuilder to a string
		studentsString := sb.String()
		fmt.Fprintf(w, studentsString)
	})


	//================================== listen port
	fmt.Println("Slave_2 starting on port 8099")
	if err := http.ListenAndServe("127.0.0.1:8099", nil); err != nil {
		panic(err)
	}
}
