package models

// Relacion modelo para grabar la relacion de un usuario con otro
type Relacion struct {
	UsuarioID         string `bson:"usuario" json:"usuarioId"`
	UsuarioRelacionID string `bson:"usuariorelacionid" json:"usuarioRelacionId"`
}
