package service

import (
	pb "geekstudy.example/go/4week-engineering/mall/api/mall/v1"
	"geekstudy.example/go/4week-engineering/mall/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewProductService)

type ProductService struct {
	pb.UnimplementedProductServiceServer

	log *log.Helper

	product *biz.ProductUsecase
}
