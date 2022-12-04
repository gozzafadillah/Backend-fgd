package follow_threads

import (
	"charum/business/comments"
	"charum/business/threads"
	"charum/business/users"
	"charum/dto"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FollowThreadUseCase struct {
	followThreadRepository Repository
	userRepository         users.Repository
	threadRepository       threads.Repository
	commentRepository      comments.Repository
	threadUseCase          threads.UseCase
}

func NewFollowThreadUseCase(ftr Repository, ur users.Repository, tr threads.Repository, cr comments.Repository, tuc threads.UseCase) UseCase {
	return &FollowThreadUseCase{
		followThreadRepository: ftr,
		userRepository:         ur,
		threadRepository:       tr,
		commentRepository:      cr,
		threadUseCase:          tuc,
	}
}

/*
Create
*/

func (ftu *FollowThreadUseCase) Create(domain *Domain) (Domain, error) {
	_, err := ftu.followThreadRepository.GetByUserIDAndThreadID(domain.UserID, domain.ThreadID)
	if err == nil {
		return Domain{}, errors.New("user already follow this thread")
	}

	_, err = ftu.userRepository.GetByID(domain.UserID)
	if err != nil {
		return Domain{}, errors.New("failed to get user")
	}

	_, err = ftu.threadRepository.GetByID(domain.ThreadID)
	if err != nil {
		return Domain{}, errors.New("failed to get thread")
	}

	domain.Id = primitive.NewObjectID()
	domain.Notification = 0
	domain.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
	domain.UpdatedAt = primitive.NewDateTimeFromTime(time.Now())

	result, err := ftu.followThreadRepository.Create(domain)
	if err != nil {
		return Domain{}, err
	}

	return result, nil
}

/*
Read
*/

func (ftu *FollowThreadUseCase) GetAllByUserID(userID primitive.ObjectID) ([]Domain, error) {
	result, err := ftu.followThreadRepository.GetAllByUserID(userID)
	if err != nil {
		return []Domain{}, errors.New("failed to get follow threads")
	}

	return result, nil
}

func (ftu *FollowThreadUseCase) CountByThreadID(threadID primitive.ObjectID) (int, error) {
	result, err := ftu.followThreadRepository.CountByThreadID(threadID)
	if err != nil {
		return 0, err
	}

	return result, nil
}

func (ftu *FollowThreadUseCase) DomainToResponse(domain Domain) (dto.ResponseFollowThread, error) {
	user, err := ftu.userRepository.GetByID(domain.UserID)
	if err != nil {
		return dto.ResponseFollowThread{}, errors.New("failed to get user")
	}

	thread, err := ftu.threadRepository.GetByID(domain.ThreadID)
	if err != nil {
		return dto.ResponseFollowThread{}, errors.New("failed to get thread")
	}

	responseThread, err := ftu.threadUseCase.DomainToResponse(thread)
	if err != nil {
		return dto.ResponseFollowThread{}, errors.New("failed to get response thread")
	}

	totalComment, err := ftu.commentRepository.CountByThreadID(domain.ThreadID)
	if err != nil {
		return dto.ResponseFollowThread{}, errors.New("failed to get total comment")
	}

	responseThread.TotalComment = totalComment

	response := dto.ResponseFollowThread{
		Id:           domain.Id,
		User:         user,
		Thread:       responseThread,
		Notification: domain.Notification,
		CreatedAt:    domain.CreatedAt.Time(),
		UpdatedAt:    domain.UpdatedAt.Time(),
	}

	return response, nil
}

func (ftu *FollowThreadUseCase) DomainToResponseArray(domains []Domain) ([]dto.ResponseFollowThread, error) {
	var responses []dto.ResponseFollowThread

	for _, domain := range domains {
		response, err := ftu.DomainToResponse(domain)
		if err != nil {
			return []dto.ResponseFollowThread{}, err
		}

		responses = append(responses, response)
	}

	return responses, nil
}

/*
Update
*/

func (ftu *FollowThreadUseCase) UpdateNotification(threadID primitive.ObjectID) error {
	err := ftu.followThreadRepository.AddOneNotification(threadID)
	if err != nil {
		return errors.New("failed to update notification")
	}

	return nil
}

func (ftu *FollowThreadUseCase) ResetNotification(threadID primitive.ObjectID, userID primitive.ObjectID) error {
	err := ftu.followThreadRepository.ResetNotification(threadID, userID)
	if err != nil {
		return errors.New("failed to reset notification")
	}

	return nil
}

/*
Delete
*/

func (ftu *FollowThreadUseCase) Delete(domain *Domain) (Domain, error) {
	_, err := ftu.userRepository.GetByID(domain.UserID)
	if err != nil {
		return Domain{}, errors.New("failed to get user")
	}

	_, err = ftu.threadRepository.GetByID(domain.ThreadID)
	if err != nil {
		return Domain{}, errors.New("failed to get thread")
	}

	result, err := ftu.followThreadRepository.GetByUserIDAndThreadID(domain.UserID, domain.ThreadID)
	if err != nil {
		return Domain{}, errors.New("failed to get follow thread")
	}

	err = ftu.followThreadRepository.Delete(result.Id)
	if err != nil {
		return Domain{}, errors.New("failed to unfollow thread")
	}

	return result, nil
}
