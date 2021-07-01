# Valist Acceptance Tests

Tests are executed by [Sharness](https://github.com/valist-io/sharness).

## Usage

Run all tests.

```bash
make
```

Run a single test with verbose output.

```bash
./t0000-sharness.sh -v
```

## Writing Tests

Create a file with a name following the format `tXXXX-about.sh`.

```bash
#!/bin/sh

test_description="..."

: "${SHARNESS_TEST_SRCDIR:=lib/sharness}"

. "$SHARNESS_TEST_SRCDIR/sharness.sh"
```

Make the test file executable `chmod +x tXXXX-about.sh`.
