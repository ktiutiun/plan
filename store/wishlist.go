package store

import "database/sql"

type Wish struct {
	Priority    string
	Wish        string
	Description string
	Link        string
}

func GetWishes(db *sql.DB) ([]Wish, error) {
	rows, err := db.Query("SELECT priority, wish, description, link FROM wishlist ORDER BY priority")
	if err != nil {
		return nil, err
	}

	wishes := []Wish{}

	for rows.Next() {
		wish := Wish{}

		err := rows.Scan(&wish.Priority, &wish.Wish, &wish.Description, &wish.Link)
		if err != nil {
			return nil, err
		}

		wishes = append(wishes, wish)
	}

	return wishes, nil
}
