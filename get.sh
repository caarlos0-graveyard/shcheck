#!/bin/sh
set -e

TAR_FILE="/tmp/shcheck.tar.gz"
DOWNLOAD_URL="https://github.com/caarlos0/shcheck/releases/download"

get_latest() {
	url="https://api.github.com/repos/caarlos0/shcheck/releases/latest"
	if test -n "$GITHUB_TOKEN"; then
		curl --fail -sSL -H "Authorization: token $GITHUB_TOKEN" "$url"
	else
		curl --fail -sSL "$url"
	fi
}

last_version() {
	get_latest |
		grep tag_name |
		cut -f4 -d'"'
}

download() {
	test -z "$VERSION" && VERSION="$(last_version)"
	test -z "$VERSION" && {
		echo "Unable to get caarlos0/shcheck version." >&2
		get_latest
		exit 1
	}
	rm -f "$TAR_FILE"
	curl -s -L -o "$TAR_FILE" \
		"$DOWNLOAD_URL/$VERSION/shcheck_$(uname -s)_$(uname -m).tar.gz"
}

download
tar -xf "$TAR_FILE" -C /tmp
# shellcheck disable=SC2048,SC2086
/tmp/shcheck $*
