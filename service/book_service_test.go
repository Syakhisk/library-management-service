package service_test

import (
	"database/sql"
	"errors"
	"library-management/mocks"
	"library-management/service"
	"library-management/shared/dto"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_bookService_GetBooks(t *testing.T) {
	mockRepo := mocks.NewBookRepository(t)
	bookService := service.NewBookService(mockRepo)

	tests := []struct {
		name    string
		mock    func()
		want    dto.GetBookResponseDTO
		wantErr func(err error)
	}{
		{
			name: "should get books",
			mock: func() {
				mockRepo.On("GetBooks").Return([]dto.Book{
					{
						ID:    1,
						Title: "Udin",
						Description: sql.NullString{
							String: "Petot",
							Valid:  true,
						},
					},
				}, nil).Once()
			},
			want: dto.GetBookResponseDTO{
				Books: []dto.BookResponse{
					{
						ID:          1,
						Title:       "Udin",
						Description: "Petot",
					},
				},
			},
			wantErr: func(err error) {
				assert.Nil(t, err)
			},
		},
		{
			name: "should return body error",
			mock: func() {
				mockRepo.On("GetBooks").Return([]dto.Book{}, errors.New("")).Once()
			},
			want: dto.GetBookResponseDTO{},
			wantErr: func(err error) {
				assert.NotNil(t, err)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Run mock
			tt.mock()

			data, err := bookService.GetBooks()

			assert.Equal(t, tt.want, data)
			tt.wantErr(err)
		})
	}
}
