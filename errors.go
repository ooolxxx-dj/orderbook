package orderbook

import "errors"

// List of errors from the orderbook package to the client
var (
	ErrInvalidQuantity      = errors.New("orderbook: invalid order quantity")
	ErrInvalidPrice         = errors.New("orderbook: invalid order price")
	ErrOrderExists          = errors.New("orderbook: order already exists")
	ErrOrderNotExists       = errors.New("orderbook: order does not exist")
	ErrInsufficientQuantity = errors.New("orderbook: insufficient quantity to calculate price")
)
