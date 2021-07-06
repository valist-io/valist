#!/bin/sh

test_description="Test npm install"

: "${SHARNESS_TEST_SRCDIR:=lib/sharness}"

. "$SHARNESS_TEST_SRCDIR/sharness.sh"

test_expect_success "install an npm package from registry" "
  npm init -y &&
  npm install --registry=https://app.valist.io/api/npm @valist/sdk
"

test_done