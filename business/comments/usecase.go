package comments

import (
	"charum/business/threads"
	"charum/business/users"
	dtoComment "charum/dto/comments"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CommentUseCase struct {
	commentRepository Repository
	threadRepository  threads.Repository
	userRepository    users.Repository
}

func NewCommentUseCase(cr Repository, tr threads.Repository, ur users.Repository) UseCase {
	return &CommentUseCase{
		commentRepository: cr,
		threadRepository:  tr,
		userRepository:    ur,
	}
}

/*
Create
*/

func (cu *CommentUseCase) Create(domain *Domain) (Domain, error) {
	var err error
	if domain.ParentID != primitive.NilObjectID {
		_, err := cu.commentRepository.GetByIDAndThreadID(domain.ParentID, domain.ThreadID)
		if err != nil {
			return Domain{}, errors.New("failed to get parent comment")
		}
	}

	_, err = cu.threadRepository.GetByID(domain.ThreadID)
	if err != nil {
		return Domain{}, errors.New("failed to get thread")
	}

	domain.Id = primitive.NewObjectID()
	domain.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
	domain.UpdatedAt = primitive.NewDateTimeFromTime(time.Now())

	comment, err := cu.commentRepository.Create(domain)
	if err != nil {
		return Domain{}, errors.New("failed to create comment")
	}

	return comment, nil
}

/*
Read
*/

func (cu *CommentUseCase) GetByThreadID(threadID primitive.ObjectID) ([]Domain, error) {
	_, err := cu.threadRepository.GetByID(threadID)
	if err != nil {
		return []Domain{}, errors.New("failed to get thread")
	}

	comments, err := cu.commentRepository.GetByThreadID(threadID)
	if err != nil {
		return []Domain{}, errors.New("failed to get comments")
	}

	return comments, nil
}

func (cu *CommentUseCase) DomainToResponse(comment Domain) (dtoComment.Response, error) {
	responseComment := dtoComment.Response{}

	user, err := cu.userRepository.GetByID(comment.UserID)
	if err != nil {
		return dtoComment.Response{}, errors.New("failed to get user")
	}

	responseComment.Id = comment.Id
	responseComment.ThreadID = comment.ThreadID
	responseComment.ParentID = comment.ParentID
	responseComment.User = user
	responseComment.Comment = comment.Comment
	responseComment.CreatedAt = comment.CreatedAt
	responseComment.UpdatedAt = comment.UpdatedAt

	return responseComment, nil
}

func (cu *CommentUseCase) DomainToResponseArray(comments []Domain) ([]dtoComment.Response, error) {
	responseComments := []dtoComment.Response{}

	for _, comment := range comments {
		responseComment, err := cu.DomainToResponse(comment)
		if err != nil {
			return []dtoComment.Response{}, errors.New("failed to get response comment")
		}

		responseComments = append(responseComments, responseComment)
	}

	return responseComments, nil
}

func (cu *CommentUseCase) CountByThreadID(threadID primitive.ObjectID) (int, error) {
	_, err := cu.threadRepository.GetByID(threadID)
	if err != nil {
		return 0, errors.New("failed to get thread")
	}

	count, err := cu.commentRepository.CountByThreadID(threadID)
	if err != nil {
		return 0, errors.New("failed to count comments")
	}

	return count, nil
}

/*
Update
*/

func (cu *CommentUseCase) Update(domain *Domain) (Domain, error) {
	comment, err := cu.commentRepository.GetByIDAndThreadID(domain.Id, domain.ThreadID)
	if err != nil {
		return Domain{}, errors.New("failed to get comment")
	}

	if comment.UserID != domain.UserID {
		return Domain{}, errors.New("user are not the owner of this comment")
	}

	_, err = cu.threadRepository.GetByID(comment.ThreadID)
	if err != nil {
		return Domain{}, errors.New("failed to get thread")
	}

	comment.Comment = domain.Comment
	comment.UpdatedAt = primitive.NewDateTimeFromTime(time.Now())

	comment, err = cu.commentRepository.Update(&comment)
	if err != nil {
		return Domain{}, errors.New("failed to update comment")
	}

	return comment, nil
}

/*
Delete
*/

func (cu *CommentUseCase) Delete(id primitive.ObjectID, userID primitive.ObjectID) (Domain, error) {
	comment, err := cu.commentRepository.GetByID(id)
	if err != nil {
		return Domain{}, errors.New("failed to get comment")
	}

	if comment.UserID != userID {
		return Domain{}, errors.New("user are not the owner of this comment")
	}

	_, err = cu.threadRepository.GetByID(comment.ThreadID)
	if err != nil {
		return Domain{}, errors.New("failed to get thread")
	}

	err = cu.commentRepository.Delete(id)
	if err != nil {
		return Domain{}, errors.New("failed to delete comment")
	}

	return comment, nil
}

func (cu *CommentUseCase) DeleteAllByUserID(userID primitive.ObjectID) error {
	err := cu.commentRepository.DeleteAllByUserID(userID)
	if err != nil {
		return errors.New("failed to delete user's comments")
	}

	return nil
}

func (cu *CommentUseCase) DeleteAllByThreadID(threadID primitive.ObjectID) error {
	err := cu.commentRepository.DeleteAllByThreadID(threadID)
	if err != nil {
		return errors.New("failed to delete thread's comments")
	}

	return nil
}
