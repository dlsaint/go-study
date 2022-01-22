package service

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel"

	pb "geekstudy.example/go/4week-engineering/mall/api/mall/v1"
	"geekstudy.example/go/4week-engineering/mall/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

func NewProductService(product *biz.ProductUsecase, logger log.Logger) *ProductService {
	return &ProductService{
		product: product,
		log:     log.NewHelper(logger),
	}
}

func (s *ProductService) ls(ctx context.Context, req *pb.CreateProductRequest) (*pb.CreateProductReply, error) {
	s.log.Infof("input data %v", req)
	err := s.product.Create(ctx, &biz.Product{
		Name:   req.Name,
		Price: req.Price,
	})
	return &pb.CreateProductReply{}, err
}

func (s *ProductService) UpdateProduct(ctx context.Context, req *pb.UpdateProductRequest) (*pb.UpdateProductReply, error) {
	s.log.Infof("input data %v", req)
	err := s.product.Update(ctx, req.Id, &biz.Product{
		Name:   req.Name,
		Price: req.Price,
	})
	return &pb.UpdateProductReply{}, err
}

func (s *ProductService) DeleteProduct(ctx context.Context, req *pb.DeleteProductRequest) (*pb.DeleteProductReply, error) {
	s.log.Infof("input data %v", req)
	err := s.product.Delete(ctx, req.Id)
	return &pb.DeleteProductReply{}, err
}

func (s *ProductService) GetProduct(ctx context.Context, req *pb.GetProductRequest) (*pb.GetProductReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetProduct")
	defer span.End()
	p, err := s.product.Get(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.GetProductReply{Product: &pb.Product{Id: p.ID, Name: p.Name, Price: p.Price, Like: p.Like}}, nil
}

func (s *ProductService) ListProduct(ctx context.Context, req *pb.ListProductRequest) (*pb.ListProductReply, error) {
	var a ProductService
	b := &ProductService{}
	fmt.Printf("%x", a)
	fmt.Printf("%x", b)

	ps, err := s.product.List(ctx)
	reply := &pb.ListProductReply{}
	for _, p := range ps {
		reply.Results = append(reply.Results, &pb.Product{
			Id:      p.ID,
			Name:   p.Name,
			Price: p.Price,
		})
	}
	return reply, err
}
