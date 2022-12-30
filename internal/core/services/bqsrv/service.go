package bqsrv

import (
	"sagebackend/internal/core/domain/repositories/rdbms"
	"sagebackend/internal/core/domain/services"
	"sagebackend/internal/core/ports"
)

type Service struct {
	chargesRepository ports.RdbmsChargesRepository
	bqRepository ports.RdbmsBqRepository
}

func New(
	chargesRepository ports.RdbmsChargesRepository,
	bqRepository ports.RdbmsBqRepository,
	) *Service {
	return &Service{
		chargesRepository: chargesRepository,
		bqRepository: bqRepository,
	}
}

func (s *Service) CreateBq(
	is_quote bool,
	enquiry_id int,
	valid_till string,
	vessel_liner string,
	vessel_total_transit string,
	vessel_free_days string,
	freight_charges_vendor string,
    freight_other_charges_vendor string,
    total_buy_amount int,
    total_sell_amount int,
    profit_percentage string,
	charges []services.Charge,
) error {
	id,err := s.bqRepository.CreateOne(
		is_quote,
		enquiry_id,
		valid_till,
		vessel_liner,
		vessel_total_transit,
		vessel_free_days,
		freight_charges_vendor,
		freight_other_charges_vendor,
		total_buy_amount,
		total_sell_amount,
		profit_percentage,
)
	if err != nil {
		return err
	}
	for _, charge := range charges {
		err := s.chargesRepository.CreateOne(
			id,
			charge.Charges_name,
			charge.Units,
			charge.Tax,
			charge.Buy_rate,
			charge.Sell_rate,
			charge.Total_buy_amount,
			charge.Total_sell_amount,
			charge.Charges_type,
		)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Service) GetBqByID(
	id int,
) (services.BqResponse, error) {
	bq, err := s.bqRepository.SelectOne(id)
	if err != nil {
		return services.BqResponse{}, err
	}
	charges, err := s.chargesRepository.SelectAll(id)
	if err != nil {
		return services.BqResponse{}, err
	}
	bqService := s.mapRepoDomainToService(bq)
	chargeService := s.mapRepoDomainToServiceCharge(charges)
	bqResponse := services.BqResponse{
		Bq: bqService,
		Charge: chargeService,
	}
	return bqResponse, nil
}

func (s *Service) mapRepoDomainToService(
	bq rdbms.Bq,
) services.Bq {
	return services.Bq{
		Id: bq.Id,
		Is_quote: bq.Is_quote,
		Valid_till: bq.Valid_till,
		Vessel_liner: bq.Vessel_liner,
		Vessel_total_transit: bq.Vessel_total_transit,
		Vessel_free_days: bq.Vessel_free_days,
		Freight_charges_vendor: bq.Freight_charges_vendor,
		Freight_other_charges_vendor: bq.Freight_other_charges_vendor,
		Total_buy_amount: bq.Total_buy_amount,
		Total_sell_amount: bq.Total_sell_amount,
		Profit_percentage: bq.Profit_percentage,
		Created_at: bq.Created_at,
	}
}

func (s *Service) mapRepoDomainToServiceCharge(
	charges []rdbms.Charges,
) []services.Charge {
	var chargeService []services.Charge
	for _, charge := range charges {
	temp :=  services.Charge{
		Id: charge.Id,
		Bq_id: charge.Bq_id,
		Charges_name: charge.Charges_name,
		Units: charge.Units,
		Tax: charge.Tax,
		Buy_rate: charge.Buy_rate,
		Sell_rate: charge.Sell_rate,
		Total_buy_amount: charge.Total_buy_amount,
		Total_sell_amount: charge.Total_sell_amount,
		Charges_type: charge.Charges_type,
	}
	chargeService = append(chargeService, temp)
}
return chargeService
}

func (s *Service) UpdateBq(
	id int,
	is_quote bool,
	enquiry_id int,
	valid_till string,
	vessel_liner string,
	vessel_total_transit string,
	vessel_free_days string,
	freight_charges_vendor string,
	freight_other_charges_vendor string,
	total_buy_amount int,
	total_sell_amount int,
	profit_percentage string,
	charges []services.Charge,
) error {
	err := s.bqRepository.UpdateOne(
		id,
		enquiry_id,
		is_quote,
		valid_till,
		vessel_liner,
		vessel_total_transit,
		vessel_free_days,
		freight_charges_vendor,
		freight_other_charges_vendor,
		total_buy_amount,
		total_sell_amount,
		profit_percentage,
	)
	if err != nil {
		return err
	}

	errf := s.chargesRepository.DeleteAll(id)
	if errf != nil {
		return errf
	}

	for _, charge := range charges {
		err := s.chargesRepository.CreateOne(
			charge.Bq_id,
			charge.Charges_name,
			charge.Units,
			charge.Tax,
			charge.Buy_rate,
			charge.Sell_rate,
			charge.Total_buy_amount,
			charge.Total_sell_amount,
			charge.Charges_type,
		)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Service) DeleteBq(id int)(error) {
	err := s.bqRepository.DeleteOne(id)
	if err != nil {
		return err
	}
	errf := s.chargesRepository.DeleteAll(id)
	if errf != nil {
		return errf
	}
	return nil
}

func (s *Service) GetAllBq() ([]services.Bq, error) {
	var bqs []services.Bq
	bq, err := s.bqRepository.SelectAll()
	if err != nil {
		return []services.Bq{}, err
	}
	for _, bq := range bq {
		bqService := s.mapRepoDomainToService(bq)
		bqs = append(bqs, bqService)
	}
	return bqs, nil
}

