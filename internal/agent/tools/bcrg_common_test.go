package tools

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

// TestBoundedWriterCapsOutput pins that a boundedWriter keeps at most limit
// bytes and always reports writes as fully consumed (no io.ErrShortWrite).
func TestBoundedWriterCapsOutput(t *testing.T) {
	t.Parallel()

	var buf bytes.Buffer
	w := &boundedWriter{w: &buf, limit: 8}

	n, err := w.Write([]byte("hello"))
	require.NoError(t, err)
	require.Equal(t, 5, n, "full write returns consumed length")

	n, err = w.Write([]byte("world-extra")) // 11 bytes
	require.NoError(t, err)
	require.Equal(t, 11, n, "writes beyond cap report full length (no short write)")
	require.Equal(t, "hellowor", buf.String(), "buffer is capped at limit")
}
