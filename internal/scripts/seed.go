package main

import (
	"context"
	"fmt"
	"log"

	"github.com/MaheshMoholkar/attendance_tracking_backend/internal/database"
	"github.com/MaheshMoholkar/attendance_tracking_backend/internal/database/postgres"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	if err := SeedDatabase(); err != nil {
		log.Fatalf("Error seeding database: %v", err)
	}
}

func SeedDatabase() error {
	// Load environment variables
	if err := godotenv.Load(".env"); err != nil {
		return fmt.Errorf("error loading .env file: %w", err)
	}

	// Open database connection
	dbQueries, err := database.OpenDB()
	if err != nil {
		return fmt.Errorf("error opening database connection: %w", err)
	}
	store := database.New(dbQueries)

	// Seed classes
	if err := SeedClasses(store); err != nil {
		return fmt.Errorf("error seeding classes: %w", err)
	}

	// Seed divisions
	if err := SeedDivisions(store); err != nil {
		return fmt.Errorf("error seeding divisions: %w", err)
	}

	// Seed students
	if err := SeedStudents(store); err != nil {
		return fmt.Errorf("error seeding students: %w", err)
	}

	// Seed staff
	if err := SeedStaff(store); err != nil {
		return fmt.Errorf("error seeding staff: %w", err)
	}

	fmt.Println("Database seeded successfully!")
	return nil
}

func SeedClasses(store *database.Store) error {
	classNames := []string{"mba", "mca"}
	for _, className := range classNames {
		if _, err := store.DB.CreateClassInfo(context.Background(), className); err != nil {
			return fmt.Errorf("error creating class %s: %w", className, err)
		}
		fmt.Println("Created class:", className)
	}
	return nil
}

func SeedDivisions(store *database.Store) error {
	divisions := []struct {
		Name    string
		ClassID int32
	}{
		{"a", 1},
		{"b", 1},
		{"c", 1},
		{"a", 2},
		{"b", 2},
	}
	for _, division := range divisions {
		if _, err := store.DB.CreateDivisionInfo(context.Background(), postgres.CreateDivisionInfoParams{
			Divisionname: division.Name,
			ClassID:      division.ClassID,
		}); err != nil {
			return fmt.Errorf("error creating division %s for class ID %d: %w", division.Name, division.ClassID, err)
		}
		fmt.Printf("Created division: %s for class ID: %d\n", division.Name, division.ClassID)
	}
	return nil
}

func SeedStudents(store *database.Store) error {
	students := []struct {
		FirstName string
		LastName  string
		RollNo    int32
		Email     string
		ClassName string
		Division  string
		Year      int32
		StudentID int32
	}{
		{"Shubham", "Gaikwad", 360, "shubham@gmail.com", "mca", "a", 2023, 52453},
		{"Mahesh", "Moholkar", 123, "mahesh@gmail.com", "mca", "b", 2023, 12345},
		{"Tejas", "Asawale", 456, "tejas@gmail.com", "mca", "a", 2023, 67890},
		{"Vaibhav", "Panchal", 789, "vaibhav@gmail.com", "mca", "b", 2023, 98765},
	}
	for _, student := range students {
		if _, err := store.DB.CreateStudentCredentials(context.Background(), postgres.CreateStudentCredentialsParams{
			StudentID:    student.StudentID,
			PasswordHash: generateHashedPassword(student.StudentID),
		}); err != nil {
			return fmt.Errorf("error creating student credentials: %w", err)
		}

		if _, err := store.DB.CreateStudentInfo(context.Background(), postgres.CreateStudentInfoParams{
			Firstname: student.FirstName,
			Lastname:  student.LastName,
			Rollno:    student.RollNo,
			Email:     student.Email,
			Classname: student.ClassName,
			Division:  student.Division,
			Year:      student.Year,
			StudentID: student.StudentID,
		}); err != nil {
			return fmt.Errorf("error creating student: %w", err)
		}
		fmt.Printf("Created student: %s %s\n", student.FirstName, student.LastName)
	}
	return nil
}

func SeedStaff(store *database.Store) error {
	// Seed staff
	staff := []struct {
		FirstName string
		LastName  string
		Email     string
		StaffID   int32
		Password  string
	}{
		{"Shikha", "Dubey", "shikha@gmail.com", 1, "1"},
		{"Swati", "Jadhav", "swati@gmail.com", 2, "2"},
	}
	for _, s := range staff {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(s.Password), bcrypt.DefaultCost)
		if err != nil {
			return fmt.Errorf("error hashing password: %w", err)
		}

		if _, err := store.DB.CreateStaffCredentials(context.Background(), postgres.CreateStaffCredentialsParams{
			StaffID:      s.StaffID,
			PasswordHash: string(hashedPassword),
		}); err != nil {
			return fmt.Errorf("error creating staff credentials: %w", err)
		}

		if _, err := store.DB.CreateStaffInfo(context.Background(), postgres.CreateStaffInfoParams{
			Firstname: s.FirstName,
			Lastname:  s.LastName,
			Email:     s.Email,
			StaffID:   s.StaffID,
		}); err != nil {
			return fmt.Errorf("error creating staff: %w", err)
		}

		fmt.Printf("Created staff: %s %s\n", s.FirstName, s.LastName)
	}
	return nil
}

func generateHashedPassword(id int32) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(fmt.Sprintf("%d", id)), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Error generating hashed password: %v", err)
	}
	return string(hashedPassword)
}
