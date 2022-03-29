package main

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "root"
	dbname   = "goLangDB"
)

func main() {
	// ********** Declaration.... ********** //
	psqlconnect := fmt.Sprintf("host = %s port = %d user = %s password = %s dbname = %s sslmode=disable", host, port, user, password, dbname)
	userchoice := 1
	// ********** Processing..... ********** //
	db, err := sql.Open("postgres", psqlconnect)
	// check error.
	CheckError(err)
	defer db.Close()

	// Starting our Januery
	for userchoice != 6 {
		menuServices()
		userchoice = TakeUserIn()
		if userchoice != 6 {
			manageUserChoice(userchoice, db)
		}
	}

	println("*** Good Bye.")

}

// Check Erro function
func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

// show our Services
func menuServices() {
	fmt.Println("\n\n\n************ Welcome to the go DB task1 home page ***********")
	fmt.Println("\nPlease Select number of Service you want.\n")
	fmt.Println("\t1: Get All Employees")
	fmt.Println("\t2: Get a Single Employee")
	fmt.Println("\t3: Add a Employee")
	fmt.Println("\t4: Update a Employee")
	fmt.Println("\t5: Delete a Employee")
	fmt.Println("\t6: Exit\n")
	fmt.Print("Your Input : ")
}

// Take Input from user...
func TakeUserIn() int {
	var userInput string
	fmt.Scan(&userInput)
	Result, err := strconv.Atoi(userInput)
	if err != nil {
		panic(err)
	}
	return Result
}

// Manage User Choice
func manageUserChoice(choice int, db *sql.DB) {
	switch choice {
	case 1:
		selectAll(db)
	case 2:
		getSingleEmployee(db)
	case 3:
		AddEmployee(db)
	case 4:
		EditEmployee(db)
	case 5:
		DeleteEmployee(db)
	default:
		main()
	}
}

// Sellect all
func selectAll(db *sql.DB) {
	rows, err := db.Query(`select * from "Employee"`)
	if err != nil {
		fmt.Println("Database error")
	}
	fmt.Println("Id \t Name")
	for rows.Next() {
		var id, Name string
		err := rows.Scan(&id, &Name)
		if err != nil {
			fmt.Println("Database error")
		}
		fmt.Println(id+"\t", Name)
	}
}

// Select single employee.
func getSingleEmployee(db *sql.DB) {
	fmt.Print("please enter the Employee id : ")
	id := TakeUserIn()
	result, err := db.Query(`select * from "Employee" where id  = $1`, id)
	CheckError(err)

	for result.Next() {
		var id, Name string
		err := result.Scan(&id, &Name)
		CheckError(err)
		fmt.Println("Id \t Name")
		fmt.Println(id+"\t", Name)
	}
	result.Close()
}

// Add new Employee
func AddEmployee(db *sql.DB) {
	var emp_Name string
	fmt.Print("Employee National Id :  ")
	emp_id := TakeUserIn()
	fmt.Print("Employee Name :  ")
	fmt.Scan(&emp_Name)
	// dynamic insert to database
	dynamicInsert := `insert into "Employee"("id","Name") values($1,$2)`
	// Execution...
	_, err := db.Exec(dynamicInsert, emp_id, emp_Name)
	// check error.
	CheckError(err)
	fmt.Println("Process done Successfully")
}

// Edit new Employee
func EditEmployee(db *sql.DB) {
	var emp_Name string
	fmt.Print("Employee National Id of Employee:  ")
	emp_id := TakeUserIn()
	fmt.Print("Employee New Name :  ")
	fmt.Scan(&emp_Name)
	// dynamic insert to database
	dynamicInsert := `UPDATE "Employee" SET "Name" = $2 WHERE "id" = $1`
	// Execution...
	_, err := db.Exec(dynamicInsert, emp_id, emp_Name)
	// check error.
	CheckError(err)
	fmt.Println("Process done Successfully")
}

// Edit new Employee
func DeleteEmployee(db *sql.DB) {

	fmt.Print("Employee National Id of Employee you want to delete :  ")
	emp_id := TakeUserIn()
	// dynamic insert to database
	dynamicInsert := `delete from "Employee" WHERE "id" = $1`
	// Execution...
	_, err := db.Exec(dynamicInsert, emp_id)
	// check error.
	CheckError(err)
	fmt.Println("Process done Successfully")
}
