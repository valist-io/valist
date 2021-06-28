#!/bin/sh

test_description="Test npm"

: "${SHARNESS_TEST_SRCDIR:=lib/sharness}"

. "$SHARNESS_TEST_SRCDIR/sharness.sh"

npm_major_version=$(npm -v | cut -d "." -f -1)

test_expect_success "npm version 7 or greater" "
  test $npm_major_version -eq 7 ||
  test $npm_major_version -gt 7
"

test_expect_success "install an npm package from registry" "
  npm init -y &&
  npm install --registry=https://app.valist.io/api/npm @valist/sdk &&
  test $(npm ls | wc -l) -gt 3
"

test_done