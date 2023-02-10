package middleware

import "context"

type ContextID int

const ID ContextID = 0

func GetID(ctx context.Context) string {
	if val, ok := ctx.Value(ID).(string); ok {
		return val
	}
	return ""
}
