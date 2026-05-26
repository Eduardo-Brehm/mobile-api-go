package db

import "time"

type User struct {
	ID           string    `json:"id"`
	Nome         *string   `json:"nome"`
	Email        string    `json:"email"`
	SenhaHash    string    `json:"senhaHash"`
	CriadoEm     time.Time `json:"criadoEm"`
	AtualizadoEm time.Time `json:"atualizadoEm"`
}

type Profile struct {
	ID            string    `json:"id"`
	UsuarioID     string    `json:"usuarioId"`
	Tipo          string    `json:"tipo"`
	NomeDeUsuario string    `json:"nomeDeUsuario"`
	Bio           *string   `json:"bio"`
	FotoUrl       *string   `json:"fotoUrl"`
	SiteUrl       *string   `json:"siteUrl"`
	CriadoEm      time.Time `json:"criadoEm"`
	AtualizadoEm  time.Time `json:"atualizadoEm"`
}

type EstablishmentProfile struct {
	ID            string    `json:"id"`
	PerfilID      string    `json:"perfilId"`
	Categoria     *string   `json:"categoria"`
	Telefone      *string   `json:"telefone"`
	EnderecoTexto *string   `json:"enderecoTexto"`
	Cidade        *string   `json:"cidade"`
	UF            *string   `json:"uf"`
	Latitude      *float64  `json:"latitude"`
	Longitude     *float64  `json:"longitude"`
	Verificado    bool      `json:"verificado"`
	CriadoEm      time.Time `json:"criadoEm"`
	AtualizadoEm  time.Time `json:"atualizadoEm"`
}

type Follow struct {
	ID         string    `json:"id"`
	SeguidorID string    `json:"seguidorId"`
	SeguindoID string    `json:"seguindoId"`
	CriadoEm   time.Time `json:"criadoEm"`
}

type Avaliacao struct {
	ID                      string    `json:"id"`
	AvaliadorPerfilID       string    `json:"avaliadorPerfilId"`
	EstabelecimentoPerfilID string    `json:"estabelecimentoPerfilId"`
	Nota                    int       `json:"nota"`
	Comentario              *string   `json:"comentario"`
	CriadoEm                time.Time `json:"criadoEm"`
	AtualizadoEm            time.Time `json:"atualizadoEm"`
}

type OpiniaoTempoReal struct {
	ID                      string    `json:"id"`
	AvaliadorPerfilID       string    `json:"avaliadorPerfilId"`
	EstabelecimentoPerfilID string    `json:"estabelecimentoPerfilId"`
	Opiniao                 string    `json:"opiniao"` // BOMBANDO, BOM, RUIM
	CriadoEm                time.Time `json:"criadoEm"`
}
