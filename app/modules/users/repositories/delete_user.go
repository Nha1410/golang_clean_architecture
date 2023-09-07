package repository

import (
	"fmt"

	"github.com/team2/real_api/app/models"
)

func(r UserRepo) DeleteUser(user *models.User) error { 
	booksAssociation := r.DB.Model(&user).Association("Books")
	books := []*models.Book{}

	if err := booksAssociation.Find(&books); err != nil {
		// Xử lý lỗi
		fmt.Println("Lỗi khi truy vấn danh sách sách:", err)
		return nil
	}

	fmt.Println("Danh sách sách của người dùng:")
	for _, book := range books {
		fmt.Printf("ID: %d, Tên: %s\n", book.ID, book.Name)
	}

	if err := r.DB.Model(&user).Association("Books").Delete(user.Books); err != nil {
		return err
	}

	if err := r.DB.Model(&user).Association("BookCategories").Delete(user.BookCategories); err != nil {
		return err
	}

	if err := r.DB.Delete(&user).Error; err != nil {
		return err
	}

	return nil

}