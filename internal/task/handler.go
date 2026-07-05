package task

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
)

type Handler struct {
	store *Store
}

func NewHandler(store *Store) *Handler {
	return &Handler{store: store}
}

func (h *Handler) Routes(mux *http.ServeMux) {
	mux.HandleFunc("GET /tasks", h.list)
	mux.HandleFunc("POST /tasks", h.create)
	mux.HandleFunc("GET /tasks/{id}", h.get)
	mux.HandleFunc("PUT /tasks/{id}", h.update)
	mux.HandleFunc("DELETE /tasks/{id}", h.delete)
}

func (h *Handler) list(w http.ResponseWriter, r *http.Request) {
	tasks, err := h.store.List(r.Context())
	if err != nil {
		writeStoreError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, tasks)
}

func (h *Handler) create(w http.ResponseWriter, r *http.Request) {
	var t Task
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := t.Validate(); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	created, err := h.store.Create(r.Context(), t)
	if err != nil {
		writeStoreError(w, err)
		return
	}

	writeJSON(w, http.StatusCreated, created)
}

func (h *Handler) get(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid task id")
		return
	}

	t, err := h.store.Get(r.Context(), id)
	if err != nil {
		writeStoreError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, t)
}

// update is a full replace (HTTP PUT): the whole task is overwritten from the
// request body, except the server-owned ID and creation time. A partial update
// would use PATCH, which we do not implement here.
func (h *Handler) update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid task id")
		return
	}

	var t Task
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := t.Validate(); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	updated, err := h.store.Update(r.Context(), id, t)
	if err != nil {
		writeStoreError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, updated)
}

func (h *Handler) delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid task id")
		return
	}

	if err := h.store.Delete(r.Context(), id); err != nil {
		writeStoreError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		log.Printf("writeJSON: %v", err)
	}
}

// errorResponse is the single shape every error takes, so a client can rely on
// one JSON envelope: {"error":"<message>"}.
type errorResponse struct {
	Error string `json:"error"`
}

// writeError writes the JSON error envelope with the given status code. Every
// error path in this package goes through it, so all errors look the same.
func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, errorResponse{Error: message})
}

// writeStoreError maps a store error to the right status in one place:
// ErrNotFound is a 404 the client caused; anything else is an unexpected 500
// whose real cause we log server-side but never leak to the client.
func writeStoreError(w http.ResponseWriter, err error) {
	if errors.Is(err, ErrNotFound) {
		writeError(w, http.StatusNotFound, "task not found")
		return
	}
	log.Printf("unexpected store error: %v", err)
	writeError(w, http.StatusInternalServerError, "something went wrong")
}
