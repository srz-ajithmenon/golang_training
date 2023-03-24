package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
)

type Student struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	RollNo    string `json:"roll_no"`
	Age       int    `json:"age"`
	School    string `json:"school"`
}

func setupDB() *sql.DB {
	connection_str := "postgres://postgres:test@localhost:5433/ajithmenon?sslmode=disable"
	db, err := sql.Open("postgres", connection_str)
	if err != nil {
		fmt.Printf(" Database Connection Error %v \n", err)
	}
	if err := db.Ping(); err != nil {
		fmt.Println("Ping Error ", err)
	}
	return db
}

func GetStudents(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	var students []Student

	db := setupDB()

	query := "SELECT * FROM STUDENT"

	rows, er := db.Query(query)
	if er != nil {
		fmt.Printf(" Query Error %v \n", er)
	}

	for rows.Next() {
		var id int
		var age int
		var first_name string
		var last_name string
		var roll_no string
		var school string

		if err := rows.Scan(&id, &first_name, &last_name, &roll_no, &age, &school); err != nil {
			fmt.Printf(" Row Error %v \n", err)
		}
		// fmt.Printf(" id : %d \n name : %s %s\n age : %d \n Roll.No : %s \n school : %s \n\n", id, first_name, last_name, age, roll_no, school)

		student := Student{ID: id, FirstName: first_name, LastName: last_name, RollNo: roll_no, Age: age, School: school}
		students = append(students, student)

	}

	// data, er := json.Marshal(students)
	// if er != nil {
	// 	fmt.Println("Error in Data", er)
	// }

	// w.Write(data)

	json.NewEncoder(w).Encode(students)

}

func AddStudent(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var student *Student
	json.NewDecoder(r.Body).Decode(&student)

	db := setupDB()

	query := "INSERT INTO STUDENT(first_name, last_name, roll_no, age, school) VALUES ($1, $2, $3, $4, $5)"

	_, er := db.Query(query, student.FirstName, student.LastName, student.RollNo, student.Age, student.School)
	if er != nil {
		fmt.Printf(" Query Error %v \n", er)
	} else {
		w.Write([]byte("Successfully Added Student : "))
		json.NewEncoder(w).Encode(student)
	}

}

func UpdateStudent(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var student *Student
	json.NewDecoder(r.Body).Decode(&student)

	db := setupDB()

	query := "UPDATE STUDENT SET last_name = $1 WHERE roll_no = $2"

	_, er := db.Query(query, student.LastName, student.RollNo)
	if er != nil {
		fmt.Printf(" Query Error %v \n", er)
	} else {
		w.Write([]byte("Successfully Updated Student"))
	}

}

func DeleteStudent(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	var student *Student
	json.NewDecoder(r.Body).Decode(&student)

	db := setupDB()

	query := "DELETE from STUDENT WHERE roll_no = $1"

	_, er := db.Query(query, student.RollNo)
	// _, er := db.Query(query, ps.ByName("rollno"))

	if er != nil {
		fmt.Printf(" Query Error %v \n", er)
	} else {
		w.Write([]byte("Deleted the Student "))
	}

}

func main() {

	router := httprouter.New()
	router.GET("/allstudents", GetStudents)
	router.POST("/newstudent", AddStudent)
	router.PUT("/updatestudent", UpdateStudent)
	router.DELETE("/deletestudent", DeleteStudent)

	log.Fatal(http.ListenAndServe(":8080", router))
}
