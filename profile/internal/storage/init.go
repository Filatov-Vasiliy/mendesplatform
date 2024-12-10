package storage

import "fmt"

const insertStatus = `INSERT INTO status VALUES('ACTIVE','Active users',current_timestamp,current_timestamp);
	 INSERT INTO status VALUES('INACTIVE','Active users',current_timestamp,current_timestamp);
	 INSERT INTO status VALUES('BLOCKED','Active users',current_timestamp,current_timestamp);
	 INSERT INTO status VALUES('DELETED','Active users',current_timestamp,current_timestamp);
`
const createTableProfile = `CREATE TABLE IF NOT EXISTS profiles (
    id uuid PRIMARY KEY,
    login VARCHAR(255) NOT NULL,
    firstname VARCHAR(255) NULL,
    surname VARCHAR(255) NULL,
    patronymic VARCHAR(255) NULL,
    password VARCHAR(255) NOT NULL,
    email VARCHAR(255) NULL,
    email_confirmed VARCHAR(255) NULL,
    status VARCHAR(255) NOT NULL,
    number VARCHAR(255) NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);`
const createTableStatus = `CREATE TABLE IF NOT EXISTS status (
	id uuid PRIMARY KEY,
	status VARCHAR(255) NOT NULL,
	description VARCHAR(255) NOT NULL,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL,
);`

func createTables(c *Storage) error {
	_, err := c.db.Exec(createTableProfile)
	if err != nil {
		fmt.Println("Profiles not init")
		return err
	}
	fmt.Println("Profiles init")
	_, err = c.db.Exec(createTableStatus)
	if err != nil {
		fmt.Println("Status not init")
		return err
	}
	fmt.Println("Status init")
	return nil
}

func insertStatusDB(c *Storage) error {
	data, err := c.db.Query(`SELECT * FROM status`)
	if err != nil {
		return err
	}
	if data == nil {
		_, err = c.db.Exec(insertStatus)
		if err != nil {
			fmt.Println("Statuses inserted")
			return err
		}
	}
	fmt.Println("Statuses existed")
	return nil
}
