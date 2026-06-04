package openapi

import (
	"encoding/json"
	"testing"
)

func TestSecuritySchemeUnmarshalReservedOpenAPIFields(t *testing.T) {
	var scheme SecurityScheme
	if err := json.Unmarshal([]byte(`{
		"type": "http",
		"in": "header",
		"scheme": "bearer",
		"bearerFormat": "JWT"
	}`), &scheme); err != nil {
		t.Fatalf("json.Unmarshal() error = %v", err)
	}
	if scheme.Get_Type() != "http" {
		t.Fatalf("type = %q, want http", scheme.Get_Type())
	}
	if scheme.Get_In() != "header" {
		t.Fatalf("in = %q, want header", scheme.Get_In())
	}
	if scheme.GetBearerFormat() != "JWT" {
		t.Fatalf("bearerFormat = %q, want JWT", scheme.GetBearerFormat())
	}
}

func TestSecuritySchemeUnmarshalThriftEscapedFields(t *testing.T) {
	var scheme SecurityScheme
	if err := json.Unmarshal([]byte(`{
		"_type": "apiKey",
		"_in": "query",
		"name": "token"
	}`), &scheme); err != nil {
		t.Fatalf("json.Unmarshal() error = %v", err)
	}
	if scheme.Get_Type() != "apiKey" {
		t.Fatalf("_type = %q, want apiKey", scheme.Get_Type())
	}
	if scheme.Get_In() != "query" {
		t.Fatalf("_in = %q, want query", scheme.Get_In())
	}
}
