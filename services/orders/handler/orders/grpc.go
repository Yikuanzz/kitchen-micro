package handler

import (
	"context"

	"github.com/yikuanzz/kitchen/services/common/genproto/orders"
	"github.com/yikuanzz/kitchen/services/orders/types"
	"google.golang.org/grpc"
)

type OrderGrpcHandler struct {
	// Service Injection
	orderService types.OrderService
	orders.UnimplementedOrderServiceServer
}

func NewGrpcOrderHandler(grpcServer *grpc.Server, orderService types.OrderService) *OrderGrpcHandler {
	gRPCHandler := &OrderGrpcHandler{
		orderService: orderService,
	}

	// Register the OrderServiceServer
	orders.RegisterOrderServiceServer(grpcServer, gRPCHandler)
	return gRPCHandler
}

func (h *OrderGrpcHandler) CreateOrder(ctx context.Context, req *orders.CreateOrderRequest) (*orders.CreateOrderResponse, error) {
	// Create the order
	order := &orders.Order{
		OrderID:    1,
		CustomerID: 2,
		ProductID:  2,
		Quantity:   10,
	}

	// Return the response
	err := h.orderService.CreateOrder(ctx, order)
	if err != nil {
		return nil, err
	}

	res := &orders.CreateOrderResponse{
		Status: "Order Created",
	}

	return res, nil
}

func (h *OrderGrpcHandler) GetOrders(ctx context.Context, req *orders.GetOrderRequest) (*orders.GetOrderResponse, error) {
	o := h.orderService.GetOrders(ctx)
	res := &orders.GetOrderResponse{
		Orders: o,
	}

	return res, nil
}
