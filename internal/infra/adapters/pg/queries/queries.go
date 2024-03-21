package queries

const InsertCategorization = `
		INSERT INTO categorization (number, category)
		VALUES ($1, $2)
		RETURNING id, number, category
	`

const GetByNumber = `
		SELECT id, number, category
		FROM categorization
		WHERE number = $1
	`
