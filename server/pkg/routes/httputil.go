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
	unsafeBytes, err := io.ReadAll(body)
	if err != nil {
		return errors.Wrap(err, "unable to read bytes")
	}

	bytes := p.SanitizeBytes(unsafeBytes)

	err = json.Unmarshal(bytes, output)
	if err != nil {
		return errors.Wrap(err, "unable to unmarshal json")
	}
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
