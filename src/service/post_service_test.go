package service_test

import (
	"net/http"
	"testing"

	"github.com/blog-service/src/domain/post"
	mock_repository "github.com/blog-service/src/mocks/repository"
	"github.com/blog-service/src/service"
	dateutils "github.com/blog-service/src/utils/date"
	stringutils "github.com/blog-service/src/utils/string"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestSave(t *testing.T) {
	ctrl := gomock.NewController(t)
	post := generateMockPostData()

	// 👇 create new mock post repository
	mockPostRepository := mock_repository.NewMockIPostRepository(ctrl)
	mockPostRepository.EXPECT().Save(gomock.Any()).Return(post, nil)

	service := service.NewTestPostService(mockPostRepository)
	returnedObject, err := service.Save(post)
	require.Nil(t, err)
	require.NotNil(t, returnedObject)
}

func TestSaveIfTitleIsInvalid(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockPostRepository := mock_repository.NewMockIPostRepository(ctrl)
	service := service.NewTestPostService(mockPostRepository)
	returnedObject, err := service.Save(&post.Post{Title: "", Description: "testksadhkashdkahkjhdkahskdhaskjdhakshdaskdhksatestksadhkashdkahkjhdkahskdhaskjdhakshdaskdhksatestksadhkashdkahkjhdkahskdhaskjdhakshdaskdhksa"})

	require.NotNil(t, err)
	require.Nil(t, returnedObject)
	require.Equal(t, http.StatusBadRequest, err.Code)
}

func TestFindById(t *testing.T) {
	ctrl := gomock.NewController(t)
	post := generateMockPostData()
	id := stringutils.GenerateUniqueId()

	// 👇 create new mock post repository
	mockPostRepository := mock_repository.NewMockIPostRepository(ctrl)
	mockPostRepository.EXPECT().FindById(id).Return(post, nil)

	service := service.NewTestPostService(mockPostRepository)
	returnedObject, err := service.FindById(id)
	require.Nil(t, err)
	require.NotNil(t, returnedObject)
}

func generateMockPostData() *post.Post {
	post := &post.Post{
		Id:          stringutils.GenerateUniqueId(),
		Title:       "testkshdkahskdhakshdka",
		Description: "testksadhkashdkahkjhdkahskdhaskjdhakshdaskdhksatestksadhkashdkahkjhdkahskdhaskjdhakshdaskdhksatestksadhkashdkahkjhdkahskdhaskjdhakshdaskdhksa",
		IsActive:    true,
		IsDeleted:   false,
		CreatedAt:   dateutils.GetNow().String(),
		UpdatedAt:   dateutils.GetNow().String(),
	}
	return post
}
