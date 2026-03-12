package serverManagment

import "encoding/json"

func (s *ServerManagement) ConvertDataToJsonType(Jsons any) ([]byte, error) {

	return json.Marshal(Jsons)
}
