#!/bin/sh

test_description="Test pip install"

: "${SHARNESS_TEST_SRCDIR:=lib/sharness}"

. "$SHARNESS_TEST_SRCDIR/sharness.sh"

test_expect_success "install a pip package from registry" "
  pip3 install https://app.valist.io/api/test/pip/latest
"

test_done