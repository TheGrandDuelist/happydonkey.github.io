package service

import (
	"github.com/ybkuroki/go-webapp-sample/container"
	"github.com/ybkuroki/go-webapp-sample/model"
)
// NewFormatService is constructor.
func NewFormatService(container container.Container) FormatService {
	return &formatService{container: container}
}


// FormatService is a service for managing master data such as format and category.
type FormatService interface {
	FindAllFormats() *[]model.Format
}

type formatService struct {
	container container.Container
}

// FindAllFormats returns the list of all formats.
func (m *formatService) FindAllFormats() *[]model.Format {
	
	result, err := format.FindAll(rep)
	rep := m.container.GetRepository()
	format := model.Format{}
	if err != nil {
		m.container.GetLogger().GetZapLogger().Errorf(err.Error())
		return nil
	}
	return result
}
