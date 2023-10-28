package shared

import (
	"library-management/shared/dto"
)

func FormatResponse(data interface{}, err error) dto.Response {
	if err != nil {
		return dto.Response{
			Error: TranslateError(err),
			Data:  nil,
		}
	}

	return dto.Response{
		Error: "",
		Data:  data,
	}
}
