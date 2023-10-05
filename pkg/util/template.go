package util

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
	"github.com/riyan-eng/golang-boilerplate-one/env"
)

type templateStruct struct{}

func NewTemplate() *templateStruct {
	return &templateStruct{}
}

func (t *templateStruct) EmailResetPassword(token string, expired *jwt.NumericDate) (template string) {
	urlFe := fmt.Sprintf(`%v/setel-ulang-password/?token=%v`, env.SERVER_HOST_FE, token)
	template = fmt.Sprintf(`
		<div style="flex: auto; text-align: center;">
			<h1>Sandbox Indonesia</h1>
			<p>Berikut adalah link untuk melakukan reset password: <a href="%v">%v</a></p>
			<p>hanya berlaku sampai %v, dan hanya bisa digunakan satu kali.</p>
		</div>
	`, urlFe, urlFe, expired)
	return
}
