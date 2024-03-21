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

const GetWithLimitAndOffset = `
		SELECT id, number, category
		FROM categorization
		ORDER BY id
		LIMIT $1 OFFSET $2
	`

const GetCount = `SELECT COUNT(*) FROM categorization`
