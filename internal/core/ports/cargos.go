package ports

import "sagebackend/internal/core/domain/repositories/rdbms"


type RdbmsCargosRepository interface {
	CreateOne(
		enquiry_id int,
		count int,
		length int,
		width int,
		height int,
		volume string,
		weight int,
		unit string,
		total_volume string,
		total_weight string,
	) error
	SelectAll(enquiry_id int) ([]rdbms.Cargo, error)	
	DeleteAll(enquiry_id int) error
}