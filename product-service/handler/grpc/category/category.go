package category

import (
	context "context"

	"github.com/nafisalfiani/p3-final-project/lib/codes"
	"github.com/nafisalfiani/p3-final-project/lib/errors"
	"github.com/nafisalfiani/p3-final-project/lib/log"
	"github.com/nafisalfiani/p3-final-project/product-service/entity"
	"github.com/nafisalfiani/p3-final-project/product-service/usecase/category"
	"go.mongodb.org/mongo-driver/bson/primitive"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type grpcCat struct {
	log log.Interface
	cat category.Interface
}

func Init(log log.Interface, cat category.Interface) CategoryServiceServer {
	return &grpcCat{
		log: log,
		cat: cat,
	}
}

func (c *grpcCat) mustEmbedUnimplementedCategoryServiceServer() {}

func (c *grpcCat) GetCategory(ctx context.Context, in *Category) (*Category, error) {
	id, err := primitive.ObjectIDFromHex(in.GetId())
	if err != nil {
		return nil, errors.NewWithCode(codes.CodeInvalidValue, err.Error())
	}

	category, err := c.cat.Get(ctx, entity.Category{
		Id:   id,
		Name: in.GetName(),
	})
	if err != nil {
		return nil, err
	}

	res := &Category{
		Id:   category.Id.Hex(),
		Name: category.Name,
	}

	return res, nil
}

func (c *grpcCat) CreateCategory(ctx context.Context, in *Category) (*Category, error) {
	category, err := c.cat.Create(ctx, entity.Category{
		Name: in.GetName(),
	})
	if err != nil {
		return nil, err
	}

	res := &Category{
		Id:   category.Id.Hex(),
		Name: category.Name,
	}

	return res, nil
}

func (c *grpcCat) UpdateCategory(ctx context.Context, in *Category) (*Category, error) {
	id, err := primitive.ObjectIDFromHex(in.GetId())
	if err != nil {
		return nil, errors.NewWithCode(codes.CodeInvalidValue, err.Error())
	}

	category, err := c.cat.Update(ctx, entity.Category{
		Id:   id,
		Name: in.GetName(),
	})
	if err != nil {
		return nil, err
	}

	res := &Category{
		Id:   category.Id.Hex(),
		Name: category.Name,
	}

	return res, nil
}

func (c *grpcCat) DeleteCategory(ctx context.Context, in *Category) (*emptypb.Empty, error) {
	id, err := primitive.ObjectIDFromHex(in.GetId())
	if err != nil {
		return nil, errors.NewWithCode(codes.CodeInvalidValue, err.Error())
	}

	if err := c.cat.Delete(ctx, entity.Category{
		Id: id,
	}); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (c *grpcCat) GetCategories(ctx context.Context, in *emptypb.Empty) (*CategoryList, error) {
	categories, err := c.cat.List(ctx)
	if err != nil {
		return nil, err
	}

	res := &CategoryList{}
	for i := range categories {
		res.Categories = append(res.Categories, &Category{
			Id:   categories[i].Id.Hex(),
			Name: categories[i].Name,
		})
	}

	return res, nil
}
