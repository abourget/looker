package looker

import (
	"github.com/abourget/looker/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new Looker client for your domain.
func New(domain string, formats strfmt.Registry) *client.LookerAPI30Reference {
	if formats == nil {
		formats = strfmt.Default
	}
	transport := httptransport.New(domain, "/api/3.0", []string{"https"})
	return New(transport, formats)
}
