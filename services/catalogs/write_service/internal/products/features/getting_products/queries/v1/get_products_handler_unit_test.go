package v1

import (
	"net/http"
	"testing"

	"emperror.dev/errors"

	customErrors "github.com/mehdihadeli/store-golang-microservice-sample/pkg/http/http_errors/custom_errors"

	"github.com/mehdihadeli/store-golang-microservice-sample/pkg/utils"
	"github.com/mehdihadeli/store-golang-microservice-sample/services/catalogs/write_service/internal/products/models"

	"github.com/mehdihadeli/store-golang-microservice-sample/services/catalogs/write_service/internal/products/mocks/testData"
	"github.com/mehdihadeli/store-golang-microservice-sample/services/catalogs/write_service/internal/shared/test_fixtures/unit_test"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type getProductsHandlerUnitTests struct {
	*unit_test.UnitTestSharedFixture
	*unit_test.UnitTestMockFixture
	getProductsHandler *GetProductsHandler
}

func TestGetProductsUnit(t *testing.T) {
	suite.Run(t, &getProductsHandlerUnitTests{UnitTestSharedFixture: unit_test.NewUnitTestSharedFixture(t)})
}

func (c *getProductsHandlerUnitTests) SetupTest() {
	// create new mocks or clear mocks before executing
	c.UnitTestMockFixture = unit_test.NewUnitTestMockFixture(c.T())
	c.getProductsHandler = NewGetProductsHandler(c.Log, c.Cfg, c.ProductRepository)
}

func (c *getProductsHandlerUnitTests) Test_Handle_Should_Return_Products_Successfully() {
	query := NewGetProducts(utils.NewListQuery(10, 1))

	items := utils.NewListResult[*models.Product](testData.Products, 10, 1, int64(len(testData.Products)))
	c.ProductRepository.On("GetAllProducts", mock.Anything, mock.Anything).
		Once().
		Return(items, nil)

	res, err := c.getProductsHandler.Handle(c.Ctx, query)
	c.Require().NoError(err)
	c.NotNil(res)
	c.NotEmpty(res.Products)
	c.Equal(len(testData.Products), len(res.Products.Items))
	c.ProductRepository.AssertNumberOfCalls(c.T(), "GetAllProducts", 1)
}

func (c *getProductsHandlerUnitTests) Test_Handle_Should_Return_Error_For_Error_In_Fetching_Data_From_Repository() {
	query := NewGetProducts(utils.NewListQuery(10, 1))

	c.ProductRepository.On("GetAllProducts", mock.Anything, mock.Anything).
		Once().
		Return(nil, errors.New("error in fetching products from repository"))

	res, err := c.getProductsHandler.Handle(c.Ctx, query)
	c.Require().Error(err)
	c.True(customErrors.IsApplicationError(err, http.StatusInternalServerError))
	c.Nil(res)
	c.ProductRepository.AssertNumberOfCalls(c.T(), "GetAllProducts", 1)
}

func (c *getProductsHandlerUnitTests) Test_Handle_Should_Return_Error_For_Mapping_List_Result() {
	query := NewGetProducts(utils.NewListQuery(10, 1))

	c.ProductRepository.On("GetAllProducts", mock.Anything, mock.Anything).
		Once().
		Return(nil, nil)

	res, err := c.getProductsHandler.Handle(c.Ctx, query)
	c.Require().Error(err)
	c.True(customErrors.IsApplicationError(err, http.StatusInternalServerError))
	c.Nil(res)
	c.ProductRepository.AssertNumberOfCalls(c.T(), "GetAllProducts", 1)
}
