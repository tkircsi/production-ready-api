package models

import "gorm.io/gorm"

type Service struct {
	DB *gorm.DB
}

type Comment struct {
	gorm.Model
	Slug   string
	Body   string
	Author string
}

type CommentService interface {
	GetComment(ID uint) (*Comment, error)
	GetCommentBySlug(slug string) ([]*Comment, error)
	PostComment(c Comment) (*Comment, error)
	UpdateComment(ID uint, new Comment) (*Comment, error)
	DeleteComment(ID uint) error
	GetAllComments() ([]*Comment, error)
}

func NewService(db *gorm.DB) *Service {
	return &Service{
		DB: db,
	}
}

func (s *Service) GetComment(ID uint) (*Comment, error) {
	comment := new(Comment)
	if tx := s.DB.First(comment, ID); tx.Error != nil {
		return comment, tx.Error
	}
	return comment, nil
}

func (s *Service) GetCommentBySlug(slug string) ([]*Comment, error) {
	var comments []*Comment
	if tx := s.DB.Find(comments).Where("slug = ?", slug); tx.Error != nil {
		return comments, tx.Error
	}
	return comments, nil
}

func (s *Service) PostComment(c Comment) (*Comment, error) {
	if tx := s.DB.Save(&c); tx.Error != nil {
		return &c, tx.Error
	}
	return &c, nil
}

func (s *Service) UpdateComment(ID uint, newComment Comment) (*Comment, error) {
	comment, err := s.GetComment(ID)
	if err != nil {
		return nil, err
	}

	if tx := s.DB.Model(comment).Updates(newComment); tx.Error != nil {
		return nil, tx.Error
	}
	return comment, nil
}

func (s *Service) DeleteComment(ID uint) error {
	if tx := s.DB.Delete(&Comment{}, ID); tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (s *Service) GetAllComments() ([]*Comment, error) {
	var comments []*Comment

	if tx := s.DB.Find(&comments); tx.Error != nil {
		return comments, tx.Error
	}
	return comments, nil
}
