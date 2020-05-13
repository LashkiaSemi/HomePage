package controller

// Field adminサイトで使うやつ
type Field struct {
	Key   string      // カラムの名前 ex) ID,タイトル
	Value interface{} // カラムの値
	Type  string      // TODO: ワンチャン使ってない
}

// FieldsResponse adminサイトで使うやつ
type FieldsResponse struct {
	ID     int      // リソースのID
	Fields []*Field // リソースの全フィールド
}
