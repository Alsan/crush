package tools

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

// TestTokensaveLive runs the tool against a real tokensave graph. Opt-in
// via TOKENSAVE_LIVE_ROOT pointing at an indexed repository root.
func TestTokensaveLive(t *testing.T) {
	root := os.Getenv("TOKENSAVE_LIVE_ROOT")
	if root == "" {
		t.Skip("TOKENSAVE_LIVE_ROOT not set")
	}
	ctx := context.Background()

	out, err := runTokensave(ctx, root, &TokensaveParams{Op: "status"})
	require.NoError(t, err)
	t.Log(out)

	out, err = runTokensave(ctx, root, &TokensaveParams{Op: "search", Query: "coordinator", Limit: 5})
	require.NoError(t, err)
	t.Log(out)

	out, err = runTokensave(ctx, root, &TokensaveParams{Op: "find", Name: "buildTools", Limit: 3})
	require.NoError(t, err)
	require.Contains(t, out, `"name": "buildTools"`)
	t.Log(out)

	out, err = runTokensave(ctx, root, &TokensaveParams{Op: "body", Name: "buildTools"})
	require.NoError(t, err)
	require.Contains(t, out, "func (c *coordinator) buildTools")

	out, err = runTokensave(ctx, root, &TokensaveParams{Op: "callees", ID: "struct_method:buildtools", Limit: 3})
	if err != nil {
		t.Logf("callees by bogus id: %v (expected)", err)
	}
	t.Log(out)
}
