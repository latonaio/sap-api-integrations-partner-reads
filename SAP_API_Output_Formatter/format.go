package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-partner-reads/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToPartnerCollection(raw []byte, l *logger.Logger) ([]PartnerCollection, error) {
	pm := &responses.PartnerCollection{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to PartnerCollection. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	partnerCollection := make([]PartnerCollection, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		partnerCollection = append(partnerCollection, PartnerCollection{
			ObjectID:                          data.ObjectID,
			PartnerID:                         data.PartnerID,
			PartnerUUID:                       data.PartnerUUID,
			LifeCycleStatusCode:               data.LifeCycleStatusCode,
			LifeCycleStatusCodeText:           data.LifeCycleStatusCodeText,
			Duns:                              data.Duns,
			LegalFormCode:                     data.LegalFormCode,
			LegalFormCodeText:                 data.LegalFormCodeText,
			PartnerABCClassificationCode:      data.PartnerABCClassificationCode,
			PartnerABCClassificationCodeText:  data.PartnerABCClassificationCodeText,
			IndustrialSectorCode:              data.IndustrialSectorCode,
			IndustrialSectorCodeText:          data.IndustrialSectorCodeText,
			ContactPermissionCode:             data.ContactPermissionCode,
			ContactPermissionCodeText:         data.ContactPermissionCodeText,
			BusinessPartnerFormattedName:      data.BusinessPartnerFormattedName,
			Name:                              data.Name,
			AdditionalName:                    data.AdditionalName,
			AdditionalName2:                   data.AdditionalName2,
			AdditionalName3:                   data.AdditionalName3,
			FormattedPostalAddressDescription: data.FormattedPostalAddressDescription,
			CountryCode:                       data.CountryCode,
			CountryCodeText:                   data.CountryCodeText,
			StateCode:                         data.StateCode,
			StateCodeText:                     data.StateCodeText,
			CareOfName:                        data.CareOfName,
			AddressLine1:                      data.AddressLine1,
			AddressLine2:                      data.AddressLine2,
			HouseNumber:                       data.HouseNumber,
			Street:                            data.Street,
			AddressLine4:                      data.AddressLine4,
			AddressLine5:                      data.AddressLine5,
			District:                          data.District,
			City:                              data.City,
			DifferentCity:                     data.DifferentCity,
			StreetPostalCode:                  data.StreetPostalCode,
			County:                            data.County,
			CompanyPostalCode:                 data.CompanyPostalCode,
			POBox:                             data.POBox,
			POBoxPostalCode:                   data.POBoxPostalCode,
			POBoxDeviatingCountryCode:         data.POBoxDeviatingCountryCode,
			POBoxDeviatingCountryCodeText:     data.POBoxDeviatingCountryCodeText,
			POBoxDeviatingStateCode:           data.POBoxDeviatingStateCode,
			POBoxDeviatingStateCodeText:       data.POBoxDeviatingStateCodeText,
			POBoxDeviatingCity:                data.POBoxDeviatingCity,
			TimeZoneCode:                      data.TimeZoneCode,
			TimeZoneCodeText:                  data.TimeZoneCodeText,
			Building:                          data.Building,
			Floor:                             data.Floor,
			Room:                              data.Room,
			Phone:                             data.Phone,
			Mobile:                            data.Mobile,
			Fax:                               data.Fax,
			Email:                             data.Email,
			WebSite:                           data.WebSite,
			LanguageCode:                      data.LanguageCode,
			LanguageCodeText:                  data.LanguageCodeText,
			BestReachedByCode:                 data.BestReachedByCode,
			BestReachedByCodeText:             data.BestReachedByCodeText,
			TotalPoints:                       data.TotalPoints,
			PartnerLevel:                      data.PartnerLevel,
			PartnerLevelText:                  data.PartnerLevelText,
			NormalisedPhone:                   data.NormalisedPhone,
			NormalisedMobile:                  data.NormalisedMobile,
			CreatedOn:                         data.CreatedOn,
			CreatedBy:                         data.CreatedBy,
			CreatedByIdentityUUID:             data.CreatedByIdentityUUID,
			ChangedOn:                         data.ChangedOn,
			ChangedBy:                         data.ChangedBy,
			ChangedByIdentityUUID:             data.ChangedByIdentityUUID,
			EntityLastChangedOn:               data.EntityLastChangedOn,
			ETag:                              data.ETag,
		})
	}

	return partnerCollection, nil
}
