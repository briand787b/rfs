package http

import (
	"fmt"
	"net/http"
)

func handleTestModel(w http.ResponseWriter, r *http.Request) {
	// var testModel models.TestModel

	// defer r.Body.Close()
	// if err := json.NewDecoder(r.Body).Decode(&testModel); err != nil {
	// 	fmt.Println(err)
	// 	http.Error(w, "malformed test_model", http.StatusUnprocessableEntity)
	// 	return
	// }

	// if err := models.NewPostgresTestModelDBStore().Save(&testModel); err != nil {
	// 	fmt.Println(err)
	// 	http.Error(w, "could not save test_model", http.StatusInternalServerError)
	// 	return
	// }

	// if err := json.NewEncoder(w).Encode(testModel); err != nil {
	// 	fmt.Println(err)
	// 	http.Error(w, "could not encode test_model", http.StatusInternalServerError)
	// 	return
	// }

	fmt.Fprintln(w, "NOT IMPLEMENTED")
}

// handleTestModelGetAll retrieves all test models from the database
func handleTestModelGetAll(w http.ResponseWriter, r *http.Request) {
	// ts, err := models.NewPostgresTestModelDBStore().GetAll()
	// if err != nil {
	// 	fmt.Println(err)
	// 	http.Error(w, "could not get all test_models", http.StatusInternalServerError)
	// 	return
	// }

	// if err := render.RenderList(w, r, NewTestModelListResponse(ts)); err != nil {
	// 	fmt.Println(err)
	// 	http.Error(w, "could not encode test_models", http.StatusInternalServerError)
	// 	return
	// }

	fmt.Fprintln(w, "NOT IMPLEMENTED")
}
