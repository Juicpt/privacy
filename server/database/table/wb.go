package table

import (
	"context"
	"database/sql"
	"github.com/kallydev/privacy/database"
	"github.com/kallydev/privacy/ent"
	"github.com/kallydev/privacy/ent/wbmodel"
)

var (
	_ database.Database = &WBDatabase{}
	_ database.Model    = &WBModel{}
)

type WBDatabase struct {
	Client *ent.Client
}

func (db *WBDatabase) QueryByQQNumber(ctx context.Context, qqNumber int64) ([]database.Model, error) {
	return []database.Model{}, nil
}

func (db *WBDatabase) QueryByWBNumber(ctx context.Context, wbNumber int64) ([]database.Model, error) {
	models, err := db.Client.WBModel.
		Query().
		Where(wbmodel.WbNumberEQ(wbNumber)).
		All(ctx)
	if err != nil {
		return nil, err
	}
	return entModelsToWBModels(models), nil
}

func (db *WBDatabase) QueryByEmail(ctx context.Context, email string) ([]database.Model, error) {
	return []database.Model{}, nil
}

func (db *WBDatabase) QueryByIDNumber(ctx context.Context, idNumber string) ([]database.Model, error) {
	return []database.Model{}, nil
}

func (db *WBDatabase) QueryByPhoneNumber(ctx context.Context, phoneNumber int64) ([]database.Model, error) {
	models, err := db.Client.WBModel.
		Query().
		Where(wbmodel.PhoneNumberEQ(phoneNumber)).
		All(ctx)
	if err != nil {
		return nil, err
	}
	return entModelsToWBModels(models), nil
}

type WBModel struct {
	WBNumber    sql.NullInt64
	PhoneNumber sql.NullInt64
}

func (model *WBModel) GetName() (name string, valid bool) {
	return "", false
}

func (model *WBModel) GetNickname() (nickname string, valid bool) {
	return "", false
}

func (model *WBModel) GetPassword() (password string, valid bool) {
	return "", false
}

func (model *WBModel) GetEmail() (email string, valid bool) {
	return "", false
}

func (model *WBModel) GetQQNumber() (qqNumber int64, valid bool) {
	return 0, false
}

func (model *WBModel) GetWBNumber() (wbNumber int64, valid bool) {
	return model.WBNumber.Int64, model.WBNumber.Valid
}

func (model *WBModel) GetIDNumber() (idNumber string, valid bool) {
	return "", false
}

func (model *WBModel) GetPhoneNumber() (phoneNumber int64, valid bool) {
	return model.PhoneNumber.Int64, model.PhoneNumber.Valid
}

func (model *WBModel) GetAddress() (address string, valid bool) {
	return "", false
}

func entModelsToWBModels(endModels []*ent.WBModel) []database.Model {
	models := make([]database.Model, len(endModels))
	for i, model := range endModels {
		models[i] = &WBModel{
			WBNumber: sql.NullInt64{
				Int64: model.WbNumber,
				Valid: model.WbNumber != 0,
			},
			PhoneNumber: sql.NullInt64{
				Int64: model.PhoneNumber,
				Valid: model.PhoneNumber != 0,
			},
		}
	}
	return models
}
