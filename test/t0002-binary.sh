#!/bin/sh

test_description="Test binary install"

: "${SHARNESS_TEST_SRCDIR:=lib/sharness}"

. "$SHARNESS_TEST_SRCDIR/sharness.sh"

test_expect_success "install a binary from registry" "
  curl -L -o binary https://app.valist.io/api/test/binary/latest
"

test_done