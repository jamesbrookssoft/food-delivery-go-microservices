package v1

import (
	"emperror.dev/errors"
	"fmt"
	customErrors "github.com/mehdihadeli/store-golang-microservice-sample/pkg/http/http_errors/custom_errors"
	"github.com/mehdihadeli/store-golang-microservice-sample/services/catalogs/write_service/internal/products/features/searching_product/queries/v1"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mehdihadeli/go-mediatr"
	"github.com/mehdihadeli/store-golang-microservice-sample/pkg/utils"
	"github.com/mehdihadeli/store-golang-microservice-sample/services/catalogs/write_service/internal/products/delivery"
	"github.com/mehdihadeli/store-golang-microservice-sample/services/catalogs/write_service/internal/products/features/searching_product/dtos"
)

type searchProductsEndpoint struct {
	*delivery.ProductEndpointBase
}

func NewSearchProductsEndpoint(productEndpointBase *delivery.ProductEndpointBase) *searchProductsEndpoint {
	return &searchProductsEndpoint{productEndpointBase}
}

func (ep *searchProductsEndpoint) MapRoute() {
	ep.ProductsGroup.GET("/search", ep.handler())
}

// SearchProducts
// @Tags Products
// @Summary Search products
// @Description Search products
// @Accept json
// @Produce json
// @Param searchProductsRequestDto query dtos.SearchProductsRequestDto false "SearchProductsRequestDto"
// @Success 200 {object} dtos.SearchProductsResponseDto
// @Router /api/v1/products/search [get]
func (ep *searchProductsEndpoint) handler() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		ep.CatalogsMetrics.SearchProductHttpRequests().Add(ctx, 1)

		listQuery, err := utils.GetListQueryFromCtx(c)

		if err != nil {
			badRequestErr := customErrors.NewBadRequestErrorWrap(err, "[searchProductsEndpoint_handler.GetListQueryFromCtx] error in getting data from query string")
			ep.Log().Errorf(fmt.Sprintf("[searchProductsEndpoint_handler.GetListQueryFromCtx] err: %v", badRequestErr))
			return err
		}

		request := &dtos.SearchProductsRequestDto{ListQuery: listQuery}
		if err := c.Bind(request); err != nil {
			badRequestErr := customErrors.NewBadRequestErrorWrap(err, "[searchProductsEndpoint_handler.Bind] error in the binding request")
			ep.Log().Errorf(fmt.Sprintf("[searchProductsEndpoint_handler.Bind] err: %v", badRequestErr))
			return badRequestErr
		}

		query := v1.NewSearchProducts(request.SearchText, request.ListQuery)

		if err := ep.Validator().StructCtx(ctx, query); err != nil {
			validationErr := customErrors.NewValidationErrorWrap(err, "[searchProductsEndpoint_handler.StructCtx]  query validation failed")
			ep.Log().Errorf("[searchProductsEndpoint_handler.StructCtx] err: {%v}", validationErr)
			return validationErr
		}

		queryResult, err := mediatr.Send[*v1.SearchProducts, *dtos.SearchProductsResponseDto](ctx, query)

		if err != nil {
			err = errors.WithMessage(err, "[searchProductsEndpoint_handler.Send] error in sending SearchProducts")
			ep.Log().Error(fmt.Sprintf("[searchProductsEndpoint_handler.Send] err: {%v}", err))
			return err
		}

		return c.JSON(http.StatusOK, queryResult)
	}
}
