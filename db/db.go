package db

import "context"

type Database interface {
	Connect(context.Context) error

	Disconnect(context.Context) error
}
