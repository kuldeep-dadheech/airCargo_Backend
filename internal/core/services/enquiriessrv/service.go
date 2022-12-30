package enquiriessrv

import (
	"aircargo/internal/core/domain/repositories/rdbms"
	"aircargo/internal/core/domain/services"
	"aircargo/internal/core/ports"
)

type Service struct {
	enquiriesRepository ports.RdbmsEnquiriesRepository
	cargosRepository    ports.RdbmsCargosRepository
}

func New(
	enquiriesRepository ports.RdbmsEnquiriesRepository,
	cargosRepository ports.RdbmsCargosRepository,
) *Service {
	return &Service{
		enquiriesRepository: enquiriesRepository,
		cargosRepository:    cargosRepository,
	}
}

func (s *Service) CreateEnquiry(
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
) error {
	id, err := s.enquiriesRepository.CreateOne(
		customer,
		terms_of_shipment,
		origin_customs,
		door_pickup,
		origin_port,
		pickup_address,
		cargo_ready_date,
		destination_customs,
		door_delivery,
		destination_port,
		delivery_address,
		hs_code,
		remarks,
	)
	if err != nil {
		return err
	}
	for _, v := range cargo {
		errf := s.cargosRepository.CreateOne(
			id,
			v.Count,
			v.Length,
			v.Width,
			v.Height,
			v.Volume,
			v.Weight,
			v.Unit,
			v.Total_volume,
			v.Total_weight,
		)
		if errf != nil {
			return errf
		}
	}

	return nil
}

func (s *Service) GetEnquiryById(id int) (services.EnquiryResponse, error) {
	enquiry, err := s.enquiriesRepository.SelectOne(id)
	if err != nil {
		return services.EnquiryResponse{}, err
	}
	cargos, cargoErr := s.cargosRepository.SelectAll(id)
	if cargoErr != nil {
		return services.EnquiryResponse{}, cargoErr
	}
	enquiryService := s.mapRepoDomainToService(enquiry)
	cargoService := s.mapRepoDomainToServiceCargo(cargos)
	enquiryResponse := &services.EnquiryResponse{
		Enquiry: enquiryService,
		Cargo:   cargoService,
	}
	return *enquiryResponse, nil
}

func (s *Service) GetAllEnquiries() ([]services.Enquiry, error) {
	var enquiriesList []services.Enquiry
	enquiry, err := s.enquiriesRepository.SelectAll()
	if err != nil {
		return []services.Enquiry{}, err
	}
	for _, v := range enquiry {
		enquiriesList = append(enquiriesList, s.mapRepoDomainToService(v))
	}
	return enquiriesList, nil
}

func (s *Service) DeleteEnquiry(id int) error {
	err := s.enquiriesRepository.DeleteOne(id)
	if err != nil {
		return err
	}
	errf := s.cargosRepository.DeleteAll(id)
	if errf != nil {
		return errf
	}
	return nil
}

func (s *Service) UpdateEnquiry(
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
) error {
	err := s.enquiriesRepository.UpdateOne(
		id,
		customer,
		terms_of_shipment,
		origin_customs,
		door_pickup,
		origin_port,
		pickup_address,
		cargo_ready_date,
		destination_customs,
		door_delivery,
		destination_port,
		delivery_address,
		hs_code,
		remarks,
	)
	if err != nil {
		return err
	}

	errf := s.cargosRepository.DeleteAll(id)
	if errf != nil {
		return errf
	}

	for _, v := range cargo {
		errr := s.cargosRepository.CreateOne(
			id,
			v.Count,
			v.Length,
			v.Width,
			v.Height,
			v.Volume,
			v.Weight,
			v.Unit,
			v.Total_volume,
			v.Total_weight,
		)
		if errr != nil {
			return errr
		}
	}
	return nil
}

func (s *Service) mapRepoDomainToService(
	enquiry rdbms.Enquiry,
) services.Enquiry {
	return services.Enquiry{
		Id:                  enquiry.Id,
		Customer:            enquiry.Customer,
		Terms_of_shipment:   enquiry.Terms_of_shipment,
		Origin_customs:      enquiry.Origin_customs,
		Door_pickup:         enquiry.Door_pickup,
		Origin_port:         enquiry.Origin_port,
		Pickup_address:      enquiry.Pickup_address,
		Cargo_ready_date:    enquiry.Cargo_ready_date,
		Destination_customs: enquiry.Destination_customs,
		Door_delivery:       enquiry.Door_delivery,
		Destination_port:    enquiry.Destination_port,
		Delivery_address:    enquiry.Delivery_address,
		Hs_code:             enquiry.Hs_code,
		Remarks:             enquiry.Remarks,
		Created_at:          enquiry.Created_at,
		Modified_at:         enquiry.Modified_at,
	}
}

func (s *Service) mapRepoDomainToServiceCargo(
	cargos []rdbms.Cargo,
) []services.Cargo {
	var cargoList []services.Cargo
	for _, v := range cargos {
		cargo := services.Cargo{
			Id:           v.Id,
			Enquiry_id:   v.Enquiry_id,
			Count:        v.Count,
			Length:       v.Length,
			Width:        v.Width,
			Height:       v.Height,
			Volume:       v.Volume,
			Weight:       v.Weight,
			Unit:         v.Unit,
			Total_volume: v.Total_volume,
			Total_weight: v.Total_weight,
		}
		cargoList = append(cargoList, cargo)
	}
	return cargoList
}
