package tasks

import "context"

type TaiJiTask interface {
	Start(ctx context.Context) error
	StartWithStep(ctx context.Context, step TaiJiTaskStep) error
}

type taiJiTask struct {
	step TaiJiTaskStep
}

func CreateTaiJiTask(fs func(ctx context.Context) error) TaiJiTask {

	return &taiJiTask{}
}

func (t *taiJiTask) Start(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (t *taiJiTask) StartWithStep(ctx context.Context, step TaiJiTaskStep) error {
	//TODO implement me
	panic("implement me")
}
