package eval

import (
	"context"

	"github.com/phenixblue/kapture/internal/checks"
	"github.com/phenixblue/kapture/internal/kube"
)

type RunRequest struct {
	Registry        []checks.Check
	Filter          checks.Filter
	PolicyFile      string
	PolicyBundle    string
	ClusterSnapshot *kube.ClusterSnapshot
}

type Evaluator interface {
	Evaluate(ctx context.Context, req RunRequest) (checks.RunResult, error)
	Name() string
}
