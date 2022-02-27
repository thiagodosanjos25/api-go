package api

var (
	sqlAddClient = `
		INSERT INTO gym_schma.clients 
			(name, gender, weight, height, imc, situation, created_at, updated_at, active)
		VALUES 
			('%s', '%s', %.2f, %.2f, %.4f, '%s', now(), now(), true) 
		RETURNING 
			id,
			name,
			gender,
			weight,
			height,
			imc,
			situation,
			created_at,
			updated_at,
			active;`

	sqlListClients = `
		SELECT 
			id,
			name,
			gender,
			weight,
			height,
			imc,
			situation,
			created_at,
			updated_at,
			active
		FROM gym_schma.clients
		WHERE
			CASE WHEN ('%s' <> '') THEN name ILIKE '%%%s%%' ELSE TRUE END AND
			CASE WHEN ('%s' <> '') THEN situation = '%s' ELSE TRUE END;`

	sqlGetClient = `
		SELECT 
			id,
			name,
			gender,
			weight,
			height,
			imc,
			situation,
			created_at,
			updated_at,
			active
		FROM gym_schma.clients
		WHERE  
			id = %d;`

	sqlUpdateClient = `
		UPDATE 
			gym_schma.clients 
		SET 
			name = '%s',
			gender = '%s',
			weight = %2.f,
			height = %2.f,
			imc = %2.f,
			situation = '%s',
			updated_at = now(),
			active = %t
		WHERE 
			id = %d
		RETURNING 
			id,
			name,
			gender,
			weight,
			height,
			imc,
			situation,
			created_at,
			updated_at,
			active;`

	sqlDeleteClient = `
		DELETE FROM 
			gym_schma.clients 
		WHERE 
			id = %d
		RETURNING 
			id,
			name,
			gender,
			weight,
			height,
			imc,
			situation,
			created_at,
			updated_at,
			active;`
)
