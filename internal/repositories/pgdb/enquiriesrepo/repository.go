package enquiriesrepo

import (
	"aircargo/internal/core/domain/repositories/rdbms"
	"aircargo/pkg/logging"
	"fmt"

	"github.com/doug-martin/goqu/v9"
)

type Repository struct {
	logger logging.Logger
	goquDB *goqu.Database
}

func New(
	logger logging.Logger,
	goquDB *goqu.Database,
) *Repository {
	return &Repository{
		logger: logger,
		goquDB: goquDB,
	}
}
type Id struct {
	Id int `db:"id"`
}
func (r *Repository) CreateOne(
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
) (int,error) {
	var id Id
	_, err := r.goquDB.Insert(TABLE).Prepared(true).Rows(
		goqu.Record{
			CUSTOMER:        customer,
			TERMS_OF_SHIPMENT:       terms_of_shipment,
			ORIGIN_CUSTOMS:       origin_customs,
			DOOR_PICKUP:     door_pickup,
			ORIGIN_PORT: 	 origin_port,
			PICKUP_ADDRESS:  pickup_address,
			CARGO_READY_DATE:     cargo_ready_date,
			DESTINATION_CUSTOMS:        destination_customs,
			DOOR_DELIVERY:       door_delivery,
			DESTINATION_PORT:       destination_port,
			DELIVERY_ADDRESS:     delivery_address,
			HS_CODE: 	 hs_code,
			REMARKS:  remarks,
		},
	).Returning(ID).Executor().ScanStruct(&id)
	// fmt.Println(err, account_id, "contactrepo")
	if err != nil {
		fmt.Println(id,"id here")
		return 0,err
	}
	
	return id.Id, nil
}

func(r *Repository) SelectOne(id int)(rd rdbms.Enquiry, err error){
	var enquiry rdbms.Enquiry
    exists, err := r.goquDB.From(TABLE).Select().Where(goqu.C(ID).Eq(id)).ScanStruct(&enquiry) 

	if err != nil {
		return rdbms.Enquiry{}, err
	}
	fmt.Print(exists)
	return enquiry, nil
}

func(r *Repository) DeleteOne(id int) error {
	_, err := r.goquDB.Delete(TABLE).Where(goqu.C(ID).Eq(id)).Executor().Exec()
	if err != nil {
		return err
	}
	return nil
}

func(r *Repository) UpdateOne(
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
) error {
	_, err := r.goquDB.Update(TABLE).Set(
		goqu.Record{
			CUSTOMER:        customer,
			TERMS_OF_SHIPMENT:       terms_of_shipment,
			ORIGIN_CUSTOMS:       origin_customs,
			DOOR_PICKUP:     door_pickup,
			ORIGIN_PORT: 	 origin_port,
			PICKUP_ADDRESS:  pickup_address,
			CARGO_READY_DATE:     cargo_ready_date,
			DESTINATION_CUSTOMS:        destination_customs,
			DOOR_DELIVERY:       door_delivery,
			DESTINATION_PORT:       destination_port,
			DELIVERY_ADDRESS:     delivery_address,
			HS_CODE: 	 hs_code,
			REMARKS:  remarks,
		},
	).Where(goqu.C(ID).Eq(id)).Executor().Exec()
	if err != nil {
		return err
	}
	return nil
}

func(r *Repository) SelectAll()([]rdbms.Enquiry, error){
	var enquiries []rdbms.Enquiry
	err := r.goquDB.From(TABLE).Select().Order(
		goqu.I(CREATED_AT).Asc(),
	).ScanStructs(&enquiries) 

	if err != nil {
		return nil, err
	}
	return enquiries, nil
}
