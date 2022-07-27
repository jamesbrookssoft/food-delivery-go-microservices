package domain

import (
	httpErrors "github.com/mehdihadeli/store-golang-microservice-sample/pkg/http_errors"
)

var (
	ErrOrderAlreadyCompleted          = httpErrors.NewDomainError("Order already completed")
	ErrOrderAlreadyCanceled           = httpErrors.NewDomainError("Order is already canceled")
	ErrOrderMustBePaidBeforeDelivered = httpErrors.NewDomainError("Order must be paid before been delivered")
	ErrCancelReasonRequired           = httpErrors.NewDomainError("Cancel reason must be provided")
	ErrOrderAlreadyCancelled          = httpErrors.NewDomainError("order already cancelled")
	ErrAlreadyPaid                    = httpErrors.NewDomainError("already paid")
	ErrAlreadySubmitted               = httpErrors.NewDomainError("already submitted")
	ErrOrderNotPaid                   = httpErrors.NewDomainError("order not paid")
	ErrOrderNotFound                  = httpErrors.NewDomainError("order not found")
	ErrAlreadyCreated                 = httpErrors.NewDomainError("order with given id already created")
	ErrOrderShopItemsIsRequired       = httpErrors.NewDomainError("order shop items is required")
	ErrInvalidDeliveryAddress         = httpErrors.NewDomainError("Invalid delivery address")
	ErrInvalidDeliveryTimeStamp       = httpErrors.NewDomainError("Invalid delivery timestamp")
	ErrInvalidAccountEmail            = httpErrors.NewDomainError("Invalid account email")
	ErrInvalidOrderID                 = httpErrors.NewDomainError("Invalid order id")
	ErrInvalidTime                    = httpErrors.NewDomainError("Invalid time")
)
