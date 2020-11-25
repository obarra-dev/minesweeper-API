// Minesweeper-Service.
//
// The purpose of this API is to provide a Rest Api for game Minesweeper. It is developed in Golang.
//
//     Schemes: https
//     BasePath: /api
//     Version: 1.0.0
//     Host: minesweeper-service-ob.herokuapp.com
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta
package docs

import (
	"minesweeper-API/minesweeper-service/model"
)

// swagger:route GET /games/{id} games getGameEndpoint
// Play a movement.
// responses:
//   200: gameResponseResponseWrapper
//	 404: notfound Not Found
//   500: internal Internal Server Errors
// Game created OK.
// swagger:response gameResponseResponseWrapper
type GameResponseResponseWrapper struct {
	// in:body
	Body model.GameResponse
}

// Game attributes to set
// swagger:parameters getGameEndpoint
type GameRequestParamsWrapper struct {

	// The game id
	// required:true
	// in:path
	GameId int64 `json:"id"`
}

// swagger:route POST /games games createGameEndpoint
// Create a game.
// responses:
//   200: gameSimpleResponseResponse
//	 400: badrequest Missing or invalid attributes in body
//   500: internal Internal Server Errors
// Game created OK.
// swagger:response gameSimpleResponseResponse
type gameSimpleResponseResponseWrapper struct {
	// in:body
	Body model.GameSimpleResponse
}

// Game attributes to set
// swagger:parameters createGameEndpoint
type gameParamsWrapper struct {
	// The game's attributes to set.
	// in:body
	Body model.GameRequest
}

// swagger:route POST /games/{id}/play games playGameEndpoint
// Play a movement.
// responses:
//   200: playResponseResponseWrapper
//	 400: badrequest Missing or invalid attributes in body
//	 404: notfound Not Found
//   500: internal Internal Server Errors
// Game created OK.
// swagger:response playResponseResponseWrapper
type PlayResponseResponseWrapper struct {
	// in:body
	Body model.PlayResponse
}

// Play attributes to set
// swagger:parameters playGameEndpoint
type PlayRequestParamsWrapper struct {
	// The game's attributes to set.
	// in:body
	Body model.PlayRequest

	// The game id
	// required:true
	// in:path
	GameId int64 `json:"id"`
}

// swagger:route POST /games/{id}/mark games markGameEndpoint
// Marka a tile.
// responses:
//   200: ok
//   200: accepted
//	 400: badrequest Missing or invalid attributes in body
//	 404: notfound Not Found
//   500: internal Internal Server Errors
// Game created OK.

// Play attributes to set
// swagger:parameters markGameEndpoint
type MarkRequestParamsWrapper struct {
	// The game's attributes to set.
	// in:body
	Body model.MarkRequest

	// The game id
	// required:true
	// in:path
	GameId int64 `json:"id"`
}