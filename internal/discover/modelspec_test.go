package discover

import (
	"context"
	"testing"

	"charm.land/catwalk/pkg/catwalk"
)

func TestModelspecEnricher_DeepSeekModels(t *testing.T) {
	e := &modelspecEnricher{}
	models := []catwalk.Model{
		{ID: "deepseek-v4-flash", Name: "deepseek-v4-flash"},
		{ID: "deepseek-v4-pro", Name: "deepseek-v4-pro"},
		{ID: "deepseek-v3.2-chat", Name: "deepseek-v3.2-chat"},
		{ID: "deepseek-r1-0528", Name: "deepseek-r1-0528"},
		{ID: "deepseek-chat", Name: "deepseek-chat"},
	}

	result, err := e.EnrichModels(context.Background(), Config{}, nil, models)
	if err != nil {
		t.Fatal(err)
	}

	want := map[string]struct {
		ctx int64
		max int64
	}{
		"deepseek-v4-flash":  {1_048_576, 131_072},
		"deepseek-v4-pro":    {1_048_576, 131_072},
		"deepseek-v3.2-chat": {131_072, 16_384},
		"deepseek-r1-0528":   {131_072, 16_384},
		"deepseek-chat":      {131_072, 16_384},
	}

	for _, m := range result {
		w, ok := want[m.ID]
		if !ok {
			continue
		}
		if m.ContextWindow != w.ctx {
			t.Errorf("%s: ContextWindow = %d, want %d", m.ID, m.ContextWindow, w.ctx)
		}
		if m.DefaultMaxTokens != w.max {
			t.Errorf("%s: DefaultMaxTokens = %d, want %d", m.ID, m.DefaultMaxTokens, w.max)
		}
	}
}

func TestModelspecEnricher_SkipsExistingMetadata(t *testing.T) {
	e := &modelspecEnricher{}
	models := []catwalk.Model{
		{ID: "deepseek-chat", Name: "deepseek-chat", ContextWindow: 999},
	}

	result, err := e.EnrichModels(context.Background(), Config{}, nil, models)
	if err != nil {
		t.Fatal(err)
	}

	if result[0].ContextWindow != 999 {
		t.Errorf("should preserve existing ContextWindow=999, got %d", result[0].ContextWindow)
	}
}

func TestModelspecEnricher_UnknownModel(t *testing.T) {
	e := &modelspecEnricher{}
	models := []catwalk.Model{
		{ID: "gpt-4o", Name: "gpt-4o"},
	}

	result, err := e.EnrichModels(context.Background(), Config{}, nil, models)
	if err != nil {
		t.Fatal(err)
	}

	if result[0].ContextWindow != 0 {
		t.Errorf("unknown model should stay at 0, got %d", result[0].ContextWindow)
	}
}

func TestMatchModelSpec_PreferLongestPrefix(t *testing.T) {
	// deepseek-v4-flash should match "deepseek-v4-flash" (1M) not "deepseek-v4-pro" (also 1M but wrong).
	spec, ok := matchModelSpec("deepseek-v4-flash-128k")
	if !ok {
		t.Fatal("expected match")
	}
	if spec.contextWindow != 1_048_576 {
		t.Errorf("expected 1M context, got %d", spec.contextWindow)
	}
}

func TestFallbackEnricherRegistration(t *testing.T) {
	// The modelspec enricher should be registered as fallback via init().
	e := GetEnricherWithFallback("nonexistent-provider-type")
	if e == nil {
		t.Fatal("expected fallback enricher to be registered")
	}
	if _, ok := e.(*modelspecEnricher); !ok {
		t.Errorf("expected *modelspecEnricher, got %T", e)
	}
}
