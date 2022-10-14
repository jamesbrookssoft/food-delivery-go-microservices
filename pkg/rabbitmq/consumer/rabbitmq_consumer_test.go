package consumer

import (
	"context"
	"fmt"
	"github.com/mehdihadeli/store-golang-microservice-sample/pkg/core/serializer/json"
	"github.com/mehdihadeli/store-golang-microservice-sample/pkg/logger/defaultLogger"
	messageConsumer "github.com/mehdihadeli/store-golang-microservice-sample/pkg/messaging/consumer"
	"github.com/mehdihadeli/store-golang-microservice-sample/pkg/messaging/pipeline"
	types2 "github.com/mehdihadeli/store-golang-microservice-sample/pkg/messaging/types"
	"github.com/mehdihadeli/store-golang-microservice-sample/pkg/otel"
	"github.com/mehdihadeli/store-golang-microservice-sample/pkg/otel/tracing"
	"github.com/mehdihadeli/store-golang-microservice-sample/pkg/rabbitmq/config"
	"github.com/mehdihadeli/store-golang-microservice-sample/pkg/rabbitmq/consumer/configurations"
	"github.com/mehdihadeli/store-golang-microservice-sample/pkg/rabbitmq/producer"
	"github.com/mehdihadeli/store-golang-microservice-sample/pkg/rabbitmq/types"
	"github.com/mehdihadeli/store-golang-microservice-sample/pkg/test"
	"github.com/mehdihadeli/store-golang-microservice-sample/pkg/test/messaging/consumer"
	errorUtils "github.com/mehdihadeli/store-golang-microservice-sample/pkg/utils/error_utils"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Consume_Message(t *testing.T) {
	test.SkipCI(t)
	defer errorUtils.HandlePanic()

	ctx := context.Background()
	tp, err := tracing.AddOtelTracing(&otel.OpenTelemetryConfig{
		ServiceName:          "test",
		Enabled:              true,
		AlwaysOnSampler:      true,
		JaegerExporterConfig: &otel.JaegerExporterConfig{AgentHost: "localhost", AgentPort: "6831"},
		ZipkinExporterConfig: &otel.ZipkinExporterConfig{Url: "http://localhost:9411/api/v2/spans"},
	})
	if err != nil {
		return
	}
	defer tp.Shutdown(ctx)

	conn, err := types.NewRabbitMQConnection(context.Background(), &config.RabbitMQConfig{
		RabbitMqHostOptions: &config.RabbitMqHostOptions{
			UserName: "guest",
			Password: "guest",
			HostName: "localhost",
			Port:     5672,
		}})
	if err != nil {
		return
	}

	fakeHandler := consumer.NewRabbitMQFakeTestConsumerHandler()
	builder := configurations.NewRabbitMQConsumerConfigurationBuilder(ProducerConsumerMessage{})
	builder.WithHandlers(func(consumerHandlerBuilder messageConsumer.ConsumerHandlerConfigurationBuilder) {
		consumerHandlerBuilder.AddHandler(NewTestMessageHandler())
		consumerHandlerBuilder.AddHandler(fakeHandler)
	})

	rabbitmqConsumer, err := NewRabbitMQConsumer(conn, builder.Build(), json.NewJsonEventSerializer(), defaultLogger.Logger)

	if rabbitmqConsumer == nil {
		t.Log("RabbitMQ consumer is nil")
		return
	}
	err = rabbitmqConsumer.Start(ctx)
	if err != nil {
		rabbitmqConsumer.Stop(ctx)
	}

	rabbitmqProducer, err := producer.NewRabbitMQProducer(
		conn,
		nil,
		defaultLogger.Logger,
		json.NewJsonEventSerializer())
	if err != nil {
		t.Fatal(err)
	}
	//time.Sleep(time.Second * 5)
	//
	//fmt.Println("closing connection")
	//conn.Close()
	//fmt.Println(conn.IsClosed())
	//
	//time.Sleep(time.Second * 10)
	//fmt.Println("after 10 second of closing connection")
	//fmt.Println(conn.IsClosed())

	err = rabbitmqProducer.PublishMessage(ctx, NewProducerConsumerMessage("test"), nil)
	for err != nil {
		err = rabbitmqProducer.PublishMessage(ctx, NewProducerConsumerMessage("test"), nil)
	}

	err = test.WaitUntilConditionMet(func() bool {
		return fakeHandler.IsHandled()
	})
	assert.NoError(t, err)

	rabbitmqConsumer.Stop(ctx)
	conn.Close()

	fmt.Println(conn.IsClosed())
	fmt.Println(conn.IsConnected())
}

type ProducerConsumerMessage struct {
	*types2.Message
	Data string
}

func NewProducerConsumerMessage(data string) *ProducerConsumerMessage {
	return &ProducerConsumerMessage{
		Data:    data,
		Message: types2.NewMessage(uuid.NewV4().String()),
	}
}

// /////////// ConsumerHandlerT
type TestMessageHandler struct {
}

func (t *TestMessageHandler) Handle(ctx context.Context, consumeContext types2.MessageConsumeContext) error {
	message := consumeContext.Message().(*ProducerConsumerMessage)
	fmt.Println(message)

	return nil
}

func NewTestMessageHandler() *TestMessageHandler {
	return &TestMessageHandler{}
}

type TestMessageHandler2 struct {
}

func (t *TestMessageHandler2) Handle(ctx context.Context, consumeContext types2.MessageConsumeContext) error {
	message := consumeContext.Message()
	fmt.Println(message)

	return nil
}

func NewTestMessageHandler2() *TestMessageHandler2 {
	return &TestMessageHandler2{}
}

// /////////////// ConsumerPipeline
type Pipeline1 struct {
}

func NewPipeline1() pipeline.ConsumerPipeline {
	return &Pipeline1{}
}

func (p Pipeline1) Handle(ctx context.Context, consumerContext types2.MessageConsumeContext, next pipeline.ConsumerHandlerFunc) error {
	fmt.Println("PipelineBehaviourTest.Handled")

	fmt.Println(fmt.Sprintf("pipeline got a message with id '%s'", consumerContext.Message().GeMessageId()))

	err := next()
	if err != nil {
		return err
	}
	return nil
}
