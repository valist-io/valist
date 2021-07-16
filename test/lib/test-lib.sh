: "${SHARNESS_TEST_SRCDIR:=lib/sharness}"
. "$SHARNESS_TEST_SRCDIR/sharness.sh"

HARDHAT_SRC=../../hardhat
VALIST_CLI=../../cli/bin/run

test_start_blockchain() {
  # need to sleep here because the telemetry prompt
  test_expect_success "blockchain starts" '
    npm run --prefix $HARDHAT_SRC blockchain &
    sleep 5 &&
    HARDHAT_PID=$!
  '
}

test_stop_blockchain() {
  test_expect_success "blockchain stops" '
    kill $HARDHAT_PID
  '
}

test_deploy_contract() {
  test_expect_success "contract deploy success" '
    npm run --prefix $HARDHAT_SRC deploy:local
  '
}

test_valist_cli() {
  exec $VALIST_CLI $@
}