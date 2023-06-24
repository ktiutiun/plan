package store

import "database/sql"

type Wish struct {
	ID          int64  `json:"id"`
	Priority    string `json:"priority"`
	Wish        string `json:"wish"`
	Description string `json:"description"`
	Link        string `json:"link"`
}

func GetWishes(db *sql.DB) ([]Wish, error) {
	rows, err := db.Query("SELECT id, priority, wish, description, link FROM wishlist ORDER BY priority")
	if err != nil {
		return nil, err
	}

	wishes := []Wish{}

	for rows.Next() {
		wish := Wish{}

		err := rows.Scan(&wish.ID, &wish.Priority, &wish.Wish, &wish.Description, &wish.Link)
		if err != nil {
			return nil, err
		}

		wishes = append(wishes, wish)
	}

	return wishes, nil
}
