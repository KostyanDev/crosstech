package domain

type Track struct {
	ID      int      `json:"track_id"`
	Source  string   `json:"source"`
	Target  string   `json:"target"`
	Signals []Signal `json:"signal_ids,omitempty"`
}

type GetTrack struct {
	ID     *int
	Source *string
	Target *string
}

type UpdateTrack struct {
	ID        int
	Source    *string
	Target    *string
	IsDeleted *bool
}
