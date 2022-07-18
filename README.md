# Safe Sanitised Env

Exit if environment variables known to be used as access tokens are set; otherwise works similar to `env`.

May be used before executing anything that might attempt to steal sensitive
environment variables, like hijacked python or node modules ;)

# Why?

Just to have a list of environment variables used by libraries to read your
sensitive authentication tokens, specifically, no wildcards used.
