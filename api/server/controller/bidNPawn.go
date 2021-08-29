package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/uss-kelvin/NFTPawningShopBackend/server/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type BidNPawnController struct {
	bid  *model.Bids
	pawn *model.Pawns
}

func NewBidNPawnController(bid *model.Bids, pawn *model.Pawns) *BidNPawnController {
	return &BidNPawnController{
		bid:  bid,
		pawn: pawn,
	}
}

func (b *BidNPawnController) InsertBidToPawn(c *gin.Context, sc mongo.SessionContext) error {
	var newBid model.BidWrite
	err := c.Bind(&newBid)
	if err != nil {
		// log.Panic(err)
		return err
	}
	_, err = b.bid.InsertOne(sc, newBid)
	if err != nil {
		// log.Panic(err)
		return err
	}
	payload := model.PawnUpdate{
		Bid: newBid.ID,
	}
	err = b.pawn.UpdateOneBy(sc, "id", newBid.Pawn, payload)
	if err != nil {
		// log.Panic(err)
		return err
	}
	return nil
}