package ports

import (
	"sagebackend/internal/core/domain/repositories/rdbms"
	"sagebackend/internal/core/domain/services"

)

type  RdbmsEnquiriesRepository interface {
	CreateOne(
		customer string,
		terms_of_shipment string,
		origin_customs bool,
		door_pickup bool, 
		origin_port string,
		pickup_address string,
		cargo_ready_date  string,
		destination_customs  bool,
		door_delivery bool,
		destination_port  string,
		delivery_address  string,
		hs_code string,
		remarks string,
	) (int,error)

	SelectOne(id int) (rdbms.Enquiry, error)

	SelectAll() ([]rdbms.Enquiry, error)

	UpdateOne(
		id int,
		customer string,
		terms_of_shipment string,
		origin_customs bool,
		door_pickup bool,
		origin_port string,
		pickup_address string,
		cargo_ready_date string,
		destination_customs bool,
		door_delivery bool,
		destination_port string,
		delivery_address string,
		hs_code string,
		remarks string,
	) error

	DeleteOne(id int) error
}

type EnquiriesService interface {
	CreateEnquiry(
		customer string,
		terms_of_shipment string,
		origin_customs bool,
		door_pickup bool,
		origin_port string,
		pickup_address string,
		cargo_ready_date string,
		destination_customs bool,
		door_delivery bool,
		destination_port string,
		delivery_address string,
		hs_code string,
		remarks string,
		cargo []services.Cargo,
	) (error)

	GetEnquiryById(id int) (services.EnquiryResponse, error)

	GetAllEnquiries() ([]services.Enquiry, error)

	UpdateEnquiry(
		id int,
		customer string,
		terms_of_shipment string,
		origin_customs bool,
		door_pickup bool,
		origin_port string,
		pickup_address string,
		cargo_ready_date string,
		destination_customs bool,
		door_delivery bool,
		destination_port string,
		delivery_address string,
		hs_code string,
		remarks string,
		cargo []services.Cargo,
	) (error)

	DeleteEnquiry(id int) (error)

}

