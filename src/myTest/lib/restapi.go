package lib

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Todo struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

var todos []Todo

func StartApi() {
	todos = append(todos, Todo{"1", "test", false})
	http.HandleFunc("/todos", getTodos)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getTodos(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("variable r.Method = %v is of type %T \n", r.Method, r.Method)

	switch {
	case r.Method == http.MethodPost:

		var t Todo

		err := json.NewDecoder(r.Body).Decode(&t)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		fmt.Printf("Todo: %+v \n\n", t)
		todos = append(todos, t)
		return
	case r.Method == http.MethodGet:
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(todos)
		return
	case r.Method == http.MethodPut:
		return
	case r.Method == http.MethodDelete:
		return
	default:
		return
	}

}

func AddTodo(todo Todo) {
	todos = append(todos, todo)
}
