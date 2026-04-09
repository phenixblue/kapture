package goeval

import (
	"context"

	"github.com/phenixblue/kapture/internal/checks"
	"github.com/phenixblue/kapture/internal/eval"
)

type Engine struct{}

func New() *Engine {
	return &Engine{}
}

func (e *Engine) Name() string {
	return "go"
}

func (e *Engine) Evaluate(ctx context.Context, req eval.RunRequest) (checks.RunResult, error) {
	return checks.RunFiltered(ctx, req.Registry, req.Filter)
}
