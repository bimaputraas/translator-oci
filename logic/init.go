package logic

import "translate/model"

type Logic interface {
	Translate(model.TranslateData) (*model.TranslateData, interface{}, error)
}

type logicBasic struct{}

type logicAdvance struct{}

func InitLogicBasic() Logic {
	return &logicBasic{}
}

func InitLogicAdvance() Logic {
	return &logicAdvance{}
}
