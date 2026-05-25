package api

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func CreateUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user User

		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			sendJSON(w, Response{Error: "invalid body"}, http.StatusBadRequest)
			return
		}

		if len(user.FirstName) < 2 ||
			len(user.FirstName) > 20 ||
			len(user.LastName) < 2 ||
			len(user.LastName) > 20 ||
			len(user.Biography) < 20 ||
			len(user.Biography) > 450 {

			sendJSON(
				w,
				Response{Error: "invalid fields"},
				http.StatusBadRequest,
			)
			return
		}

		user.ID = uuid.NewString()

		if err := CreateUserDB(db, user); err != nil {
			sendJSON(w, Response{Error: "failed to create user"}, http.StatusInternalServerError)
			return
		}

		sendJSON(w, Response{Data: user}, http.StatusCreated)
	}
}

func FindAll(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := ListUsersDB(db)
		if err != nil {
			sendJSON(w, Response{Error: "failed to fetch users"}, http.StatusInternalServerError)
			return
		}

		if users == nil {
			users = make([]User, 0)
		}

		sendJSON(w, Response{Data: users}, http.StatusOK)
	}
}

func FindById(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		user, err := GetUserDB(db, id)
		if err != nil {
			sendJSON(w, Response{Error: "user not found"}, http.StatusNotFound)
			return
		}

		sendJSON(w, Response{Data: user}, http.StatusOK)
	}
}

func Update(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		_, err := GetUserDB(db, id)
		if err != nil {
			sendJSON(w, Response{Error: "user not found"}, http.StatusNotFound)
			return
		}

		var user User

		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			sendJSON(w, Response{Error: "invalid body"}, http.StatusBadRequest)
			return
		}

		if len(user.FirstName) < 2 ||
			len(user.FirstName) > 20 ||
			len(user.LastName) < 2 ||
			len(user.LastName) > 20 ||
			len(user.Biography) < 20 ||
			len(user.Biography) > 450 {

			sendJSON(
				w,
				Response{Error: "invalid fields"},
				http.StatusBadRequest,
			)
			return
		}

		user.ID = id

		if err := UpdateUserDB(db, user); err != nil {
			sendJSON(w, Response{Error: "failed to update user"}, http.StatusInternalServerError)
			return
		}

		sendJSON(w, Response{Data: user}, http.StatusOK)
	}
}

func Delete(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		user, err := DeleteUserDB(db, id)
		if err != nil {
			sendJSON(w, Response{Error: "user not found"}, http.StatusNotFound)
			return
		}

		sendJSON(w, Response{Data: user}, http.StatusOK)
	}
}
