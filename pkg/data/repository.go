package data

type DataRepository interface {
	GetAll() ([]*Type, error)
	Save(data *Type) error
}
