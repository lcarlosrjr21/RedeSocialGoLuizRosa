package modelos

import (
	"errors"
	"strings"
	"time"
)

// Publicacao representa uma publicacao (post) feita pelo usuario
type Publicacao struct {
	ID        uint64    `json:"id,omitempy`
	Titulo    string    `json:"titulo,omitempy`
	Conteudo  string    `json:"conteudo,omitempy`
	AutorID   uint64    `json:"autorId,omitempy`
	AutorNick string    `json:"autorNick,omitempy`
	Curtidas  uint64    `json:"curtidas`
	CriadaEm  time.Time `json:"criadaEm,omitempy`
}

func (publicacao *Publicacao) Preparar() error {
	if erro := publicacao.validar(); erro != nil {
		return erro
	}

	publicacao.formatar()
	return nil
}

// verifica se o campo titulo e conteudo estao vazios
func (publicacao *Publicacao) validar() error {
	if publicacao.Titulo == "" {
		return errors.New("O titulo nao pode ser vazio")
	}

	if publicacao.Conteudo == "" {
		return errors.New("O conteudo nao pode ser vazio")
	}

	return nil

}

// remove os espacos de titulo e conteudo
func (publicacao *Publicacao) formatar() {
	publicacao.Titulo = strings.TrimSpace(publicacao.Titulo)
	publicacao.Conteudo = strings.TrimSpace(publicacao.Conteudo)
}
