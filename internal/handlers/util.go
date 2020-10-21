package handlers

import (
	"encoding/json"
	"github.com/apex/log"
	"github.com/gorilla/schema"
	"github.com/insan1k/one-qr-dot-me/internal/configuration"
	"io"
	"net/http"
)

const (
	Success       = "success"
	NoContent     = "no-content"
	Redirect      = "redirect"
	Bad           = "bad"
	Unauthorized  = "unauthorized"
	Forbidden     = "forbidden"
	NotAllowed    = "not-allowed"
	NotFound      = "not-found"
	InternalError = "server"
)

type Endpoints struct {
	responses map[string]response
	c         configuration.Configuration
}

func New()(e Endpoints){
	e.load()
	return
}


type response struct {
	Code   int
	String string
}

func (r response) Body() []byte {
	return []byte(r.String)
}

func (e *Endpoints) load() {
	e.registerResponses()
}

func (e *Endpoints) registerResponses() {
	e.responses = map[string]response{
		Success:       {200, `{"code": 200, "message": "Asked and done"}`},
		NoContent:     {204, `{"code": 204, "message": "No content!"}`},
		Redirect:      {302, `{"code": 302, "message": "Never gonna give you up..."}`},
		Bad:           {400, `{"code": 400, "message": "Are you talking to me?"}`},
		Unauthorized:  {401, `{"code": 401, "message": "Who are you?"}`},
		Forbidden:     {403, `{"code": 403, "message": "Can't touch this!"}`},
		NotAllowed:    {405, `{"code": 405, "message": "This is not the way"}`},
		NotFound:      {404, `{"code": 404, "message": "Resource not found"}`},
		InternalError: {500, `{"code": 500, "message": "Get to the chopper!"}`},
	}
}

func (e Endpoints) DecodeJson(body io.Reader, entity interface{}) (err error) {
	d := json.NewDecoder(body)
	err = d.Decode(entity)
	return
}

func (e Endpoints) DecodeQueryParameters(dst interface{}, src map[string][]string) (err error) {
	d := schema.NewDecoder()
	err = d.Decode(dst, src)
	return
}

func (e Endpoints) EncodeJson(toMarshall interface{}, ident bool) (j []byte, err error) {
	if !ident {
		j, err = json.Marshal(toMarshall)
	} else {
		j, err = json.MarshalIndent(toMarshall, "", " ")
	}
	return
}

func (e *Endpoints) Response(w http.ResponseWriter, resp string, body []byte) {
	// get response from default responses
	r, ok := e.responses[resp]
	if !ok {
		r.Code = 500
		r.String = `{"code": 500, "message": "Response Not Implemented!"}`
	}
	w.WriteHeader(r.Code)
	if body != nil {
		body = r.Body()
	}
	_, err := w.Write(body)
	log.Errorf("failed writing body", err)
	return
}
