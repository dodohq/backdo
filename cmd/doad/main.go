package main

import (
	"database/sql"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"golang.org/x/crypto/bcrypt"

	_ "github.com/lib/pq"
	"golang.org/x/crypto/ssh/terminal"
)

func main() {
	actionPtr := flag.String("action", "", "[Reqruired] 'create' or 'changepwd'")
	dbPtr := flag.String("db", "", "[Required] database connection string")
	flag.Parse()

	if *actionPtr == "" || *dbPtr == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	dbConn, err := sql.Open("postgres", *dbPtr)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer dbConn.Close()

	if strings.ToLower(*actionPtr) == "create" {
		email := prompForEmail()
		hashBytes, err := prompForPassword()
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		hashedPwd := string(hashBytes)
		query := `INSERT INTO admins(email, password) VALUES ($1, $2) RETURNING id`
		var ID int64
		err = dbConn.QueryRow(query, email, hashedPwd).Scan(&ID)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		fmt.Printf("\nAdmin with id %d created\n", ID)
	} else if strings.ToLower(*actionPtr) == "changepwd" {
		email := prompForEmail()
		query := `SELECT id FROM admins WHERE email = $1`
		var ID int64
		rows, err := dbConn.Query(query, email)
		if rows.Next() {
			err = rows.Scan(&ID)
			if err != nil {
				log.Fatal(err)
				os.Exit(1)
			}
		}
		if ID < 1 {
			log.Fatal(errors.New("Email Not Found"))
			os.Exit(1)
		}

		hashBytes, err := prompForPassword()
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		query = `UPDATE admins SET password = $1 WHERE id = $2 RETURNING id`
		err = dbConn.QueryRow(query, hashBytes, ID).Scan(&ID)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		fmt.Printf("\nPassword for admin %d has been changed\n", ID)
	}
}

func prompForEmail() string {
	fmt.Print("Enter Email: ")
	var email string
	fmt.Scan(&email)
	return email
}

func prompForPassword() ([]byte, error) {
	fmt.Print("Enter Password: ")
	passwordBytes, err := terminal.ReadPassword(0)
	if err != nil {
		return nil, err
	}
	for len(string(passwordBytes)) < 10 {
		fmt.Println("Password too short")
		fmt.Print("Enter Password: ")
		passwordBytes, err = terminal.ReadPassword(0)
		if err != nil {
			return nil, err
		}
	}
	fmt.Print("\nConfirm Password: ")
	reenterBytes, err := terminal.ReadPassword(0)
	if err != nil {
		return nil, err
	}
	for string(reenterBytes) != string(passwordBytes) {
		fmt.Println("Wrong Confirmation")
		fmt.Print("Confirm Password: ")
		reenterBytes, err = terminal.ReadPassword(0)
		if err != nil {
			return nil, err
		}
	}

	hashBytes, err := bcrypt.GenerateFromPassword(passwordBytes, 14)
	if err != nil {
		return nil, err
	}

	return hashBytes, nil
}
