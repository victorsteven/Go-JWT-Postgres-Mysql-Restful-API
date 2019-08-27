package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/badoux/checkmail"
	"github.com/jinzhu/gorm"
	"github.com/victorsteven/fullstack/api/utils/channels"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Nickname  string    `gorm:"size:255;not null;unique" json:"nickname"`
	Email     string    `gorm:"size:100;not null;unique" json:"email"`
	Password  string    `gorm:"size:100;not null;" json:"password"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (u *User) BeforeSave() error {
	hashedPassword, err := Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) Prepare() {
	u.ID = 0
	u.Nickname = html.EscapeString(strings.TrimSpace(u.Nickname))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

func (u *User) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if u.Nickname == "" {
			return errors.New("Required Nickname")
		}
		// if u.Password == "" {
		// 	return errors.New("Required Password")
		// }
		if u.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}

		return nil
	case "login":
		if u.Password == "" {
			return errors.New("Required Password")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}
		return nil

	default:
		if u.Nickname == "" {
			return errors.New("Required Nickname")
		}
		if u.Password == "" {
			return errors.New("Required Password")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}
		return nil
	}
}

func (u *User) SaveUser(db *gorm.DB) (*User, error) {
	var err error
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = db.Debug().Create(&u).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)

	if channels.OK(done) {
		return u, nil
	}
	return &User{}, err
}

func (u *User) FindAllUsers(db *gorm.DB) (*[]User, error) {
	var err error
	users := []User{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = db.Debug().Model(&User{}).Limit(100).Find(&users).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)

	if channels.OK(done) {
		return &users, nil
	}
	return nil, err
}

func (u *User) FindUserByID(db *gorm.DB, uid uint32) (*User, error) {
	var err error
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = db.Debug().Model(User{}).Where("id = ?", uid).Take(&u).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)

	if channels.OK(done) {
		return u, nil
	}
	if gorm.IsRecordNotFoundError(err) {
		return &User{}, errors.New("User Not Found")
	}
	return &User{}, err
}

func (u *User) UpdateAUser(db *gorm.DB, uid uint32) (*User, error) {
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		db = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&User{}).UpdateColumns(
			map[string]interface{}{
				"nickname":  u.Nickname,
				"email":     u.Email,
				"update_at": time.Now(),
			},
		)
		ch <- true
	}(done)

	if channels.OK(done) {
		if db.Error != nil {
			return &User{}, db.Error
		}
		// This is the display the updated user
		err := db.Debug().Model(&User{}).Where("id = ?", uid).Take(&u).Error
		if err != nil {
			return &User{}, db.Error
		}
		return u, nil
	}
	return &User{}, db.Error
}

func (u *User) DeleteAUser(db *gorm.DB, uid uint32) (int64, error) {
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		db = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&User{}).Delete(&User{})
		ch <- true
	}(done)

	if channels.OK(done) {
		if db.Error != nil {
			return 0, db.Error
		}
		return db.RowsAffected, nil
	}
	return 0, db.Error
}
