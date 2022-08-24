package model

import "time"

type History struct {
	Id          int       `json:"id"`
	Address     string    `json:"address,omitempty"`
	Recipient   string    `json:"recipient,omitempty"`
	SrcChain    string    `json:"src_chain,omitempty"`
	DestChain   string    `json:"dest_chain,omitempty"`
	TokenSymbol string    `json:"token_symbol,omitempty"`
	Amount      string    `json:"amount,omitempty"`
	SrcHash     string    `json:"src_hash,omitempty"`
	DestHash    string    `json:"dest_hash,omitempty"`
	SrcLink     string    `json:"src_link,omitempty"`
	DestLink    string    `json:"dest_link,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
