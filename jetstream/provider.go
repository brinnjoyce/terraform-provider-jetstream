package jetstream

import (
	"regexp"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

var streamIdRegex = regexp.MustCompile("^JETSTREAM_STREAM_(.+)$")
var consumerIdRegex = regexp.MustCompile("^JETSTREAM_STREAM_(.+?)_CONSUMER_(.+)$")
var streamTemplateIdRegex = regexp.MustCompile("^JETSTREAM_STREAMTEMPLATE_(.+)$")
var kvIdRegex = regexp.MustCompile("^JETSTREAM_KV_(.+)$")

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"servers": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Comma separated list of NATS servers to connect to",
				DefaultFunc: schema.EnvDefaultFunc("NATS_URL", nil),
			},
			"credentials": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Path to the NATS 2.0 credentials file to use for authentication",
				DefaultFunc: schema.EnvDefaultFunc("NATS_CREDS", nil),
			},
			"credential_data": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The contents of the NATS 2.0 Credentials file to use",
			},
			"tls": {
				Type:     schema.TypeSet,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ca_file": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Path to the server root CA file (in PEM format). Needed when the NATS server has TLS enabled with an unknown root CA",
						},
						"ca_file_data": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The root CA (file) content, in PEM format. Needed when the NATS server has TLS enabled with an unknown root CA",
						},
					},
				},
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"jetstream_stream":          resourceStream(),
			"jetstream_consumer":        resourceConsumer(),
			"jetstream_stream_template": resourceStreamTemplate(),
			"jetstream_kv_bucket":       resourceKVBucket(),
		},

		ConfigureFunc: connectMgr,
	}
}
