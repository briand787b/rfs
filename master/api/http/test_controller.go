package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/briand787b/rfs/app/models"
)

func init() {
	http.HandleFunc("/", handleTestModel)
}

func handleTestModel(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(w, "started working")
	var testModel models.TestModel

	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&testModel); err != nil {
		fmt.Println(err)
		http.Error(w, "malformed test_model", http.StatusUnprocessableEntity)
		return
	}

	if err := models.NewPostgresTestModelDBStore().Save(&testModel); err != nil {
		fmt.Println(err)
		http.Error(w, "could not save test_model", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(testModel); err != nil {
		fmt.Println(err)
		http.Error(w, "could not encode test_model", http.StatusInternalServerError)
		return
	}

	// fmt.Fprintln(w, "done working")
}