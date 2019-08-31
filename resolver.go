package login_template_graphql

import (
	"context"
	"demo/models"
	"fmt"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

var usuarios = []*models.Usuario{
	&models.Usuario{
		Id: "789798",
		Nome:"Nestor",
		Email:    "teste@uol.com.br",
		Password: "temtem"},
	&models.Usuario{
		Id: "123",
		Nome:"Cacau",
		Email:    "joao@terra.com.br",
		Password: "123"},
	&models.Usuario{
		Id: "789798",
		Nome:"Rosdolfo",
		Email:    "beto@zaptec.com.br",
		Password: "321"},
	&models.Usuario{
		Id: "0000",
		Nome:"Luis Otavio",
		Email:    "agioanldo@terra.com.br",
		Password: "456"},
	&models.Usuario{
		Id: "000",
		Nome:"Humberto",
		Email:    "giogio@afsdfdsa.com.br",
		Password: "7777"},
}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CriarUsuario(ctx context.Context, input models.Usuario) (*models.Usuario, error) {
	
	fmt.Println(input)
	usuarios = append(usuarios, &input)
	return &input, nil
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Usuarios(ctx context.Context) ([]*models.Usuario, error) {
	
	return usuarios, nil
}
func (r *queryResolver) Login(ctx context.Context, email string, password string) (*models.Usuario, error) {
	return usuarios[0], nil
}
