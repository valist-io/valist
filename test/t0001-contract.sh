#!/bin/sh

test_description="Start blockchain and deploy contract"

. "lib/test-lib.sh"

test_start_blockchain

test_deploy_contract

test_stop_blockchain

test_done