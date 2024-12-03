package repository

import (
	"a21hc3NpZ25tZW50/model"
	"fmt"

	"gorm.io/gorm"
)

type ClassRepository interface {
	FetchAll() ([]model.Class, error)
}

type classRepoImpl struct {
	db *gorm.DB
}

func NewClassRepo(db *gorm.DB) *classRepoImpl {
	return &classRepoImpl{db}
}

func (s *classRepoImpl) FetchAll() ([]model.Class, error) {
    var classes []model.Class

    // Melakukan query untuk mengambil semua data dari tabel 'classes'
    err := s.db.Find(&classes).Error

    if err != nil {
        // Jika terjadi error, kembalikan slice kosong dan error yang terjadi
        return nil, fmt.Errorf("failed to fetch classes: %v", err)
    }

    // Jika berhasil, kembalikan slice kelas dan nil sebagai error
    return classes, nil
}

