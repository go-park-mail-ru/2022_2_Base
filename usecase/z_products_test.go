package usecase

import (
	"log"
	"serv/domain/model"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	orders "serv/microservices/orders/gen_files"
	mocks "serv/mocks"
)

func TestGetProducts(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	grcpConnOrders, err := grpc.Dial(
		":8083",
		//"localhost:8083",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Println("cant connect to grpc orders")
	} else {
		log.Println("connected to grpc orders service")
	}
	defer grcpConnOrders.Close()
	prodStoreMock := mocks.NewMockProductStoreInterface(ctrl)
	ordersManager := orders.NewOrdersWorkerClient(grcpConnOrders)
	prodUsecase := NewProductUsecase(prodStoreMock, &ordersManager)

	mockLastItemID := 0
	mockCount := 1
	mockSort := ""
	testProducts := new([3]*model.Product)
	err = faker.FakeData(testProducts)
	testProductsSlice := testProducts[:]

	assert.NoError(t, err)
	//mockArticleRepo.AssertExpectations(t)

	prodStoreMock.EXPECT().GetProductsFromStore(mockLastItemID, mockCount, mockSort).Return(testProductsSlice, nil)
	for _, testProd := range testProductsSlice {
		log.Println(testProd.Rating, *testProd.CommentsCount)
		prodStoreMock.EXPECT().GetProductsRatingAndCommsCountFromStore(testProd.ID).Return(testProd.Rating, *testProd.CommentsCount, nil)
	}
	//prodStoreMock.EXPECT().GetProductsRatingAndCommsCountFromStore(mockLastItemID, mockCount, mockSort).Return(0, 0, nil)
	products, err := prodUsecase.GetProducts(mockLastItemID, mockCount, mockSort)
	//resBID, err := boardUseCase.CreateBoard(testBoard)
	assert.NoError(t, err)
	assert.Equal(t, testProductsSlice, products)

	// mockArticleRepo := new(mocks.ArticleRepository)
	// mockArticle := domain.Article{
	// 	Title:   "Hello",
	// 	Content: "Content",
	// }

	// t.Run("success", func(t *testing.T) {
	// 	mockArticleRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(mockArticle, nil).Once()

	// 	mockArticleRepo.On("Delete", mock.Anything, mock.AnythingOfType("int64")).Return(nil).Once()

	// 	mockAuthorrepo := new(mocks.AuthorRepository)
	// 	u := ucase.NewArticleUsecase(mockArticleRepo, mockAuthorrepo, time.Second*2)

	// 	err := u.Delete(context.TODO(), mockArticle.ID)

	// 	assert.NoError(t, err)
	// 	mockArticleRepo.AssertExpectations(t)
	// 	mockAuthorrepo.AssertExpectations(t)
	// })
}
