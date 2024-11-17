package postgres_test

import (
	"fmt"
	"tender/model"
	"tender/storage/postgres"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateBid(t *testing.T) {
	db := postgres.ConnectDB(postgres.Logger)
	Bidrepo := postgres.NewBidRepository(db)

	resp, err := Bidrepo.CreateBid(
		model.CreateBidInput{
			TenderID:     "2de1f098-a446-4f47-aa14-d2c8a4f985bc",
			ContractorID: "99f13ede-f029-4ba6-bc90-c28269f82fdc",
			Price:        5000,
			DeliveryTime: "07-10-2024",
			Comments:     "Sample comment",
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(resp)
}

func TestGetTendersByFilters(t *testing.T) {
	db := postgres.ConnectDB(postgres.Logger)
	Bidrepo := postgres.NewBidRepository(db)

	resp, err := Bidrepo.GetTendersByFilters(
		model.GetTendersInput{
			Status: "close",
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, []model.Tender(nil), resp)
	// fmt.Println(resp)
}

func TestGetBidsForTenderWithFilters(t *testing.T) {
	db := postgres.ConnectDB(postgres.Logger)
	Bidrepo := postgres.NewBidRepository(db)

	resp, err := Bidrepo.GetBidsForTenderWithFilters(
		model.GetBidsInput{
			TenderID: "ff80539c-3e53-4cd8-a62a-7e0bc2465640",
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	// assert.Equal(t, []model.Tender(nil), resp)
	fmt.Println(resp)
}

func TestGetMyBidHistory(t *testing.T) {
	db := postgres.ConnectDB(postgres.Logger)
	Bidrepo := postgres.NewBidRepository(db)

	resp, err := Bidrepo.GetMyBidHistory(
		model.GetMyBidsInput{
			UserID: "0b214112-8f65-434d-bb4c-b5f7be5337ab",
		},
	)
	if err != nil {
		t.Fatal(err)
	}

	// assert.Equal(t, []model.Tender(nil), resp)
	fmt.Println(resp)
}
