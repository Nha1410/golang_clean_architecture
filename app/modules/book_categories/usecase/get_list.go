package usecase

import "github.com/team2/real_api/app/models"

func (u BookCategoryUseCase) GetList() ([]*models.BookCategoryResponse, error) {
	bookCategories, err := u.bookCategoryRepo.GetList()
	if err != nil {
		return nil, err
	}

	var bookCategoryResponses []*models.BookCategoryResponse
	for _, category := range bookCategories {
		bookCategoryResponses = append(bookCategoryResponses, models.FilterBookCategoryRecord(category))
	}


	return bookCategoryResponses, nil
}