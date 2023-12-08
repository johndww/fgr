package routes

import (
	"encoding/json"
	"github.com/microcosm-cc/bluemonday"
	"github.com/pkg/errors"
	"io"
	"net/http"
)

var p *bluemonday.Policy

func init() {
	p = bluemonday.UGCPolicy()
}

func ReadBody(body io.ReadCloser, output any) error {
	bytes, err := io.ReadAll(body)
	if err != nil {
		return errors.Wrap(err, "unable to read bytes")
	}

	err = json.Unmarshal(bytes, output)
	if err != nil {
		return errors.Wrap(err, "unable to unmarshal json")
	}

	SanitizeFields(output)

	return nil
}

func WriteResponse(w http.ResponseWriter, input any) error {
	bytes, err := json.Marshal(input)
	if err != nil {
		return errors.Wrap(err, "unable to marshal json")
	}

	_, err = w.Write(bytes)
	if err != nil {
		return errors.Wrap(err, "unable to write bytes")
	}
	return nil
}
