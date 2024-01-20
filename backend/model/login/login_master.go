package login

type Model struct {
	PersonID int64 `json:"PersonID"`
}

func (m *Model) ResponseWithToken(token string) map[string]any {
	v := map[string]any{
		"token":    token,
		"personId": m.PersonID,
	}
	return v
}
