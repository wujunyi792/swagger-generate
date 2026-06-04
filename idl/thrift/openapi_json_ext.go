package openapi

import "encoding/json"

// UnmarshalJSON keeps OpenAPI reserved words usable from JSON based option
// parsing. The generated thrift fields for type/in are named _Type/_In to
// avoid Go keywords, which makes them invisible to encoding/json by default.
func (p *SecurityScheme) UnmarshalJSON(data []byte) error {
	type securitySchemeJSON struct {
		Type                   string      `json:"type"`
		UnderscoreType         string      `json:"_type"`
		Description            string      `json:"description"`
		Name                   string      `json:"name"`
		In                     string      `json:"in"`
		UnderscoreIn           string      `json:"_in"`
		Scheme                 string      `json:"scheme"`
		BearerFormat           string      `json:"bearer_format"`
		BearerFormatCamel      string      `json:"bearerFormat"`
		Flows                  *OauthFlows `json:"flows"`
		OpenIDConnectURL       string      `json:"open_id_connect_url"`
		OpenIDConnectURLCamel  string      `json:"openIdConnectUrl"`
		SpecificationExtension []*NamedAny `json:"specification_extension"`
	}
	var aux securitySchemeJSON
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	p._Type = firstNonEmpty(aux.Type, aux.UnderscoreType)
	p.Description = aux.Description
	p.Name = aux.Name
	p._In = firstNonEmpty(aux.In, aux.UnderscoreIn)
	p.Scheme = aux.Scheme
	p.BearerFormat = firstNonEmpty(aux.BearerFormat, aux.BearerFormatCamel)
	p.Flows = aux.Flows
	p.OpenIDConnectURL = firstNonEmpty(aux.OpenIDConnectURL, aux.OpenIDConnectURLCamel)
	p.SpecificationExtension = aux.SpecificationExtension
	return nil
}

func firstNonEmpty(values ...string) string {
	for _, value := range values {
		if value != "" {
			return value
		}
	}
	return ""
}
