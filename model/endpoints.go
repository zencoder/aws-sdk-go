package model

import (
	"bytes"
	"go/format"
	"io"
	"text/template"
)

// Generate writes a Go file to the given writer.
func GenerateEndpoints(endpoints interface{}, w io.Writer) error {
	tmpl, err := template.New("endpoints").Parse(t)
	if err != nil {
		return err
	}

	out := bytes.NewBuffer(nil)
	if err := tmpl.Execute(out, endpoints); err != nil {
		return err
	}

	b, err := format.Source(bytes.TrimSpace(out.Bytes()))
	if err != nil {
		return err
	}

	_, err = io.Copy(w, bytes.NewReader(b))
	return err
}

const t = `
package aws

type endpointStruct struct {
	Version   int
	Endpoints map[string]endpointEntry
}

type endpointEntry struct {
	Endpoint string
}

var endpointMap = endpointStruct{
	Version: {{ .Version }},
	Endpoints: map[string]endpointEntry{
		{{ range $key, $entry := .Endpoints }}"{{ $key }}": endpointEntry{
			Endpoint: "{{ $entry.Endpoint }}",
		},
		{{ end }}
	},
}
`
