package api

import (
	"encoding/json"
	"fmt"
	"io"
	"slices"
	"net/http"
	openapi_types "github.com/oapi-codegen/runtime/types"

	//"artifacts/internal/storage"
	"artifacts/internal/storage/storage_error"
)

func (s Server) GetCharts(w http.ResponseWriter, r *http.Request) {
	data, err := s.storageHandler.Read("", "", "")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(data)
}

func (s Server) GetChart(w http.ResponseWriter, r *http.Request, repository string, resource string, version string) {
	w.Header().Set("Content-Type", "application/json")

	data, err := s.storageHandler.Read(repository, resource, version)

	if err != nil {
		if err == storage_error.NotFound {
			w.WriteHeader(http.StatusNotFound)
			_ = json.NewEncoder(w).Encode(Error{
				Code: new(int(http.StatusNotFound)),
				Message: new(string(fmt.Sprintf("Resource '%v:%v' was not found", resource, version))),
			})
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(Error{
			Code: new(int(http.StatusInternalServerError)),
			Message: new(string(fmt.Sprintf("Internal server error occurred for '%v:%v'", resource, version))),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(data)
}

func (Server) GetChartVersions(w http.ResponseWriter, r *http.Request, name string) {
	// TODO: interface to storage backend for querying
	res := fmt.Sprintf("%v", r)
	_ = json.NewEncoder(w).Encode(res)
}

func (s Server) AddChart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")  // TODO: set globally, or default?

	// Set max header size (50MB)
	r.Body = http.MaxBytesReader(w, r.Body, 50<<20) // TODO: set globally, or default?

	// Limit parsed form size (32MB)
	if err := r.ParseMultipartForm(32 << 20); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(Error{
			Code: new(int(http.StatusBadRequest)),
			Message: new(string("Invalid form data")),
		})
		return
	}
	defer func() {
		if r.MultipartForm != nil {
			// Clean up temp files
			r.MultipartForm.RemoveAll()
		}
	}()

	repo := r.FormValue("repository")
	if repo == "" {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(Error{
			Code: new(int(http.StatusBadRequest)),
			Message: new(string("Missing 'repository' parameter")),
		})
		return
	}

	// Extract chart name from form data
	name := r.FormValue("name")
	if name == "" {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(Error{
			Code: new(int(http.StatusBadRequest)),
			Message: new(string("Missing 'name' parameter")),
		})
		return
	}

	// Extract chart bytes from form data
	f, fh, err := r.FormFile("chart")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(Error{
			Code: new(int(http.StatusBadRequest)),
			Message: new(string("Missing chart")),
		})
		return
	}
	defer f.Close()

	// Reject unsupported MIME type
	if eq := slices.Compare(fh.Header["Content-Type"], chartMIMEType); eq != 0 {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		_ = json.NewEncoder(w).Encode(Error{
			Code: new(int(http.StatusUnsupportedMediaType)),
			Message: new(string(fmt.Sprintf("Unsupported type: %v", fh.Header["Content-Type"]))),
		})
		return
	}

	data, err := io.ReadAll(f)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(Error{
			Code: new(int(http.StatusBadRequest)),
			Message: new(string("Malformed file contents")),
		})
	}

	// Construct runtime type for compatibility with generated types
	file := &openapi_types.File{}
	file.InitFromBytes(data, fh.Filename)

	bytes, err := file.Bytes()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(Error{
			Code: new(int(http.StatusBadRequest)),
			Message: new(string("Malformed file upload")),
		})
	}

	if err := s.storageHandler.Write(repo, name, bytes); err != nil {  // TODO: return created version
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(Error{
			Code: new(int(http.StatusInternalServerError)),
			Message: new(string("Error occurred during file IO")),
		})
		return
	}

	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(Error{
		Code: new(int(http.StatusCreated)),
		Message: new(string(fmt.Sprintf("Chart '%v' created", name))),
	})
}
