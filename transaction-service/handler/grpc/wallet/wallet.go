package wallet

import (
	context "context"

	"github.com/nafisalfiani/p3-final-project/lib/log"
	"github.com/nafisalfiani/p3-final-project/transaction-service/usecase/wallet"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type grpcWallet struct {
	log    log.Interface
	wallet wallet.Interface
}

func Init(log log.Interface, wallet wallet.Interface) WalletServiceServer {
	return &grpcWallet{
		log:    log,
		wallet: wallet,
	}
}

func (w *grpcWallet) mustEmbedUnimplementedWalletServiceServer() {}

func (w *grpcWallet) GetWallet(ctx context.Context, in *Wallet) (*Wallet, error) {
	wallet, err := w.wallet.Get(ctx, fromProto(in))
	if err != nil {
		return nil, err
	}

	return toProto(wallet), nil
}

func (w *grpcWallet) CreateWallet(ctx context.Context, in *Wallet) (*Wallet, error) {
	wallet, err := w.wallet.Create(ctx, fromProto(in))
	if err != nil {
		return nil, err
	}

	return toProto(wallet), nil
}

func (w *grpcWallet) UpdateWallet(ctx context.Context, in *Wallet) (*Wallet, error) {
	wallet, err := w.wallet.Update(ctx, fromProto(in))
	if err != nil {
		return nil, err
	}

	return toProto(wallet), nil
}

func (w *grpcWallet) DeleteWallet(ctx context.Context, in *Wallet) (*emptypb.Empty, error) {
	if err := w.wallet.Delete(ctx, in.GetId()); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (w *grpcWallet) GetWallets(ctx context.Context, in *Wallet) (*WalletList, error) {
	wallets, err := w.wallet.List(ctx)
	if err != nil {
		return nil, err
	}

	res := &WalletList{}
	for i := range wallets {
		res.Wallets = append(res.Wallets, toProto(wallets[i]))
	}

	return res, nil
}
