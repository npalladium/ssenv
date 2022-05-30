package main

import (
	"log"
	"os"
)

func main() {
	var apikeys = []string{
		"ACCESS-TOKEN",
		"APIFY_TOKEN",
		"API_KEY",
		"api-token",
		"ATLAS_TOKEN",
		"AUTH0_MGMT_API_TOKEN",
		"authToken",
		"AUTH_TOKEN",
		"AWS_ACCESS_KEY_ID",
		"AWS_SECRET_ACCESS_KEY",
		"AZURE_STORAGE_KEY",
		"AZURE_STORAGE_SAS_TOKEN",
		"CALIBRE_API_TOKEN",
		"CARGO_REGISTRY_TOKEN",
		"CIRCLE_TOKEN",
		"CLOUDFLARE_AUTH_TOKEN",
		"CLOUDSMITH_API_KEY",
		"COMPOSER_AUTH",
		"CPLEX_STUDIO_KEY",
		"DAPR_API_TOKEN",
		"DATOCMS_API_TOKEN",
		"DIGITALOCEAN_ACCESS_TOKEN",
		"DIGITALOCEAN_TOKEN",
		"DOPPLER_TOKEN",
		"DX_SECURITY_CONTEXT",
		"ELSEVIER_SCOPUS_KEY",
		"FRED_API_KEY",
		"GENIUS_API_TOKEN",
		"GH_TOKEN",
		"GHTOKEN",
		"GITHUB_COM_TOKEN",
		"GITHUB_TOKEN",
		"GOOGLE_API_KEY",
		"GOOGLE_APPLICATION_CREDENTIALS",
		"hcloud_token", //hetzner cloud
		"HEROKU_API_KEY",
		"IEX_TOKEN",
		"JB_SPACE_CLIENT_TOKEN",
		"K6_CLOUD_TOKEN",
		"K8S_AUTH_API_KEY",
		"KAGGLE_KEY",
		"LAST_FM_API_SECRET",
		"LONGBRIDGE_ACCESS_TOKEN",
		"NEPTUNE_API_TOKEN",
		"NPM_TOKEN",
		"NUTANIX_SESSION_AUTH",
		"OPENAI_API_KEY",
		"OVHCLOUD_WEBPAAS_CLI_TOKEN",
		"OVIRT_TOKEN",
		"PERCY_TOKEN",
		"PHRASE_ACCESS_TOKEN",
		"PLATFORMIO_AUTH_TOKEN",
		"PLATFORMSH_CLI_TOKEN",
		"POLARIS_ACCESS_TOKEN",
		"PRIVATE_KEY",
		"PRIVATE-TOKEN",
		"PSONO_CI_API_KEY_ID",
		"PYPI_TOKEN",
		"REACT_APP_API_KEY",
		"REDIVIS_API_TOKEN",
		"RS_TOKEN",
		"SAS_VIYA_TOKEN",
		"SECRET",
		"SECRET_TOKEN",
		"SENDGRID_API_KEY",
		"SLACK_BOT_TOKEN",
		"SNYK_TOKEN",
		"SPOTIPY_CLIENT_ID",
		"SPOTIPY_CLIENT_SECRET",
		"STRIPE_TOKEN",
		"TAP_FACEBOOK_ACCESS_TOKEN", //facebook marketing api
		"TFE_TOKEN",
		"TMDB_API_TOKEN",
		"Token",
		"TOKEN",
		"TRIPLYDB_TOKEN",
		"TWITTER_KEY",
		"TWITTER_SECRET",
		"TX_TOKEN",
		"UNIFORM_API_TOKEN",
		"USYM_UPLOAD_AUTH_TOKEN",
		"VAULT_TOKEN",
		"VERCEL_ARTIFACTS_TOKEN",
	}
	for _, s := range apikeys {
		_, present := os.LookupEnv(s)
		if present {
			log.Fatalf("Found key: [%s]", s)
		}
	}
}
