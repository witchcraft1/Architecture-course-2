package dormitories

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/witchcraft1/Architecture-course-2/server/tools"
)

type HttpHandlerFunc http.HandlerFunc

func HttpHandler(store *Store) HttpHandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			handleGetDormitory(r, rw, store)
		} else if r.Method == "POST"{
			handleStudentAdd(r, rw, store)
		} else {
			rw.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}

func handleStudentAdd(r *http.Request, rw http.ResponseWriter, store *Store) {
	var c Student
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		log.Printf("Error decoding dormitory input: %s", err)
		tools.WriteJsonBadRequest(rw, "bad JSON payload")
		return
	}

	err := store.AddStudent(c.Name, c.DromitoryId, c.Specialty)
	if err == nil {
		tools.WriteJsonOk(rw, &c)
	} else {
		log.Printf("Error inserting record: %s", err)
		tools.WriteJsonInternalError(rw)
	}
}

func handleGetDormitory(r *http.Request, rw http.ResponseWriter, store *Store) {
	var c Student
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		log.Printf("Error decoding dormitory input: %s", err)
		tools.WriteJsonBadRequest(rw, "bad JSON payload")
		return
	}
	res, err := store.GetDormitory(c.Specialty)
	if err != nil {
		log.Printf("Error making query to the db: %s", err)
		tools.WriteJsonInternalError(rw)
		return
	}
	tools.WriteJsonOk(rw, res)
}
