package usecase

import (
	"context"
)

type SaveNewsLetterEmail interface {
	SaveEmail(ctx context.Context, email string) error
}

type UseCase struct {
	saveNewsLetterEmail SaveNewsLetterEmail
}

func NewUseCase(saveNewsLetterEmail SaveNewsLetterEmail) *UseCase {
	return &UseCase{saveNewsLetterEmail: saveNewsLetterEmail}
}

func (uc UseCase) Execute(ctx context.Context, email string) error {

	errSaveRecoveryToken := uc.saveNewsLetterEmail.SaveEmail(ctx, email)
	if errSaveRecoveryToken != nil {
		return errSaveRecoveryToken
	}

	return nil
}
