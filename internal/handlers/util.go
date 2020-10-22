package handlers

import (
	"encoding/json"
	"github.com/apex/log"
	"github.com/gorilla/schema"
	"io"
	"net/http"
)

const (
	// Success represents a success response in the predefined responses map
	Success = iota
	// NoContent represents a no content error response in the predefined responses map
	NoContent = iota
	// Redirect represents a redirect response in the predefined responses map
	Redirect = iota
	// Bad represents bad request response in the predefined responses map
	Bad = iota
	// Unauthorized  represents a unauthorized request response in the predefined responses map
	Unauthorized = iota
	// Forbidden represents a forbidden request response in the predefined responses map
	Forbidden = iota
	// NotAllowed represents a not allowed request response in the predefined responses map
	NotAllowed = iota
	// NotFound represents a not found request response in the predefined responses map
	NotFound = iota
	// InternalError represents a internal server error response in the predefined responses map
	InternalError = iota
)

//Endpoints struct that holds methods and parameters useful for our endpoints and the endpoint methods themselves
type Endpoints struct {
	responses map[int]response
}

//New creates the Endpoints struct
func New() (e Endpoints) {
	e.load()
	return
}

type response struct {
	Code   int
	String string
}

//Body returns the []byte of a predefined response
func (r response) Body() []byte {
	return []byte(r.String)
}

func (e *Endpoints) load() {
	e.responses = map[int]response{
		Success:       {200, `{"code": 200, "message": "Asked and done"}`},
		NoContent:     {204, `{"code": 204, "message": "No content!"}`},
		Redirect:      {302, `{"code": 302, "message": "Never gonna give you up..."}`},
		Bad:           {400, `{"code": 400, "message": "Are you talking to me?"}`},
		Unauthorized:  {401, `{"code": 401, "message": "Who are you?"}`},
		Forbidden:     {403, `{"code": 403, "message": "Can't touch this!"}`},
		NotAllowed:    {405, `{"code": 405, "message": "This is not the way"}`},
		NotFound:      {404, `{"code": 404, "message": "You do not know the way"}`},
		InternalError: {500, `{"code": 500, "message": "Get to the chopper!"}`},
	}
}

// DecodeJSON decodes a JSON
func (e Endpoints) DecodeJSON(body io.Reader, entity interface{}) (err error) {
	d := json.NewDecoder(body)
	err = d.Decode(entity)
	return
}

// DecodeQueryParameters decodes query string parameters
func (e Endpoints) DecodeQueryParameters(dst interface{}, src map[string][]string) (err error) {
	d := schema.NewDecoder()
	err = d.Decode(dst, src)
	return
}

//EncodeJSON encodes a JSON to []byte
func (e Endpoints) EncodeJSON(toMarshall interface{}, ident bool) (j []byte, err error) {
	if !ident {
		j, err = json.Marshal(toMarshall)
	} else {
		j, err = json.MarshalIndent(toMarshall, "", " ")
	}
	return
}

// Response writes a response to http.ResponseWriter
func (e Endpoints) Response(w http.ResponseWriter, resp int, body ...byte) {
	// get response from default responses
	r, ok := e.responses[resp]
	if !ok {
		r.Code = 500
		r.String = `{"code": 500, "message": "Response Not Implemented!"}`
	}
	w.WriteHeader(r.Code)
	if body == nil {
		body = r.Body()
	}
	_, err := w.Write(body)
	if err != nil {
		log.Errorf("failed writing body", err)
	}
	return
}

// Response performs a http.Redirect to the supplied url
func (e Endpoints) Redirect(w http.ResponseWriter, redirectURL string) {
	w.Header().Set("Location", redirectURL)
	w.WriteHeader(e.responses[Redirect].Code)
	_, err := w.Write(e.responses[Redirect].Body())
	if err != nil {
		log.Errorf("failed writing body", err)
	}
	return
}
