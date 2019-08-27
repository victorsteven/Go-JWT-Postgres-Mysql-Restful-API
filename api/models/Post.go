package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/victorsteven/fullstack/api/utils/channels"
)

type Post struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Title     string    `gorm:"size:255;not null;unique" json:"title"`
	Content   string    `gorm:"size:255;not null;" json:"content"`
	Author    User      `json:"author"`
	AuthorID  uint32    `gorm:"not null" json:"author_id"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	// CreatedAt time.Time `json:"created_at"`
	// UpdatedAt time.Time `json:"updated_at"`
}

func (p *Post) Prepare() {
	p.ID = 0
	p.Title = html.EscapeString(strings.TrimSpace(p.Title))
	p.Content = html.EscapeString(strings.TrimSpace(p.Content))
	p.Author = User{}
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
}

func (p *Post) Validate() error {

	if p.Title == "" {
		return errors.New("Required Title")
	}
	if p.Content == "" {
		return errors.New("Required Content")
	}
	if p.AuthorID < 1 {
		return errors.New("Required Author")
	}
	return nil
}

func (p *Post) SavePost(db *gorm.DB) (*Post, error) {
	var err error
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = db.Debug().Model(&Post{}).Create(&p).Error
		if err != nil {
			ch <- false
			return
		}
		if p.ID != 0 {
			err = db.Debug().Model(&User{}).Where("id = ?", p.AuthorID).Take(&p.Author).Error
			if err != nil {
				ch <- false
				return
			}
		}
		ch <- true
	}(done)

	if channels.OK(done) {
		return p, nil
	}
	return &Post{}, err
}

func (p *Post) FindAllPosts(db *gorm.DB) ([]Post, error) {
	var err error
	posts := []Post{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = db.Debug().Model(&Post{}).Limit(100).Find(&posts).Error
		if err != nil {
			ch <- false
			return
		}
		if len(posts) > 0 {
			for i, _ := range posts {
				err := db.Debug().Model(&User{}).Where("id = ?", posts[i].AuthorID).Take(&posts[i].Author).Error
				if err != nil {
					ch <- false
					return
				}
			}
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return posts, nil
	}
	return nil, err
}

func (p *Post) FindPostByID(db *gorm.DB, pid uint64) (*Post, error) {
	var err error
	done := make(chan bool)
	go func(ch chan<- bool) {
		// defer close(ch)
		err = db.Debug().Model(&Post{}).Where("id = ?", pid).Take(&p).Error
		if err != nil {
			ch <- false
			return
		}
		if p.ID != 0 {
			err = db.Debug().Model(&User{}).Where("id = ?", p.AuthorID).Take(&p.Author).Error
			if err != nil {
				ch <- false
				return
			}
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return p, nil
	}
	return &Post{}, err
}

func (p *Post) UpdateAPost(db *gorm.DB, pid uint64) (*Post, error) {

	var err error
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		db = db.Debug().Model(&Post{}).Where("id = ?", pid).Take(&Post{}).UpdateColumns(
			map[string]interface{}{
				"title":      p.Title,
				"content":    p.Content,
				"updated_at": time.Now(),
			},
		)
		err = db.Debug().Model(&Post{}).Where("id = ?", pid).Take(&p).Error
		if err != nil {
			ch <- false
			return
		}
		if p.ID != 0 {
			err = db.Debug().Model(&User{}).Where("id = ?", p.AuthorID).Take(&p.Author).Error
			if err != nil {
				ch <- false
				return
			}
		}
		ch <- true
	}(done)

	if channels.OK(done) {
		return p, nil
	}
	return &Post{}, err
}

func (p *Post) DeleteAPost(db *gorm.DB, pid uint64, uid uint32) (int64, error) {

	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		db = db.Debug().Model(&Post{}).Where("id = ? and author_id = ?", pid, uid).Take(&Post{}).Delete(&Post{})
		ch <- true
	}(done)
	if channels.OK(done) {
		if db.Error != nil {
			if gorm.IsRecordNotFoundError(db.Error) {
				return 0, errors.New("Post not found")
			}
			return 0, db.Error
		}
		return db.RowsAffected, nil
	}
	return 0, db.Error
}
