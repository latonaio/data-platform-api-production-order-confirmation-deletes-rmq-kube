package dpfm_api_caller

import (
	dpfm_api_input_reader "data-platform-api-production-order-confirmation-deletes-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-production-order-confirmation-deletes-rmq-kube/DPFM_API_Output_Formatter"
	"fmt"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func (c *DPFMAPICaller) HeaderRead(
	input *dpfm_api_input_reader.SDC,
	log *logger.Logger,
) *dpfm_api_output_formatter.Header {
	where := fmt.Sprintf("WHERE header.ProductionOrderConfirmation = %d ", input.Header.ProductionOrderConfirmation)
	rows, err := c.db.Query(
		`SELECT 
			header.ProductionOrderConfirmation
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_production_order_CONFIRMAION_header_data as header 
		` + where + ` ;`)
	if err != nil {
		log.Error("%+v", err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToHeader(rows)
	if err != nil {
		log.Error("%+v", err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) ItemsRead(
	input *dpfm_api_input_reader.SDC,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Item {
	where := fmt.Sprintf("WHERE item.ProductionOrderConfirmation IS NOT NULL\nAND header.ProductionOrderConfirmation = %d", input.Header.ProductionOrderConfirmation)
	rows, err := c.db.Query(
		`SELECT 
			item.ProductionOrderConfirmation, item.ProductionOrderConfirmationItem
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_production_order_CONFIRMAION_item_data as item
		INNER JOIN DataPlatformMastersAndTransactionsMysqlKube.data_platform_production_order_CONFIRMAION_header_data as header
		ON header.ProductionOrderConfirmation = item.ProductionOrderConfirmation ` + where + ` ;`)
	if err != nil {
		log.Error("%+v", err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToItem(rows)
	if err != nil {
		log.Error("%+v", err)
		return nil
	}

	return data
}
