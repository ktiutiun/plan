package store

import "database/sql"

type Habit struct {
	Name  string
	Scale int64
}

func GetHealthHabits(db *sql.DB) ([]Habit, error) {
	rows, err := db.Query("SELECT habit_name, scale FROM health_habits ORDER BY habit_name")
	if err != nil {
		return nil, err
	}

	habits := []Habit{}

	for rows.Next() {
		habit := Habit{}

		err := rows.Scan(&habit.Name, &habit.Scale)
		if err != nil {
			return nil, err
		}

		habits = append(habits, habit)
	}

	return habits, nil
}
