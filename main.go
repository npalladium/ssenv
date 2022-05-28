package main

import (
	"flag"
	"log"
	"os"
)

func main() {
	exit := flag.Bool("exit", false, "Exit instead of unsetting environment vars")
	flag.Parse()
	var apikeys = []string{
		"AWS_ACCESS_KEY_ID",
		"AWS_SECRET_ACCESS_KEY",
		"GITHUB_TOKEN",
		"PRIVATE-TOKEN",
		"API_KEY",
		"TOKEN",
		"SECRET",
		"SECRET_TOKEN",
		"GOOGLE_API_KEY",
		"ATLAS_TOKEN",
		"authToken",
		"PLATFORMSH_CLI_TOKEN",
		"TRIPLYDB_TOKEN",
		"TX_TOKEN",
		"GOOGLE_APPLICATION_CREDENTIALS",
		"DAPR_API_TOKEN",
		"COMPOSER_AUTH",
		"DIGITALOCEAN_TOKEN",
		"DIGITALOCEAN_ACCESS_TOKEN",
		"APIFY_TOKEN",
		"PHRASE_ACCESS_TOKEN",
		"DOPPLER_TOKEN",
		"Token",
		"VERCEL_ARTIFACTS_TOKEN",
		"VAULT_TOKEN",
		"AUTH0_MGMT_API_TOKEN",
		"TFE_TOKEN",
		"K6_CLOUD_TOKEN",
		"K8S_AUTH_API_KEY",
		"HEROKU_API_KEY",
		"SENDGRID_API_KEY",
		"CLOUDFLARE_AUTH_TOKEN",
		"DATOCMS_API_TOKEN",
		"OVHCLOUD_WEBPAAS_CLI_TOKEN",
		"STRIPE_TOKEN",
		"SAS_VIYA_TOKEN",
		"ELSEVIER_SCOPUS_KEY",
		"UNIFORM_API_TOKEN",
		"REDIVIS_API_TOKEN",
		"POLARIS_ACCESS_TOKEN",
		"ACCESS-TOKEN",
		"USYM_UPLOAD_AUTH_TOKEN",
		"LONGBRIDGE_ACCESS_TOKEN",
		"GHTOKEN",
		"KAGGLE_KEY",
		"GH_TOKEN",
		"SPOTIPY_CLIENT_ID",
		"SPOTIPY_CLIENT_SECRET",
		"AZURE_STORAGE_KEY",
		"AZURE_STORAGE_SAS_TOKEN",
		"TWITTER_SECRET",
		"TWITTER_KEY",
		"REACT_APP_API_KEY",
		"NPM_TOKEN",
		"PRIVATE_KEY",
		"OPENAI_API_KEY",
		"PSONO_CI_API_KEY_ID",
		"CPLEX_STUDIO_KEY",
		"api-token",
		"PYPI_TOKEN",
		"IEX_TOKEN",
		"TMDB_API_TOKEN",
		"FRED_API_KEY",
		"NEPTUNE_API_TOKEN",
		"GITHUB_COM_TOKEN",
		"CLOUDSMITH_API_KEY",
		"SLACK_BOT_TOKEN",
	}
	for _, s := range apikeys {
		_, present := os.LookupEnv(s)
		if present {
			if *exit {
				log.Fatalf("Found key: [%s]", s)
			}
			log.Printf("Unsetting key: [%s]", s)
			os.Unsetenv(s)
		}
	}
}
