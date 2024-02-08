package email

import (
	"github.com/nafisalfiani/p3-final-project/notification-service/entity"
)

func fromProto(in *Email) entity.Email {
	return entity.Email{
		Body:        in.Body,
		BodyType:    in.BodyType,
		Subject:     in.Subject,
		SenderName:  in.SenderName,
		SenderEmail: in.SenderEmail,
		Recipients: entity.Recipient{
			ToEmails:  in.RecipientTo,
			CCEmails:  in.RecipientCc,
			BCCEmails: in.RecipientBcc,
		},
		Attachments: in.Attachments,
	}
}

func fromUserProto(in *User) entity.User {
	return entity.User{
		Id:        in.GetId(),
		Name:      in.GetName(),
		Username:  in.GetUsername(),
		Email:     in.GetEmail(),
		Password:  in.GetPassword(),
		CreatedAt: in.GetCreatedAt().AsTime(),
		CreatedBy: in.GetCreatedBy(),
		UpdatedAt: in.GetUpdatedAt().AsTime(),
		UpdatedBy: in.GetUpdatedBy(),
	}
}
