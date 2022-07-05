package delivery

import (
	"context"
	"github.com/mehdihadeli/store-golang-microservice-sample/pkg/mediatr"
	"github.com/mehdihadeli/store-golang-microservice-sample/services/catalogs/read_service/internal/shared/configurations"
	"github.com/segmentio/kafka-go"
)

type ProductConsumersBase struct {
	*configurations.Infrastructure
	Mediator *mediatr.Mediator
}

func NewProductConsumersBase(infra *configurations.Infrastructure, mediator *mediatr.Mediator) *ProductConsumersBase {
	return &ProductConsumersBase{Infrastructure: infra, Mediator: mediator}
}

func (pm *ProductConsumersBase) CommitMessage(ctx context.Context, r *kafka.Reader, m kafka.Message) {
	pm.Metrics.SuccessKafkaMessages.Inc()
	pm.Log.KafkaLogCommittedMessage(m.Topic, m.Partition, m.Offset)

	if err := r.CommitMessages(ctx, m); err != nil {
		pm.Log.WarnMsg("commitMessage", err)
	}
}

func (pm *ProductConsumersBase) LogProcessMessage(m kafka.Message, workerID int) {
	pm.Log.KafkaProcessMessage(m.Topic, m.Partition, string(m.Value), workerID, m.Offset, m.Time)
}

func (pm *ProductConsumersBase) CommitErrMessage(ctx context.Context, r *kafka.Reader, m kafka.Message) {
	pm.Metrics.ErrorKafkaMessages.Inc()
	pm.Log.KafkaLogCommittedMessage(m.Topic, m.Partition, m.Offset)
	if err := r.CommitMessages(ctx, m); err != nil {
		pm.Log.WarnMsg("commitMessage", err)
	}
}
