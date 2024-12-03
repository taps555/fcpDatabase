package repository

import (
	"a21hc3NpZ25tZW50/model"
	"fmt"

	"gorm.io/gorm"
)

type StudentRepository interface {
	FetchAll() ([]model.Student, error)
	FetchByID(id int) (*model.Student, error)
	Store(s *model.Student) error
	Update(id int, s *model.Student) error
	Delete(id int) error
	FetchWithClass() (*[]model.StudentClass, error)
}

type studentRepoImpl struct {
	db *gorm.DB
}

func NewStudentRepo(db *gorm.DB) *studentRepoImpl {
	return &studentRepoImpl{db}
}

func (s *studentRepoImpl) FetchAll() ([]model.Student, error) {
    var students []model.Student

    // Mengambil semua data dari tabel students
    err := s.db.Find(&students).Error
    if err != nil {
        // Jika terjadi error, kembalikan slice nil dan error
        return nil, fmt.Errorf("failed to fetch students: %w", err)
    }

    // Kembalikan hasil query beserta error nil
    return students, nil
}


func (s *studentRepoImpl) Store(student *model.Student) error {
    // Menyimpan data mahasiswa ke dalam tabel students
    err := s.db.Create(student).Error
    if err != nil {
        // Jika terjadi error, kembalikan error
        return fmt.Errorf("failed to store student: %w", err)
    }

    // Jika berhasil, kembalikan nil
    return nil
}


func (s *studentRepoImpl) Update(id int, student *model.Student) error {
    // Menjalankan query UPDATE berdasarkan ID
    err := s.db.Model(&model.Student{}).Where("id = ?", id).Updates(student).Error
    if err != nil {
        // Jika terjadi error, kembalikan error
        return fmt.Errorf("failed to update student with ID %d: %w", id, err)
    }

    // Jika berhasil, kembalikan nil
    return nil
}

func (s *studentRepoImpl) Delete(id int) error {
    // Menjalankan query DELETE untuk menghapus data mahasiswa berdasarkan ID
    err := s.db.Where("id = ?", id).Delete(&model.Student{}).Error
    if err != nil {
        // Jika terjadi error, kembalikan error
        return fmt.Errorf("failed to delete student with ID %d: %w", id, err)
    }

    // Jika berhasil, kembalikan nil
    return nil
}


func (s *studentRepoImpl) FetchByID(id int) (*model.Student, error) {
    var student model.Student

    // Query untuk mencari mahasiswa berdasarkan ID
    err := s.db.First(&student, id).Error
    if err != nil {
        // Jika error, kembalikan nil dan error
        if err == gorm.ErrRecordNotFound {
            return nil, fmt.Errorf("student with ID %d not found", id)
        }
        return nil, fmt.Errorf("failed to fetch student with ID %d: %w", id, err)
    }

    // Jika berhasil, kembalikan pointer ke student dan error nil
    return &student, nil
}

func (s *studentRepoImpl) FetchWithClass() (*[]model.StudentClass, error) {
    // Menyiapkan slice untuk menampung hasil query
    var studentClasses []model.StudentClass

    // Melakukan JOIN antara tabel 'students' dan 'classes'
    err := s.db.Table("students").
        Select("students.name, students.address, classes.name as class_name, classes.professor, classes.room_number").
        Joins("join classes on students.class_id = classes.id").
        Scan(&studentClasses).Error

    if err != nil {
        // Jika terjadi error, kembalikan nil dan error
        return nil, fmt.Errorf("failed to fetch students with class: %v", err)
    }

    // Jika tidak ada data yang ditemukan, kembalikan slice kosong
    if len(studentClasses) == 0 {
        return &[]model.StudentClass{}, nil
    }

    // Jika berhasil, kembalikan slice studentClasses dan nil sebagai error
    return &studentClasses, nil
}



