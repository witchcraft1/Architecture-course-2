package dormitories

import (
	"database/sql"
	"fmt"
)

type ScenaryDormitory struct {
	Id            int            `json:"id"`
	Name          string         `json:"name"`
	StudentsCount map[string]int `json:"studentsCount"`
}

type Student struct {
	Name        string `json:"name"`
	DromitoryId int    `json:"dormitoryId"`
	Specialty   string `json:"specialty"`
}

type Store struct {
	Db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{Db: db}
}

func (s *Store) GetDormitory(specialty string) (*ScenaryDormitory, error) {
	if len(specialty) < 0 {
		return nil, fmt.Errorf("Specialty is not provided")
	}
	rows, err := s.Db.Query("SELECT dormitory.id, COUNT(student.*) FROM dormitory LEFT JOIN Speciality ON true LEFT JOIN student ON student.dormitoryId = dormitory.id and student.specialityId = Speciality.id WHERE speciality.name = $1 GROUP BY dormitory.id, Speciality.id ORDER BY dormitory.id;", specialty)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	id, count := 0, 0
	b := true
	for rows.Next() {
		var c [2]int
		if err := rows.Scan(&c[0], &c[1]); err != nil {
			return nil, err
		}
		if b {
			count = c[1]
			id = c[0]
			b = false
		}
		if count > c[1] {
			count = c[1]
			id = c[0]
		}
	}
	rows, err = s.Db.Query("SELECT dormitory.id, dormitory.name, Speciality.name, COUNT(student.*) FROM dormitory LEFT JOIN Speciality ON true LEFT JOIN student ON student.dormitoryId = dormitory.id and student.specialityId = Speciality.id WHERE dormitory.id = $1 GROUP BY dormitory.id, Speciality.id;", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res ScenaryDormitory
	m := make(map[string]int)
	for rows.Next() {
		var spec string
		var count int
		if err := rows.Scan(&res.Id, &res.Name, &spec, &count); err != nil {
			return nil, err
		}
		m[spec] = count
	}
	res.StudentsCount = m
	return &res, nil
}

func (s *Store) AddStudent(name string, dormitoryId int, specialty string) error {
	if len(name) < 0 || len(specialty) < 0 {
		return fmt.Errorf("Student name is not provided")
	}
	var specialityId int
	err := s.Db.QueryRow("SELECT speciality.id FROM speciality WHERE speciality.name = $1;", specialty).Scan(&specialityId)
	switch {
	case err == sql.ErrNoRows:
		return fmt.Errorf("Speciality is not provided")
	default:
		_, err = s.Db.Exec("INSERT INTO student (name,dormitoryId,specialityId)VALUES ($1, $2, $3)", name, dormitoryId, specialityId)
		return err
	}
}
