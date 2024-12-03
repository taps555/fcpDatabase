package repository

import (
	"a21hc3NpZ25tZW50/model"
	"fmt"

	"github.com/jackc/pgconn"
	"gorm.io/gorm"
)

type SessionsRepository interface {
	AddSessions(session model.Session) error
	DeleteSession(token string) error
	UpdateSessions(session model.Session) error
	SessionAvailName(name string) error
	SessionAvailToken(token string) (model.Session, error)
}

type sessionsRepoImpl struct {
	db *gorm.DB
}
// 
func NewSessionRepo(db *gorm.DB) *sessionsRepoImpl {
	return &sessionsRepoImpl{db}
}

func (s *sessionsRepoImpl) AddSessions(session model.Session) error {
    fmt.Println("Received session:", session)
    if result := s.db.Create(&session); result.Error != nil {
        return result.Error
    }
    return nil
}


func (s *sessionsRepoImpl) DeleteSession(token string) error {
	session := model.Session{}
	if result := s.db.Table("sessions").Where("token = ?", token).Delete(&session); result.Error != nil {
		return fmt.Errorf("error deleting teacher: %v", result.Error)
	}
	return nil // TODO: replace this
}

func (s *sessionsRepoImpl) UpdateSessions(session model.Session) error {
	if err := s.db.Table("sessions").
	Where("username = ?", session.Username).
	Update("token", session.Token).
	Error; err != nil {
		return err
	}

	
	
	return nil // TODO: replace this
}


func (s *sessionsRepoImpl) SessionAvailName(username string) error {
    var session model.Session

    // Cari session berdasarkan username
    err := s.db.Where("username = ?", username).First(&session).Error
    if err != nil {
        if err == gorm.ErrRecordNotFound {
            // Jika tidak ditemukan, kembalikan error yang lebih jelas
            return fmt.Errorf("session with username %s not found", username)
        }
        // Jika ada error lain, periksa apakah itu karena kolom 'username' tidak ditemukan
        if pqErr, ok := err.(*pgconn.PgError); ok && pqErr.Code == "42703" {
            return fmt.Errorf("column 'name' does not exist in the sessions table. Please check the database schema")
        }
        // Jika error lain, langsung kembalikan error tersebut
        return err
    }

    // Jika ditemukan, cetak data session (opsional)
    fmt.Println("Session Found:", session)
    return nil
}




func (s *sessionsRepoImpl) SessionAvailToken(token string) (model.Session, error) {
    var session model.Session

    // Cari session berdasarkan token
    if err := s.db.Where("token = ?", token).First(&session).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            // Jika tidak ditemukan, kembalikan session kosong dan pesan error
            return model.Session{}, fmt.Errorf("session with token %s not found", token)
        }
        // Jika ada error lain, langsung kembalikan error tersebut
        return model.Session{}, err
    }

    // Jika ditemukan, kembalikan session dan error nil
    return session, nil
}
