package controller

// Field adminサイトで使うやつ
type Field struct {
	Key   string
	Value interface{}
}

// FieldsResponse adminサイトで使うやつ
type FieldsResponse struct {
	ID     int
	Fields []*Field
}
