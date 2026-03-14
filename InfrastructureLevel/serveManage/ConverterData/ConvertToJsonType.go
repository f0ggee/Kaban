package ConverterData

import (
	"encoding/json"
)

func (s *ConvertingData) ConvertDataToJsonType(Jsons any) ([]byte, error) {

	return json.Marshal(Jsons)
}
