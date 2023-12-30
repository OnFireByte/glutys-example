// This file is generated by glutys. DO NOT EDIT.

package routegen

import (
	"encoding/json"
	glutys "github.com/onfirebyte/glutys"
	"net/http"
	cache "github.com/OnFireByte/glutys-example/server/di/cache"
	reqcontext "github.com/OnFireByte/glutys-example/server/reqcontext"
	math "github.com/OnFireByte/glutys-example/server/route/math"
	todolist "github.com/OnFireByte/glutys-example/server/route/todolist"
)

type Handler struct {
	Cache cache.Cache
}

func NewHandler(cache cache.Cache) *Handler {
	return &Handler{Cache: cache}
}
func (h *Handler) TodolistGetAllHandler(w http.ResponseWriter, r *http.Request, body *glutys.RequestBody) {
	Username0, errUsername0 := reqcontext.ParseUsername(r)
	if errUsername0 != nil {
		response := map[string]interface{}{
			"error": "Invalid Context",
			"msg":   errUsername0.Error(),
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}
	res := todolist.GetAll(Username0)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
	return
}
func (h *Handler) TodolistUpdateHandler(w http.ResponseWriter, r *http.Request, body *glutys.RequestBody) {
	Username0, errUsername0 := reqcontext.ParseUsername(r)
	if errUsername0 != nil {
		response := map[string]interface{}{
			"error": "Invalid Context",
			"msg":   errUsername0.Error(),
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}
	var todo1 *todolist.Todo
	errtodo1 := json.Unmarshal(body.Args[0], &todo1)
	if errtodo1 != nil {
		response := map[string]interface{}{
			"error": "Invalid JSON",
			"msg":   errtodo1.Error(),
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}
	res, err := todolist.Update(Username0, todo1)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": "Bad Request",
			"msg":   err.Error(),
		})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
	return
}
func (h *Handler) MathFibHandler(w http.ResponseWriter, r *http.Request, body *glutys.RequestBody) {
	var int1 int
	errint1 := json.Unmarshal(body.Args[0], &int1)
	if errint1 != nil {
		response := map[string]interface{}{
			"error": "Invalid JSON",
			"msg":   errint1.Error(),
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}
	res := math.Fib(h.Cache, int1)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
	return
}
func (h *Handler) TodolistAddHandler(w http.ResponseWriter, r *http.Request, body *glutys.RequestBody) {
	Username0, errUsername0 := reqcontext.ParseUsername(r)
	if errUsername0 != nil {
		response := map[string]interface{}{
			"error": "Invalid Context",
			"msg":   errUsername0.Error(),
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}
	var string1 string
	errstring1 := json.Unmarshal(body.Args[0], &string1)
	if errstring1 != nil {
		response := map[string]interface{}{
			"error": "Invalid JSON",
			"msg":   errstring1.Error(),
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}
	var string2 []string
	errstring2 := json.Unmarshal(body.Args[1], &string2)
	if errstring2 != nil {
		response := map[string]interface{}{
			"error": "Invalid JSON",
			"msg":   errstring2.Error(),
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}
	res := todolist.Add(Username0, string1, string2)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
	return
}
func (h *Handler) TodolistBulkAddHandler(w http.ResponseWriter, r *http.Request, body *glutys.RequestBody) {
	Username0, errUsername0 := reqcontext.ParseUsername(r)
	if errUsername0 != nil {
		response := map[string]interface{}{
			"error": "Invalid Context",
			"msg":   errUsername0.Error(),
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}
	var string1 []string
	errstring1 := json.Unmarshal(body.Args[0], &string1)
	if errstring1 != nil {
		response := map[string]interface{}{
			"error": "Invalid JSON",
			"msg":   errstring1.Error(),
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}
	res, err := todolist.BulkAdd(Username0, string1)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": "Bad Request",
			"msg":   err.Error(),
		})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
	return
}
func (h *Handler) TodolistGetHandler(w http.ResponseWriter, r *http.Request, body *glutys.RequestBody) {
	Username0, errUsername0 := reqcontext.ParseUsername(r)
	if errUsername0 != nil {
		response := map[string]interface{}{
			"error": "Invalid Context",
			"msg":   errUsername0.Error(),
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}
	var int1 int
	errint1 := json.Unmarshal(body.Args[0], &int1)
	if errint1 != nil {
		response := map[string]interface{}{
			"error": "Invalid JSON",
			"msg":   errint1.Error(),
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}
	res := todolist.Get(Username0, int1)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
	return
}

type HandlerFunc func(http.ResponseWriter, *http.Request, *glutys.RequestBody)

func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	body := glutys.RequestBody{}
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		response := map[string]interface{}{
			"error": "Bad Request",
			"msg":   "Invalid JSON",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
	}
	switch body.Method {
	case "math.fib":
		h.MathFibHandler(w, r, &body)
	case "todolist.add":
		h.TodolistAddHandler(w, r, &body)
	case "todolist.bulkAdd":
		h.TodolistBulkAddHandler(w, r, &body)
	case "todolist.get":
		h.TodolistGetHandler(w, r, &body)
	case "todolist.getAll":
		h.TodolistGetAllHandler(w, r, &body)
	case "todolist.update":
		h.TodolistUpdateHandler(w, r, &body)
	}
}
