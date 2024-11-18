/* Autores: Jazmin Chillagana & Kevin López
Fecha de creacion: 11/18/2024
Descripcion: Sistema de Gestión de Libros Electrónicos
*/

package main

import (
	"time"
)

//Definicion de Estructuras

// Administrador
type Administrador struct {
	AdministradorID int
	Nombre          string
	Mail            string
	Contrasena      string
	Rol             string
	FechaCreacion   time.Time
	UltimoAcceso    time.Time
}

// Usuario
type Usuario struct {
	UsuarioID  int
	Nombre     string
	Mail       string
	Contrasena string
	Rol        string
}

// Inventario
type Inventario struct {
	InventarioId int
	LibroID      int
	Disponible   bool
}

// Libro
type Libro struct {
	LibroID          int
	Titulo           string
	Autor            string
	FechaPublicacion time.Time
	Genero           string
	URL              string
}

// Prestamo
type Prestamo struct {
	PrestamoID      int
	LibroID         int
	UsuarioID       int
	FechaReserva    time.Time
	FechaDevolucion time.Time
}

// Funcion Principal

func main() {

	/*Creacion de administradores
	Utilizamos un slice [] para crear varios administradores y pueda ser dinamico
	en caso de que se requiera crear mas en el futuro*/

	administradores := []Administrador{
		{
			AdministradorID: 100,
			Nombre:          "Kevin Lopez",
			Mail:            "kevin.lopez@correo.com",
			Contrasena:      "contrasena100",
			Rol:             "Admin1",
			FechaCreacion:   time.Now(),
			UltimoAcceso:    time.Now(),
		},
		{
			AdministradorID: 200,
			Nombre:          "Jazmin Chillagana",
			Mail:            "jazmin.chillagana@correo.com",
			Contrasena:      "contrasena200",
			Rol:             "Admin2",
			FechaCreacion:   time.Now(),
			UltimoAcceso:    time.Now(),
		},
	}
}
