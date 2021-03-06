// Minesweeper-Service.
//
// The purpose of this API is to provide a Rest Api for game Minesweeper. It is developed in Golang.
//
//     Schemes: https
//     BasePath: /v1
//     Version: 1.0.0
//     Host: minesweeper-engine-ob.herokuapp.com
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta
package docs

// ------------------------------------------------------------
// General Responses
// ------------------------------------------------------------

// Success response
// swagger:response ok
type swaggerSuccessResponse struct {
	// in:body
	Body struct {
	}
}

// accepted response
// swagger:response accepted
type swaggerAcceptedResponse struct {
	// in:body
	Body struct {
	}
}

// Error BadRequest 400 - Missing or invalid attributes
// swagger:response badrequest
type swaggerErrorBadRequest struct {
	// in:body
	Body struct {
	}
}

// Error Unauthorized 401 - Missing Authorization token
// swagger:response unauthorized
type swaggerErrorUnauthorized struct {
	// in:body
	Body struct {
	}
}

// Error Forbidden 403
// swagger:response forbidden
type swaggerErrorForbidden struct {
	// in:body
	Body struct {
	}
}

// Error NotFaund 404
// swagger:response notfound
type swaggerErrorNotfound struct {
	// in:body
	Body struct {
	}
}

// Error Internal Server Error 500
// swagger:response internal
type swaggerErrorInternalServerError struct {
	// in:body
	Body struct {
	}
}
