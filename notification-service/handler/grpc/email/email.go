package email

import (
	context "context"

	"github.com/nafisalfiani/p3-final-project/lib/log"
	"github.com/nafisalfiani/p3-final-project/notification-service/usecase/mailer"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type grpcMail struct {
	log    log.Interface
	mailer mailer.Interface
}

func Init(log log.Interface, mailer mailer.Interface) EmailServiceServer {
	return &grpcMail{
		log:    log,
		mailer: mailer,
	}
}

func (r *grpcMail) mustEmbedUnimplementedEmailServiceServer() {}

func (r *grpcMail) SendRegistrationMail(ctx context.Context, in *Email) (*emptypb.Empty, error) {
	err := r.mailer.Send(ctx, fromProto(in))
	if err != nil {
		return &emptypb.Empty{}, err
	}

	return &emptypb.Empty{}, nil
}

func (r *grpcMail) SendInvoiceMail(ctx context.Context, in *Email) (*emptypb.Empty, error) {
	err := r.mailer.Send(ctx, fromProto(in))
	if err != nil {
		return &emptypb.Empty{}, err
	}

	return &emptypb.Empty{}, nil
}
