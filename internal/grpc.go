package internal

import (
	"context"
	product "github.com/monstrum/grpc-iris-demo/pkg/proto"
)

func CreateGrpcHandler() product.ApiServer {
	grpc := grpcHandler{}
	return &grpc
}

type grpcHandler struct {
}

func (s *grpcHandler) Create(
	_ context.Context,
	in *product.CreateProductReq,
) (*product.CreateProductResp, error) {
	return &product.CreateProductResp{
		Product: &product.Product{
			Id:          "test",
			Name:        in.Product.GetName(),
			Description: in.Product.GetDescription(),
		},
	}, nil
}

func (s *grpcHandler) Read(context.Context, *product.ReadProductReq) (*product.ReadProductResp, error) {
	return nil, nil
}

func (s *grpcHandler) Update(context.Context, *product.UpdateProductReq) (*product.UpdateProductResp, error) {
	return nil, nil
}

func (s *grpcHandler) Delete(context.Context, *product.DeleteProductReq) (*product.EmptyResp, error) {
	return nil, nil
}
