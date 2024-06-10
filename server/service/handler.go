package service

import (
	"github.com/labstack/echo/v4"
	"strconv"
)

type QueryResponse struct {
	Names        []string `json:"names"`
	Nicknames    []string `json:"nicknames"`
	PhoneNumbers []string `json:"phone_numbers"`
	IDNumbers    []string `json:"id_numbers"`
	QQNumbers    []string `json:"qq_numbers"`
	WBNumbers    []string `json:"wb_uids"`
	Passwords    []string `json:"passwords"`
	Emails       []string `json:"emails"`
	Addresses    []string `json:"addresses"`
}

func NewQueryResponse() *QueryResponse {
	return &QueryResponse{
		Names:        make([]string, 0),
		Nicknames:    make([]string, 0),
		PhoneNumbers: make([]string, 0),
		IDNumbers:    make([]string, 0),
		QQNumbers:    make([]string, 0),
		WBNumbers:    make([]string, 0),
		Passwords:    make([]string, 0),
		Emails:       make([]string, 0),
		Addresses:    make([]string, 0),
	}
}

func (svc *Service) queryHandlerFunc(ctx echo.Context) error {
	value := ctx.QueryParam("value")
	if value == "" {
		return InvalidParameterError
	}

	result := NewQueryResult()

	if value[0] == '@' {
		trimmedValue := value[1:]
		wbNumber, err := strconv.ParseInt(trimmedValue, 10, 64)
		if err != nil {
			return err
		}
		if err := result.queryWBNumber(svc.databases, wbNumber); err != nil {
			return err
		}

		for phoneNumber := range result.PhoneNumbers {
			if !result.PhoneNumbers[phoneNumber] {
				if err := result.queryPhoneNumber(svc.databases, int64(phoneNumber)); err != nil {
					return err
				}
			}
		}
	} else if value[0] == '+' {
		trimmedValue := value[1:]
		phoneNumber, err := strconv.ParseInt(trimmedValue, 10, 64)
		if err != nil {
			return err
		}
		if err := result.queryPhoneNumber(svc.databases, phoneNumber); err != nil {
			return err
		}
	} else {
		types := whatType(value)
		for _, t := range types {
			switch value := t.(type) {
			case QQNumber:
				result.addQQNumber(int64(value))
			case Email:
				result.addEmail(string(value))
			case IDNumber:
				result.addIDNumber(string(value))
			}
		}
	}

	response := result.Build(svc.config.Mask)
	return NewResponse(ctx, nil, response)
}

