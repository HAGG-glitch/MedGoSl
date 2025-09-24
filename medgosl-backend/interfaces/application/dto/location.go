package dto

import "time"

type LocationUpdate struct {
    OrderID  uint      `json:"order_id"`
    DriverID uint      `json:"driver_id"`
    Lat      float64   `json:"lat"`
    Lng      float64   `json:"lng"`
    At       time.Time `json:"at"`
}
