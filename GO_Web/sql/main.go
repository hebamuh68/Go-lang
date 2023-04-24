package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

// Student data type for export
type Student struct {
	ID        int
	FirstName string
	LastName  string
}

/*tpl and db are defined as global variables. tpl is an instance of 
the *template.Template struct, which holds the parsed HTML templates. 
db is a *sql.DB object, which is used to connect to the MySQL database.
*/
var tpl *template.Template
var db *sql.DB

func main() {

    //parses all the HTML templates in the templates folder using template.ParseGlob()
	tpl, _ = template.ParseGlob("templates/*.html")
	var err error

    //sets up a connection to the MySQL database using sql.Open()
	db, err = sql.Open("mysql", "heba:heba333@tcp(localhost:3306)/grade_1")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

    //sets up an HTTP server using http.HandleFunc(), which maps incoming HTTP requests to the StudentSearchHandler
	http.HandleFunc("/", StudentSearchHandler)
    http.HandleFunc("/insert", StudentInsertHandler)

	http.ListenAndServe("localhost:8080", nil)
}

func StudentSearchHandler(w http.ResponseWriter, r *http.Request) {

    //checks if the HTTP request is a GET or POST request using r.Method
	if r.Method == "GET" {

        //If it is a GET request, it renders the studentsearch.html template with no data by passing nil to tpl.ExecuteTemplate(). 
		tpl.ExecuteTemplate(w, "studentsearch.html", nil)
		return
	}

    //If it is a POST request, it retrieves the firstname parameter from the HTTP request using r.FormValue()
	r.ParseForm()
	firstname := r.FormValue("firstname")

    //constructs a SQL query to search for the student with that firstname, and executes the query using db.QueryRow()
	stmt := "SELECT * FROM students WHERE firstname = ?;"
	row := db.QueryRow(stmt, firstname)

    //If there is a result, it scans the result into a Student object using row.Scan()
	var s Student
	err := row.Scan(&s.ID, &s.FirstName, &s.LastName)
	if err != nil {
	}

    //and renders the studentsearch.html template with the data by passing the Student object to tpl.ExecuteTemplate()
	tpl.ExecuteTemplate(w, "studentsearch.html", s)
}

func StudentInsertHandler(w http.ResponseWriter, r *http.Request) {
    
    fmt.Println("*****insertHandler running*****")
	if r.Method == "GET" {
		tpl.ExecuteTemplate(w, "insert.html", nil)
		return
	}
	r.ParseForm()
	// func (r *Request) FormValue(key string) string
	ID := r.FormValue("ID")
	FirstName := r.FormValue("FirstName")
	LastName := r.FormValue("LastName")
	var err error
	if ID == "" || FirstName == "" || LastName == "" {
		fmt.Println("Error inserting row:", err)
		tpl.ExecuteTemplate(w, "insert.html", "Error inserting data, please check all fields.")
		return
	}
	var ins *sql.Stmt
	// don't use _, err := db.Query()
	// func (db *DB) Prepare(query string) (*Stmt, error)
	ins, err = db.Prepare("INSERT INTO `grade_1`.`students` (`id`, `firstname`, `lastname`) VALUES (?, ?, ?);")
	if err != nil {
		panic(err)
	}
	defer ins.Close()
	// func (s *Stmt) Exec(args ...interface{}) (Result, error)
	res, err := ins.Exec(ID, FirstName, LastName)

	// check rows affectect???????
	rowsAffec, _ := res.RowsAffected()
	if err != nil || rowsAffec != 1 {
		fmt.Println("Error inserting row:", err)
		tpl.ExecuteTemplate(w, "insert.html", "Error inserting data, please check all fields.")
		return
	}
	lastInserted, _ := res.LastInsertId()
	rowsAffected, _ := res.RowsAffected()
	fmt.Println("ID of last row inserted:", lastInserted)
	fmt.Println("number of rows affected:", rowsAffected)
	tpl.ExecuteTemplate(w, "insert.html", "Product Successfully Inserted")
}