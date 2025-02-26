package domain

type Signal struct {
	ID        int     `json:"signal_id"`
	Name      string  `json:"signal_name" `
	ELR       string  `json:"elr"`
	Mileage   float64 `json:"mileage" `
	TrackID   int     `json:"track_id"`
	IsDeleted bool    `json:"is_deleted"`
}

type GetSignal struct {
	ID      *int
	TrackID *int
}

type UpdateSignal struct {
	ID        int
	Name      *string
	ELR       *string
	Mileage   *float64
	TrackID   *int
	IsDeleted *bool
}
