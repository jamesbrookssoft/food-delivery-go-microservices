package v1

import (
	"context"
	"emperror.dev/errors"
	"fmt"
	"github.com/mehdihadeli/go-mediatr"
	customErrors "github.com/mehdihadeli/store-golang-microservice-sample/pkg/http/http_errors/custom_errors"
	"github.com/mehdihadeli/store-golang-microservice-sample/pkg/logger"
	messageTracing "github.com/mehdihadeli/store-golang-microservice-sample/pkg/messaging/otel/tracing"
	types2 "github.com/mehdihadeli/store-golang-microservice-sample/pkg/messaging/types"
	"github.com/mehdihadeli/store-golang-microservice-sample/pkg/otel/tracing"
	"github.com/mehdihadeli/store-golang-microservice-sample/pkg/otel/tracing/attribute"
	deletingProductV1 "github.com/mehdihadeli/store-golang-microservice-sample/services/catalogs/read_service/internal/products/features/deleting_products/commands/v1"
	"github.com/mehdihadeli/store-golang-microservice-sample/services/catalogs/read_service/internal/shared/contracts"
	uuid "github.com/satori/go.uuid"
)

type productDeletedConsumer struct {
	contracts.InfrastructureConfiguration
}

func NewProductDeletedConsumer(infra contracts.InfrastructureConfiguration) *productDeletedConsumer {
	return &productDeletedConsumer{InfrastructureConfiguration: infra}
}

func (c *productDeletedConsumer) Handle(ctx context.Context, consumeContext types2.MessageConsumeContext) error {
	message, ok := consumeContext.Message().(*ProductDeletedV1)
	if !ok {
		return errors.New("error in casting message to ProductDeletedV1")
	}

	ctx, span := tracing.Tracer.Start(ctx, "productDeletedConsumer.Handle")
	span.SetAttributes(attribute.Object("Message", consumeContext.Message()))
	defer span.End()

	productUUID, err := uuid.FromString(message.ProductId)
	if err != nil {
		badRequestErr := customErrors.NewBadRequestErrorWrap(err, "[productDeletedConsumer_Handle.uuid.FromString] error in the converting uuid")
		c.GetLog().Errorf(fmt.Sprintf("[productDeletedConsumer_Handle.uuid.FromString] err: %v", messageTracing.TraceMessagingErrFromSpan(span, badRequestErr)))

		return err
	}

	command := deletingProductV1.NewDeleteProduct(productUUID)
	if err := c.GetValidator().StructCtx(ctx, command); err != nil {
		validationErr := customErrors.NewValidationErrorWrap(err, "[productDeletedConsumer_Handle.StructCtx] command validation failed")
		c.GetLog().Errorf(fmt.Sprintf("[productDeletedConsumer_Consume.StructCtx] err: {%v}", messageTracing.TraceMessagingErrFromSpan(span, validationErr)))

		return err
	}

	_, err = mediatr.Send[*deletingProductV1.DeleteProduct, *mediatr.Unit](ctx, command)

	if err != nil {
		err = errors.WithMessage(err, "[productDeletedConsumer_Handle.Send] error in sending DeleteProduct")
		c.GetLog().Errorw(fmt.Sprintf("[productDeletedConsumer_Handle.Send] id: {%s}, err: {%v}", command.ProductId, messageTracing.TraceMessagingErrFromSpan(span, err)), logger.Fields{"Id": command.ProductId})
	}

	return nil
}
