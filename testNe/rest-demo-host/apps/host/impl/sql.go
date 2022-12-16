package impl

const (
	InsertResourceSQL = `
	INSERT INTO resource (
		id,
		vendor,
		region,
		create_at,
		expire_at,
		type,
		name,
		description,
		status,
		update_at,
		sync_at,
		accout,
		public_ip,
		private_ip
	)
	VALUES
		(?,?,?,?,?,?,?,?,?,?,?,?,?,?);
	`

	// INSERT INTO `host` ( resource_id, cpu, memory, gpu_amount, gpu_spec, os_type, os_name, serial_number )
	// VALUES
	// ( "111", 1, 2048, 1, 'n', 'linux', 'centos8', '00000' );
	InsertDescribeSQL = `
	INSERT INTO host ( resource_id, cpu, memory, gpu_amount, gpu_spec, os_type, os_name, serial_number )
	VALUES
		( ?,?,?,?,?,?,?,? );
	`

	// 	--  1页2条数据, 2  offset: (page_number-1)*page_size
	// SELECT * FROM resource AS r LEFT JOIN `host` AS h ON r.id=h.resource_id  WHERE id LIKE 'ins-0%'  LIMIT 2,2 ;
	QueryHostSQL = `
	SELECT
	r.*,h.cpu, h.memory, h.gpu_spec, h.gpu_amount, h.os_type, h.os_name, h.serial_number
	FROM
		resource AS r
		LEFT JOIN host AS h ON r.id = h.resource_id
`

	updateResourceSQL = `UPDATE resource SET vendor=?,region=?,expire_at=?,name=?,description=? WHERE id = ?`
	updateHostSQL     = `UPDATE host SET cpu=?,memory=? WHERE resource_id = ?`
)
