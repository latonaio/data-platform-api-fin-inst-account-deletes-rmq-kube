package dpfm_api_caller

import (
	dpfm_api_input_reader "data-platform-api-fin-inst-account-deletes-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-fin-inst-account-deletes-rmq-kube/DPFM_API_Output_Formatter"
	"fmt"
	"strings"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func (c *DPFMAPICaller) HeaderRead(
	input *dpfm_api_input_reader.SDC,
	log *logger.Logger,
) *dpfm_api_output_formatter.Header {
	where := strings.Join([]string{
		fmt.Sprintf("WHERE header.FinInstCountry = %d ", input.Header.FinInstCountry),
	}, "")

	rows, err := c.db.Query(
		`SELECT 
    	header.FinInstCountry
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_fin_inst_account_header_data as header 
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
	where := strings.Join([]string{
		fmt.Sprintf("WHERE item.FinInstCountry = %d ", input.Header.FinInstCountry),
	}, "")

	rows, err := c.db.Query(
		`SELECT 
    	item.FinInstCountry
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_fin_inst_account_item_data as item
		` + where + ` ;`)
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
