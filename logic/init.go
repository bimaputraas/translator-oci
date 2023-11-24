package logic

import "translate/model"

type Logic interface {
	Translate(model.TranslateData) (*model.TranslateData, interface{}, error)
}

type logicBasic struct{}

type logicAdvance struct{}

func InitLogicBasic() *logicBasic {
	return &logicBasic{}
}

func InitLogicAdvance() *logicAdvance {
	return &logicAdvance{}
}
