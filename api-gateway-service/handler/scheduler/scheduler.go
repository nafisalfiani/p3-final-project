package scheduler

import (
	"context"

	"github.com/nafisalfiani/p3-final-project/api-gateway-service/usecase"
	"github.com/nafisalfiani/p3-final-project/lib/auth"
	"github.com/nafisalfiani/p3-final-project/lib/log"
	"github.com/nafisalfiani/p3-final-project/lib/scheduler"
)

type Config struct {
	HelloWorld scheduler.TaskConf
}

type Interface interface {
	Run()
	TriggerScheduler(name string) error
}

type schedule struct {
	scheduler scheduler.Interface
	log       log.Interface
}

func Init(conf Config, log log.Interface, auth auth.Interface, uc *usecase.Usecases) Interface {
	s := &schedule{
		scheduler: scheduler.Init(log, auth),
		log:       log,
	}

	s.scheduler.AssignTask(conf.HelloWorld, s.HelloWorld)

	return s
}

func (s *schedule) Run() {
	s.scheduler.Run()
}

func (s *schedule) TriggerScheduler(name string) error {
	return s.scheduler.TriggerScheduler(name)
}

func (s *schedule) HelloWorld(ctx context.Context) error {
	s.log.Info(ctx, "Hello, 世界!")

	return nil
}
