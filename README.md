# Safe Sanitised Env

Exit if environment variables known to be used as access tokens are set; otherwise works similar to `env`.

May be used before executing anything that might attempt to steal sensitive
environment variables, like hijacked python or node modules ;)

# FAQs

## Why?

I came across [abbbi's ssenv](https://github.com/abbbi/ssenv) which checks for environment that represent access keys and exits if they are present. For me, the logical extension of this was an `env`-like tool. (If you have been using it to only print environment variables, you have been missing out; env's man page says it can "set environment and execute command, or print environment". As an aside, you may have used this functionality in your shebangs such as `#!/usr/bin/env bash`.)

## Do you use it?

Currently, yes. But I may stop using it once the novelty wears off.

## Will it really secure me against malicious tools?

I'm not really sure. Keep in mind that:
- If you have a malicious tool installed, you may have bigger problems.
- The "jail" environment created by this tool is probably escapable.

# License

Portions Copyright (C) 2014, The GO-Coreutils Developers.
Portions Copyright (C) 2022, Michael Ablassmeier (abbbi).
Portions Copyright (C) 2022, npalladium

ssenv is licensed under the GNU General Public License v3.

This program comes with ABSOLUTELY NO WARRANTY; for details see
LICENSE. This is free software, and you are welcome to redistribute
it under certain conditions in LICENSE.
