package api
import (
	"encoding/json"
	"net/http"
)
type coinBalanceParams struct {
	Username string
}

type coinBalanceResponse struct {
	Code    int
	Balance int64
}

type Error struct {
	Code    int
	Message string
}
