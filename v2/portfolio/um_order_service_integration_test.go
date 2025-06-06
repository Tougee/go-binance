//go:build integration
// +build integration

package portfolio

import (
	"context"
	"testing"
)

type umOrderServiceIntegrationTestSuite struct {
	*baseIntegrationTestSuite
}

func TestUMOrderServiceIntegration(t *testing.T) {
	base := SetupTest(t)
	suite := &umOrderServiceIntegrationTestSuite{
		baseIntegrationTestSuite: base,
	}

	t.Run("PlaceOrder", func(t *testing.T) {
		service := &UMOrderService{c: suite.client}
		order, err := service.Symbol("BTCUSDT").
			Side(SideTypeBuy).
			Type(OrderTypeLimit).
			TimeInForce(TimeInForceTypeGTC).
			Quantity("0.001").
			Price("20000").
			Do(context.Background())
		if err != nil {
			t.Fatalf("Failed to place order: %v", err)
		}

		// Basic validation of returned data
		if order.Symbol != "BTCUSDT" {
			t.Error("Expected symbol to be BTCUSDT")
		}

		if order.Side != SideTypeBuy {
			t.Error("Expected side to be BUY")
		}

		if order.Type != OrderTypeLimit {
			t.Error("Expected type to be LIMIT")
		}

		if order.TimeInForce != TimeInForceTypeGTC {
			t.Error("Expected timeInForce to be GTC")
		}

		if order.Status == "" {
			t.Error("Expected non-empty status")
		}

		if order.OrderID == 0 {
			t.Error("Expected non-zero order ID")
		}
	})
}
