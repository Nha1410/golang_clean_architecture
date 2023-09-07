package repository

import (
	"github.com/team2/real_api/app/models"
)

func (r BookCategoryRepo) DeleteBookCategory(bookCategory *models.BookCategory) error {
    // Xóa tất cả các sách liên quan bằng cách sử dụng mối quan hệ
    if err := r.DB.Model(&bookCategory).Association("Books").Delete(bookCategory.Books); err != nil {
        return err
    }

    // Xóa danh mục sách
    if err := r.DB.Delete(&bookCategory).Error; err != nil {
        return err
    }

    return nil
}
