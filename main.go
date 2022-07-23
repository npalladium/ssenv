package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var (
	version = "none"
	commit  = "none"
)

const (
	helpText = `Usage: env [OPTION]... [-] [NAME=VALUE]... [COMMAND [ARG]...]
Set each NAME to VALUE in the environment and run COMMAND.
  -ignore-environment  start with an empty environment
  -null           end each output line with 0 byte rather than newline
  -sanitise       delete the sensitive environment variables
  -unset=NAME     remove variable from the environment
  -help           display this help and exit
  -version        output version information and exit
`
)

var versionText = "Version: " + version + ". " + "Commit: " + commit + "." + `

Portions Copyright (C) 2014, The GO-Coreutils Developers.
Portions Copyright (C) 2022, Michael Ablassmeier (abbbi).
Portions Copyright (C) 2022, npalladium

ssenv is licensed under the GNU General Public License v3.
`

func apiKeys() []string {
	return []string{
		"ACCESS-TOKEN",
		"APIFY_TOKEN",
		"API_KEY",
		"api-token",
		"ATLAS_PRIVATE_KEY",
		"ATLAS_TOKEN",
		"AUTH0_MGMT_API_TOKEN",
		"authToken",
		"AUTH_TOKEN",
		"AWS_ACCESS_KEY_ID",
		"AWS_SECRET_ACCESS_KEY",
		"ASANA_CLIENT_SECRET",
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
		"EXPO_TOKEN",
		"FRED_API_KEY",
		"FIREFLY_III_ACCESS_TOKEN",
		"FLY_ACCESS_TOKEN",
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
		"OPSGENIE_API_KEY",
		"OVHCLOUD_WEBPAAS_CLI_TOKEN",
		"OVIRT_TOKEN",
		"PARTICLE_AUTH",
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
		"SRC_ACCESS_TOKEN",
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
}

var (
	help       = flag.Bool("help", false, "help")
	versionOpt = flag.Bool("version", false, "version_text")
	ignoreEnv  = flag.Bool("ignore-environment", false, "start with an empty environment")
	nullOpt    = flag.Bool("null", false, "end each output line with 0 byte rather than newline")
	sanitise   = flag.Bool("sanitise", false, "remove the api keys from env")
	unset      = flag.String("unset", "", "remove variable from the environment")
	environ    = os.Environ()
)

func setenv(name, value string) {
	err := os.Setenv(name, value)
	if err != nil {
		return
	}
	for i := 0; i < len(environ); i++ {
		e := strings.SplitN(environ[i], "=", 2)
		if e[0] == name {
			environ[i] = name + "=" + value
			return
		}
	}
	environ = append(environ, name+"="+value)
}
func unsetenv(name string) {
	_ = os.Unsetenv(name)
	for i := 0; i < len(environ); {
		e := strings.SplitN(environ[i], "=", 2)
		if e[0] == name && len(e) == 2 {
			environ = append(environ[:i], environ[i+1:]...) // delete
		} else {
			i++
		}
	}
}

func tokenInEnv() bool {
	for _, s := range apiKeys() {
		_, present := os.LookupEnv(s)
		if present {
			fmt.Printf("Found key: [%s]\n", s)
			return true
		}
	}
	return false
}

func main() {
	optNullTerminateOutput := false
	flag.Parse()
	if *help {
		fmt.Println(helpText)
		os.Exit(0)
	}
	if *versionOpt {
		fmt.Println(versionText)
		os.Exit(0)
	}
	if *ignoreEnv {
		environ = make([]string, 0)
	}
	if *nullOpt {
		optNullTerminateOutput = true
	}
	if *unset != "" {
		unsetenv(*unset)
	}
	if *sanitise {
		for _, key := range apiKeys() {
			unsetenv(key)
		}
	}
	if !*sanitise && tokenInEnv() {
		os.Exit(1)
	}
	arg := flag.Args()
	if len(arg) >= 1 && arg[0] == "-" {
		environ = make([]string, 0)
		arg = arg[1:]
	}
	if len(arg) >= 1 {
		for i, _ := range arg {
			if strings.Index(arg[i], "=") > 0 {
				e := strings.SplitN(arg[i], "=", 2)
				setenv(e[0], e[1])
			} else {
				// run COMMAND
				if optNullTerminateOutput {
					fmt.Println("ssenv: cannot specify -null with command: No such file or directory")
					fmt.Println("Try 'ssenv -help' for more information.")
					os.Exit(1)
				}
				cmd := exec.Command(arg[i], arg[i+1:]...)
				cmd.Stdin = os.Stdin
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				cmd.Env = environ
				err := cmd.Run()
				if err != nil {
					_, lookPathErr := exec.LookPath(arg[i])
					if lookPathErr != nil {
						fmt.Printf("env: %s: No such file or directory\n", arg[i])
					} else {
						fmt.Println(err)
					}
					os.Exit(1)
				}
				os.Exit(0)
			}
		}
	}
	// print all Environment
	for _, s := range environ {
		if optNullTerminateOutput {
			fmt.Printf("%v%c", s, '\000')
		} else {
			fmt.Println(s)
		}
	}
}
