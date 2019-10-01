package main

import (
	"errors"
	"time"
	"database/sql"
	"github.com/lib/pq"
)

// Juegos estructura
type Juegos struct {
	ID        int
	Name      string
	Age       int64
	Active    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

//Crear registro de un juego en BD

func Crear(e Juegos) error {
	q := "INSERT INTO juegos (name, age, active) VALUES ($1, $2, $3)"

	intNull := sql.NullInt64{}
	strNull := sql.NullString{}

	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(q)
	if err != nil {
		return err
	}
	defer stmt.Close()

	if e.Age == 0 {
		intNull.Valid = false
	}else{
		intNull.Valid = true
		intNull.Int64 = int64(e.Age)
	}
	if e.Name == "" {
		strNull.Valid = false
	}else{
		strNull.Valid = true 
		strNull.String = e.Name
	}

	r, err := stmt.Exec(strNull, intNull, e.Active)
	if err != nil {
		return err
	}
	i, _ := r.RowsAffected()
	if i != 1 {
		return errors.New("Error: Se esperaba una fila afectada.")
	}

	return nil
}

//Buscar o consulta informacion juegos (todos)


func Consultar()(juegos []Juegos, err error){
	q := "SELECT id, name, age, active, created_at, update_at FROM juegos"

	timeNull := pq.NullTime{}
	intNull := sql.NullInt64{}
	strNull := sql.NullString{}
	boolNull := sql.NullBool{}


	db := getConnection()
	defer db.Close()

	rows, err := db.Query(q)
	if err != nil {
		return 
	}
	defer rows.Close()

	for rows.Next() {
		j := Juegos{}
		err = rows.Scan(&j.ID, &strNull, &intNull, &boolNull, &j.CreatedAt, &timeNull)
		if err != nil {
			return
		}

		j.UpdatedAt = timeNull.Time
		j.Name = strNull.String
		j.Age = intNull.Int64
		j.Active = boolNull.Bool

		juegos = append(juegos, j)
	}
	return juegos, nil 

}

//Actualizar datos en la BD

func Actualizar(e Juegos) error {
	q := "UPDATE juegos SET name = $1, age = $2, active = $3, update_at = now() WHERE id = $4"

	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(q)
	if err != nil {
		return err
	}
	defer stmt.Close()
	r, err := stmt.Exec(e.Name, e.Age, e.Active, e.ID)
	if err != nil {
		return err
	}

	i, _ := r.RowsAffected()
	if i != 1 {
		return errors.New("Error: Se esperaba una fila afectada.")
	}

	return nil

}

func Borrar (ID int)error {
	q := "DELETE FROM juegos WHERE id = $1"

	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(q)
	if err != nil {
		return err
	}
	defer stmt.Close()

	r, err := stmt.Exec(ID)

	if err != nil {
		return err
	}

	i, _ := r.RowsAffected()
	if i != 1 {
		return errors.New("Error: Se esperaba una fila afectada.")
	}

	return nil
}