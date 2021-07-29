package build

import (
	"context"
	"testing"

	"github.com/moby/buildkit/client/llb"
	"github.com/moby/buildkit/frontend/gateway/client"
	"github.com/moby/buildkit/util/testutil/integration"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

func runBuildkit(t *testing.T, sb integration.Sandbox) {
	ctx := sb.Context()

	c, err := New(ctx, sb.Address())
	require.NoError(t, err)
	defer c.Close()

	product := "buildkit_test"
	optKey := "test-string"

	b := func(ctx context.Context, c client.Client) (*client.Result, error) {
		if c.BuildOpts().Product != product {
			return nil, errors.Errorf("expected product %q, got %q", product, c.BuildOpts().Product)
		}
		opts := c.BuildOpts().Opts
		testStr, ok := opts[optKey]
		if !ok {
			return nil, errors.Errorf(`build option %q missing`, optKey)
		}

		run := llb.Image("busybox:latest").Run(
			llb.ReadonlyRootFS(),
			llb.Args([]string{"/bin/sh", "-ec", `echo -n '` + testStr + `' > /out/foo`}),
		)
		st := run.AddMount("/out", llb.Scratch())

		def, err := st.Marshal(ctx)
		if err != nil {
			return nil, errors.Wrap(err, "failed to marshal state")
		}
	}
}
