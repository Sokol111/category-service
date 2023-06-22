package handler

import (
	"context"
	"fmt"
	"github.com/Sokol111/category-service/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
)

type handler struct {
	proto.UnimplementedCategoryServiceServer
	categoryService ports.CategoryService
}

func NewCategoryServiceServer(s ports.CategoryService) proto.CategoryServiceServer {
	return &handler{categoryService: s}
}

func (h *handler) CreateCategory(_ context.Context, in *proto.CreateCategoryRequest) (*proto.CategoryResponse, error) {
	created, err := h.categoryService.Create(domain.Category{
		Name:    in.Name,
		Enabled: in.Enabled,
	})

	if err != nil {
		log.Printf("Failed to create category from grpc handler: %v\n", err)
		return nil, status.New(codes.Internal, "Failed to create category").Err()
	}
	return &proto.CategoryResponse{
		Id:               created.ID,
		Version:          created.Version,
		Name:             created.Name,
		Enabled:          created.Enabled,
		CreatedDate:      created.CreatedDate.String(),
		LastModifiedDate: created.LastModifiedDate.String(),
	}, nil
}

func (h *handler) GetCategoryById(_ context.Context, in *proto.GetCategoryByIdRequest) (*proto.CategoryResponse, error) {
	category, err := h.categoryService.GetById(in.Id)

	if err != nil {
		log.Printf("Failed to get category by id [%v] from grpc handler: %v\n", in.Id, err)
		return nil, status.New(codes.Internal, fmt.Sprintf("Failed to get category by id [%v]", in.Id)).Err()
	}

	return &proto.CategoryResponse{
		Id:               category.ID,
		Version:          category.Version,
		Name:             category.Name,
		Enabled:          category.Enabled,
		CreatedDate:      category.CreatedDate.String(),
		LastModifiedDate: category.LastModifiedDate.String(),
	}, nil
}

func (h *handler) GetCategoryByName(_ context.Context, in *proto.GetCategoryByNameRequest) (*proto.CategoryResponse, error) {
	category, err := h.categoryService.GetByName(in.Name)

	if err != nil {
		log.Printf("Failed to get category by name [%v] from grpc handler: %v\n", in.Name, err)
		return nil, status.New(codes.Internal, fmt.Sprintf("Failed to get category by name [%v]", in.Name)).Err()
	}

	return &proto.CategoryResponse{
		Id:               category.ID,
		Version:          category.Version,
		Name:             category.Name,
		Enabled:          category.Enabled,
		CreatedDate:      category.CreatedDate.String(),
		LastModifiedDate: category.LastModifiedDate.String(),
	}, nil
}

func (h *handler) UpdateCategory(_ context.Context, in *proto.UpdateCategoryRequest) (*proto.CategoryResponse, error) {
	updated, err := h.categoryService.Update(domain.Category{
		ID:      in.Id,
		Version: in.Version,
		Name:    in.Name,
		Enabled: in.Enabled,
	})

	if err != nil {
		log.Printf("Failed to update category with id [%v] from grpc handler: %v\n", in.Id, err)
		return nil, status.New(codes.Internal, fmt.Sprintf("Failed to update category with id [%v]", in.Id)).Err()
	}
	return &proto.CategoryResponse{
		Id:               updated.ID,
		Version:          updated.Version,
		Name:             updated.Name,
		Enabled:          updated.Enabled,
		CreatedDate:      updated.CreatedDate.String(),
		LastModifiedDate: updated.LastModifiedDate.String(),
	}, nil
}

func (h *handler) GetCategories(_ context.Context, _ *emptypb.Empty) (*proto.CategoryListResponse, error) {
	categories, err := h.categoryService.GetCategories()

	if err != nil {
		log.Printf("Failed to get categories from grpc handler: %v\n", err)
		return nil, status.New(codes.Internal, "Failed to get categories").Err()
	}

	responses := make([]*proto.CategoryResponse, len(categories))
	for i, c := range categories {
		responses[i] = &proto.CategoryResponse{
			Id:               c.ID,
			Version:          c.Version,
			Name:             c.Name,
			Enabled:          c.Enabled,
			CreatedDate:      c.CreatedDate.String(),
			LastModifiedDate: c.LastModifiedDate.String(),
		}
	}

	return &proto.CategoryListResponse{Categories: responses}, nil
}
