package region

import (
	context "context"

	"github.com/nafisalfiani/p3-final-project/lib/codes"
	"github.com/nafisalfiani/p3-final-project/lib/errors"
	"github.com/nafisalfiani/p3-final-project/lib/log"
	"github.com/nafisalfiani/p3-final-project/product-service/entity"
	"github.com/nafisalfiani/p3-final-project/product-service/usecase/region"
	"go.mongodb.org/mongo-driver/bson/primitive"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type grpcReg struct {
	log log.Interface
	reg region.Interface
}

func Init(log log.Interface, reg region.Interface) RegionServiceServer {
	return &grpcReg{
		log: log,
		reg: reg,
	}
}

func (r *grpcReg) mustEmbedUnimplementedRegionServiceServer() {}

func (r *grpcReg) GetRegion(ctx context.Context, in *Region) (*Region, error) {
	id, err := primitive.ObjectIDFromHex(in.GetId())
	if err != nil {
		return nil, errors.NewWithCode(codes.CodeInvalidValue, err.Error())
	}

	region, err := r.reg.Get(ctx, entity.Region{
		Id:   id,
		Name: in.GetName(),
	})
	if err != nil {
		return nil, err
	}

	res := &Region{
		Id:   region.Id.Hex(),
		Name: region.Name,
	}

	return res, nil
}

func (r *grpcReg) CreateRegion(ctx context.Context, in *Region) (*Region, error) {
	region, err := r.reg.Create(ctx, entity.Region{
		Name: in.GetName(),
	})
	if err != nil {
		return nil, err
	}

	res := &Region{
		Id:   region.Id.Hex(),
		Name: region.Name,
	}

	return res, nil
}

func (r *grpcReg) UpdateRegion(ctx context.Context, in *Region) (*Region, error) {
	id, err := primitive.ObjectIDFromHex(in.GetId())
	if err != nil {
		return nil, errors.NewWithCode(codes.CodeInvalidValue, err.Error())
	}

	region, err := r.reg.Update(ctx, entity.Region{
		Id:   id,
		Name: in.GetName(),
	})
	if err != nil {
		return nil, err
	}

	res := &Region{
		Id:   region.Id.Hex(),
		Name: region.Name,
	}

	return res, nil
}

func (r *grpcReg) DeleteRegion(ctx context.Context, in *Region) (*emptypb.Empty, error) {
	id, err := primitive.ObjectIDFromHex(in.GetId())
	if err != nil {
		return nil, errors.NewWithCode(codes.CodeInvalidValue, err.Error())
	}

	if err := r.reg.Delete(ctx, entity.Region{
		Id: id,
	}); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (r *grpcReg) GetRegions(ctx context.Context, in *emptypb.Empty) (*RegionList, error) {
	regions, err := r.reg.List(ctx)
	if err != nil {
		return nil, err
	}

	res := &RegionList{}
	for i := range regions {
		res.Regions = append(res.Regions, &Region{
			Id:   regions[i].Id.Hex(),
			Name: regions[i].Name,
		})
	}

	return res, nil
}
