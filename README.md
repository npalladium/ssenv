# Safe Sanitised Env

Exit if environment variables known to be used as access tokens are set; otherwise works similar to `env`.

May be used before executing anything that might attempt to steal sensitive
environment variables, like hijacked python or node modules ;)

# Why?

Just to have a list of environment variables used by libraries to read your
sensitive authentication tokens, specifically, no wildcards used.

# License

Portions Copyright (C) 2014, The GO-Coreutils Developers.
Portions Copyright (C) 2022, Michael Ablassmeier (abbbi).
Portions Copyright (C) 2022, npalladium

ssenv is licensed under the GNU General Public License v3.

This program comes with ABSOLUTELY NO WARRANTY; for details see
LICENSE. This is free software, and you are welcome to redistribute
it under certain conditions in LICENSE.
