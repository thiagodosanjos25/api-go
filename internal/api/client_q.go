package api

var (
	sqlAddClient = `
		INSERT INTO arrecada.circulares 
			(id_arrecadadora, id_sub_rede, id_estabelecimento, id_terminal, titulo,
			mensagem, data_inicio, data_fim, ativo, id_usuario)
		VALUES 
			(%d, (CASE WHEN (%d <> 0) THEN %d ELSE NULL END), (CASE WHEN (%d <> 0) THEN %d ELSE NULL END), (CASE WHEN (%d <> 0) THEN %d ELSE NULL END), '%s', 
			'%s', '%s', '%s', %t, %d) 
			RETURNING 
				id_circular;`

	sqlListClients = `
		SELECT 
			id_circular, 
			COALESCE(id_sub_rede, 0) as id_sub_rede,
			COALESCE(id_estabelecimento, 0) as id_estabelecimento,
			COALESCE(id_terminal, 0) as id_terminal,
			titulo,
			mensagem,
			data_inicio,
			data_fim,
			COALESCE(id_usuario, 0) as id_usuario,
			ativo
		FROM arrecada.circulares
		WHERE
			id_arrecadadora = %d AND
			data_inicio >= '%s' AND
			data_fim >= '%s' AND
			CASE WHEN ('%s' <> '') THEN titulo = '%s' ELSE TRUE END AND
			CASE WHEN (%d <> 0) THEN id_sub_rede = %d ELSE TRUE END AND
			CASE WHEN (%d <> 0) THEN id_estabelecimento = %d ELSE TRUE END AND
			CASE WHEN (%d <> 0) THEN id_terminal = %d ELSE TRUE END;`

	sqlGetClient = `
		SELECT 
			id_circular, 
			COALESCE(id_sub_rede, 0) as id_sub_rede,
			COALESCE(id_estabelecimento, 0) as id_estabelecimento,
			COALESCE(id_terminal, 0) as id_terminal,
			titulo,
			mensagem,
			data_inicio,
			data_fim,
			COALESCE(id_usuario, 0) as id_usuario,
			ativo
		FROM arrecada.circulares
		WHERE  
			id_arrecadadora = %d AND
			id_circular = %d;`

	sqlUpdateClient = `
		UPDATE 
			arrecada.circulares 
		SET 
			id_sub_rede = (CASE WHEN (%d <> 0) THEN %d ELSE NULL END),
			id_estabelecimento = (CASE WHEN (%d <> 0) THEN %d ELSE NULL END),
			id_terminal = (CASE WHEN (%d <> 0) THEN %d ELSE NULL END),
			titulo = '%s',
			mensagem = '%s',
			data_inicio = '%s',
			data_fim = '%s',
			id_usuario = %d,
			ativo = %t
		WHERE 
			id_arrecadadora = %d AND
			id_circular = %d
		RETURNING 
			id_circular;`

	sqlDeleteClient = `
		DELETE FROM 
			arrecada.circulares
		WHERE 
			id_arrecadadora = %d AND
			id_circular = %d
		RETURNING 
			id_circular, 
			COALESCE(id_sub_rede, 0) as id_sub_rede,
			COALESCE(id_estabelecimento, 0) as id_estabelecimento,
			COALESCE(id_terminal, 0) as id_terminal,
			titulo,
			mensagem,
			data_inicio,
			data_fim,
			COALESCE(id_usuario, 0) as id_usuario,
			ativo;`
)
