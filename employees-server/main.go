package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Employee struct to map to the employees table
type Employee struct {
	ID           uint   `json:"id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Email        string `json:"email"`
	HireDate     string `json:"hire_date"`
	Salary       string `json:"salary"`
	DepartmentID int    `json:"department_id"`
}

// Initialize the database connection
func initDB() *gorm.DB {
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	return db
}

var db *gorm.DB

// Create a new employee
func createEmployee(w http.ResponseWriter, r *http.Request) {
	var employee Employee
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&employee); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := db.Create(&employee).Error; err != nil {
		http.Error(w, "Could not create employee", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(employee)
}

// Get all employees
func getEmployees(w http.ResponseWriter, r *http.Request) {
	var employees []Employee
	if err := db.Find(&employees).Error; err != nil {
		http.Error(w, "Could not retrieve employees", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(employees)
}

// Get employee by ID
func getEmployee(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/employees/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var employee Employee
	if err := db.First(&employee, id).Error; err != nil {
		http.Error(w, "Employee not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(employee)
}

func main() {
	// Initialize the database connection
	// var err error
	db = initDB()
	defer db.Close()

	// Automatically create the employees table if it doesn't exist
	db.AutoMigrate(&Employee{})

	// Set up the standard HTTP router
	http.HandleFunc("/employees", getEmployees)
	http.HandleFunc("/employees/", getEmployee)
	http.HandleFunc("/employees/create", createEmployee)

	// Start the server
	fmt.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
