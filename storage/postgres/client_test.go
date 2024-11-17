package postgres

import (
	"database/sql"
	"log/slog"
	"tender/model"
	"tender/pkg/logs"
	"testing"
)

func ConnectDB(logger *slog.Logger) *sql.DB {
	connector := "host = localhost user = postgres port = 5432 dbname = tender_management password = 03212164 sslmode = disable"
	db, err := sql.Open("postgres", connector)
	if err != nil {
		logger.Error(err.Error())
		return nil
	}
	return db
}

var Logger = logs.InitLogger()

func Test_CreateTender(t *testing.T) {
	db := ConnectDB(Logger)
	c := NewClientRepo(db, Logger)

	_, err := c.CreateTender(&model.CreateTenderReq{
		ClientId:    "f76b4c4b-da00-4957-bb88-b197e0ce739a",
		Title:       "Uy qurilishi",
		Description: "9 qavatli uy qurish kerak",
		Diadline:    "10-11-2025",
		Budget:      50_000_000.0,
	})
	if err != nil {
		Logger.Error(err.Error())
		t.Fatal(err)
	}
}

func Test_UpdateTender(t *testing.T) {
	db := ConnectDB(Logger)
	c := NewClientRepo(db, Logger)

	_, err := c.UpdateTender(&model.UpdateTenderReq{
		Id:          "97c71420-b902-434d-9ba0-e2d718451ead",
		ClientId:    "f76b4c4b-da00-4957-bb88-b197e0ce739a",
		Title:       "Uy qurilishi",
		Description: "9 qavatli uy qurish kerak",
		Diadline:    "10-11-2025",
		Budget:      50_000_000.0,
		Status:      "awarded",
	})

	if err != nil {
		Logger.Error(err.Error())
		t.Fatal(err)
	}
}

func Test_GetAllTenders(t *testing.T) {
	db := ConnectDB(Logger)
	c := NewClientRepo(db, Logger)

	_, err := c.GetAllTenders(&model.GetAllTendersReq{
		ClientId: "f76b4c4b-da00-4957-bb88-b197e0ce739a",
		Limit:    10,
		Page:     1,
	})

	if err != nil {
		Logger.Error(err.Error())
		t.Fatal(err)
	}
}

func Test_DeleteTender(t *testing.T) {
	db := ConnectDB(Logger)
	c := NewClientRepo(db, Logger)

	_, err := c.DeleteTender(&model.DeleteTenderReq{
		Id:       "708b0159-3763-4002-846f-c7e1d017a009",
		ClientId: "7308f557-bca4-4bd3-8dc5-67e0c0de6955",
	})

	if err != nil {
		Logger.Error(err.Error())
		t.Fatal(err)
	}
}

func Test_GetTenderBids(t *testing.T) {
	db := ConnectDB(Logger)
	c := NewClientRepo(db, Logger)

	_, err := c.GetTenderBids(&model.GetTenderBidsReq{
		ClientId:   "7308f557-bca4-4bd3-8dc5-67e0c0de6955",
		TenderId:   "b572c2d2-265c-43af-9fa5-671dcfa5f6c9",
		StartPrice: 10_000_000.0,
		EndPrice:   100_000_000.0,
		StartDate:  "01-01-2025",
		EndDate:    "01-01-2026",
		Limit:      10,
		Page:       1,
	})

	if err != nil {
		Logger.Error(err.Error())
		t.Fatal(err)
	}
}

func Test_SubmitBit(t *testing.T) {
	db := ConnectDB(Logger)
	c := NewClientRepo(db, Logger)

	_, err := c.SubmitBit(&model.SubmitBitReq{
		ClientId: "7308f557-bca4-4bd3-8dc5-67e0c0de6955",
		TenderId: "b572c2d2-265c-43af-9fa5-671dcfa5f6c9",
		BidId:    "",
		Status:   "Awarded",
	})

	if err != nil {
		Logger.Error(err.Error())
		t.Fatal(err)
	}
}

func Test_AwardTender(t *testing.T) {
	db := ConnectDB(Logger)
	c := NewClientRepo(db, Logger)

	_, err := c.AwardTender(&model.AwardTenderReq{
		ClientId: "7308f557-bca4-4bd3-8dc5-67e0c0de6955",
		TenderId: "b572c2d2-265c-43af-9fa5-671dcfa5f6c9",
		Bidid:    "",
	})

	if err != nil {
		t.Fatal(err)
	}
}
