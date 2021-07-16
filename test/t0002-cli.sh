#!/bin/sh

test_description="Run basic cli commands"

. "lib/test-lib.sh"

test_expect_success "returns error on empty address" '
  test_valist_cli account:get
'

test_done