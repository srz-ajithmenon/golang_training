package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func GetSchools(db *sql.DB) {

	query := "SELECT * FROM SCHOOL"
	rows, er := db.Query(query)
	if er != nil {
		fmt.Printf(" Query Error %v \n", er)
	}

	for rows.Next() {
		var id int
		var name string
		var address string
		var phone int

		if err := rows.Scan(&id, &name, &address, &phone); err != nil {
			fmt.Printf(" Row Error %v \n", err)
		}

		fmt.Printf(" id : %d \n name : %s \n address : %s \n phone No : %d \n\n", id, name, address, phone)
	}

}

func NewSchool(db *sql.DB) {
	var name string
	var address string
	var phone int

	fmt.Print("Enter The Name : ")
	fmt.Scanln(&name)
	fmt.Print("Enter The Address : ")
	fmt.Scanln(&address)
	fmt.Print("Enter The Phone No. : ")
	fmt.Scanln(&phone)

	query := "INSERT INTO SCHOOL(name, address, phone_no) VALUES ($1, $2, $3 );"

	_, er := db.Query(query, name, address, phone)
	if er != nil {
		fmt.Printf(" Query Error %v \n", er)
	} else {
		fmt.Print("Inserted into School \n\n")
	}

}

func DeleteSchool(db *sql.DB) {

	var name string
	fmt.Print("Enter The Name of School to delete : ")
	fmt.Scanln(&name)

	query := "DELETE FROM SCHOOL WHERE  name = $1;"

	_, er := db.Query(query, name)
	if er != nil {
		fmt.Printf(" Query Error %v \n", er)
	} else {
		fmt.Print("Deleted From School \n\n")
	}

}

func UpdateSchool(db *sql.DB) {

	query := "UPDATE SCHOOL SET address = 'Mamabazar Pavaraty' WHERE name = 'Wisdom'"

	_, er := db.Query(query)
	if er != nil {
		fmt.Printf(" Query Error %v \n", er)
	} else {
		fmt.Print("Updated School \n\n")
	}

}

func tableSchool(db *sql.DB) {

	var ch_operations int

temp:
	for {
		fmt.Println(" 1. Read Data from Table \n 2. Insert Data into Table \n 3. Update data of Table \n 4. Delete data from Table")
		fmt.Print("\nEnter you Choice : ")
		fmt.Scanln(&ch_operations)

		switch ch_operations {
		case 1:
			GetSchools(db)
		case 2:
			NewSchool(db)
		case 3:
			UpdateSchool(db)
		case 4:
			DeleteSchool(db)
		default:
			fmt.Println("Invalid choice : Exit")
			break temp
		}

	}

}

func tableStudent(db *sql.DB) {

	var ch_operations int
	var query string

temp:
	for {

		fmt.Println("Enter your choice")
		fmt.Println(" 1. Read Data from Table \n 2. Insert Data into Table \n 3. Update data of Table \n 4. Delete data from Table")
		fmt.Scanln(&ch_operations)

		switch ch_operations {
		case 1:
			query = "SELECT * FROM STUDENT"
		case 2:
			query = "INSERT INTO STUDENT(first_name, last_name, roll_no, age, school) VALUES ('Akshay', 'p', '007', 25, 'Holy Grace');"
		case 3:
			query = "UPDATE STUDENT SET last_name = 'Paliyath' WHERE roll_no = '007';"
		case 4:
			query = "DELETE FROM STUDENT WHERE  roll_no = '007';"
		default:
			fmt.Println("Invalid choice : Exit")
			break temp
		}

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
			fmt.Printf(" id : %d \n name : %s %s\n age : %d \n Roll.No : %s \n school : %s \n\n", id, first_name, last_name, age, roll_no, school)
		}

	}

}

func main() {
	fmt.Println("Hi")

	var ch_table int

	connection_str := "postgres://postgres:test@localhost:5433/ajithmenon?sslmode=disable"
	db, err := sql.Open("postgres", connection_str)
	if err != nil {
		fmt.Printf(" Database Connection Error %v \n", err)
	}
	if err := db.Ping(); err != nil {
		fmt.Println("Ping Error ", err)
	}

	fmt.Println("SELCCT Table for operations ")
	fmt.Println(" 1. SCHOOL \n 2. STUDENT ")
	fmt.Scanln(&ch_table)

	switch ch_table {
	case 1:
		tableSchool(db)

	case 2:
		tableStudent(db)
	}

}
