package requests

import (
	"encoding/json"
	"io"
	"net/http"
)

func NewMarshall(w http.ResponseWriter, variable any) ([]byte, error) {
	w.Header().Set("Content-Type", "application/json")
	return json.MarshalIndent(variable, "", "    ")
}

func NewDec(r *http.Request, variable any) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, variable)
}
