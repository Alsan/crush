package discover

import (
	"context"
	"strings"

	"charm.land/catwalk/pkg/catwalk"
)

func init() {
	RegisterFallbackEnricher(&modelspecEnricher{})
}

// modelSpec holds known metadata for a model whose /v1/models endpoint
// does not return context window or max token information.
type modelSpec struct {
	contextWindow    int64
	defaultMaxTokens int64
}

// knownModelSpecs maps model ID prefixes to their known metadata.
// A discovered model matches if its ID starts with the key. More
// specific (longer) prefixes take precedence when multiple match.
// Only models whose APIs omit metadata need entries here.
var knownModelSpecs = map[string]modelSpec{
	// DeepSeek — /v1/models returns only {id, object, owned_by}.
	"deepseek-v4-flash": {contextWindow: 1_048_576, defaultMaxTokens: 131_072},
	"deepseek-v4-pro":   {contextWindow: 1_048_576, defaultMaxTokens: 131_072},
	"deepseek-v3.2":     {contextWindow: 131_072, defaultMaxTokens: 16_384},
	"deepseek-v3":       {contextWindow: 131_072, defaultMaxTokens: 16_384},
	"deepseek-r1":       {contextWindow: 131_072, defaultMaxTokens: 16_384},
	"deepseek-chat":     {contextWindow: 131_072, defaultMaxTokens: 16_384},
	"deepseek-coder":    {contextWindow: 131_072, defaultMaxTokens: 16_384},
}

// modelspecEnricher fills in model metadata from a static lookup table
// of known model ID prefixes. It runs as a fallback enricher for
// providers that do not have their own enricher (e.g. cloud providers
// whose /v1/models endpoint omits context_window and max_tokens).
type modelspecEnricher struct{}

func (e *modelspecEnricher) EnrichModels(_ context.Context, _ Config, _ Resolver, models []catwalk.Model) ([]catwalk.Model, error) {
	for i := range models {
		if models[i].ContextWindow != 0 {
			continue // User-specified metadata takes precedence.
		}
		if spec, ok := matchModelSpec(models[i].ID); ok {
			models[i].ContextWindow = spec.contextWindow
			if spec.defaultMaxTokens > 0 && models[i].DefaultMaxTokens == 0 {
				models[i].DefaultMaxTokens = spec.defaultMaxTokens
			}
		}
	}
	return models, nil
}

// matchModelSpec finds the longest matching prefix in knownModelSpecs.
func matchModelSpec(modelID string) (modelSpec, bool) {
	id := strings.ToLower(modelID)
	var bestKey string
	var bestSpec modelSpec
	for prefix, spec := range knownModelSpecs {
		if strings.HasPrefix(id, prefix) && len(prefix) > len(bestKey) {
			bestKey = prefix
			bestSpec = spec
		}
	}
	if bestKey == "" {
		return modelSpec{}, false
	}
	return bestSpec, true
}
// test
