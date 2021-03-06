package ultradns

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/terra-farm/udnssdk"
)

// Provider returns a terraform.ResourceProvider.
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"username": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("ULTRADNS_USERNAME", nil),
				Description: "UltraDNS Username.",
			},

			"password": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("ULTRADNS_PASSWORD", nil),
				Description: "UltraDNS User Password",
			},
			"baseurl": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("ULTRADNS_BASEURL", nil),
				Default:     udnssdk.DefaultLiveBaseURL,
				Description: "UltraDNS Base URL",
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"ultradns_dirpool":    resourceUltradnsDirpool(),
			"ultradns_probe_http": resourceUltradnsProbeHTTP(),
			"ultradns_probe_ping": resourceUltradnsProbePing(),
			"ultradns_record":     resourceUltradnsRecord(),
			"ultradns_tcpool":     resourceUltradnsTcpool(),
			"ultradns_rdpool":     resourceUltradnsRdpool(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		Username: d.Get("username").(string),
		Password: d.Get("password").(string),
		BaseURL:  d.Get("baseurl").(string),
	}

	return config.Client()
}
