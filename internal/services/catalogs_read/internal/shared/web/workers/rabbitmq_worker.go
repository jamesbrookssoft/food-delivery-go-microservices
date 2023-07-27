package workers

import (
    "context"

    "github.com/mehdihadeli/go-ecommerce-microservices/internal/pkg/logger"
    "github.com/mehdihadeli/go-ecommerce-microservices/internal/pkg/messaging/bus"
    "github.com/mehdihadeli/go-ecommerce-microservices/internal/pkg/web"
)

func NewRabbitMQWorker(logger logger.Logger, bus bus.Bus) web.Worker {
	return web.NewBackgroundWorker(func(ctx context.Context) error {
		err := bus.Start(ctx)
		if err != nil {
			logger.Errorf("[RabbitMQWorkerWorker.Start] error in the starting rabbitmq worker: {%v}", err)
			return err
		}
		return nil
	}, func(ctx context.Context) error {
		return bus.Stop(ctx)
	})
}
