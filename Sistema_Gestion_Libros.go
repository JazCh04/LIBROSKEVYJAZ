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

	/*Creacion de usuarios
	Utilizamos un slice [] para crear varios usuarios ya que constantemente se puede
	requerir crear mas en el futuro*/

	usuarios := []Usuario{
		{
			UsuarioID:  001,
			Nombre:     "Juan Perez",
			Mail:       "juan.perez@correo.com",
			Contrasena: "librosjuan1",
			Rol:        "Usuario",
		},
		{
			UsuarioID:  002,
			Nombre:     "Maria Enriquez",
			Mail:       "maria.enriquez@correo.com",
			Contrasena: "mislibros123",
			Rol:        "Usuario",
		},
		{
			UsuarioID:  003,
			Nombre:     "Pedro Alvarez",
			Mail:       "pedro.alvarez@correo.com",
			Contrasena: "miperro5",
			Rol:        "Usuario",
		},
		{
			UsuarioID:  004,
			Nombre:     "Pablo Hernandez",
			Mail:       "pablo.hernandez@correo.com",
			Contrasena: "contra123",
			Rol:        "Usuario",
		},
		{
			UsuarioID:  005,
			Nombre:     "Samantha Rivera",
			Mail:       "samy.rivera@correo.com",
			Contrasena: "riosol159",
			Rol:        "Usuario",
		},
	}

	/*Creacion de libros
	Utilizamos un slice [] para crear varios libros ya que constantemente se puede
	requerir crear mas en el futuro*/

	libros := []Libro{
		{
			LibroID:          001,
			Titulo:           "Cartas de un Estoico",
			Autor:            "Lucio A. Séneca",
			FechaPublicacion: "2023",
			Genero:           "Filosofía",
			URL:              "www.libros.com/cartas_estoico",
		},
		{
			LibroID:          002,
			Titulo:           "Los Discursos de Epicteto",
			Autor:            "Epicteto",
			FechaPublicacion: "2023",
			Genero:           "Filosofía",
			URL:              "www.libros.com/discursos_epicteto",
		},
		{
			LibroID:          003,
			Titulo:           "Manual de Epicteto",
			Autor:            "Epicteto",
			FechaPublicacion: "1980",
			Genero:           "Filosofía",
			URL:              "www.libros.com/manual_epicteto",
		},
		{
			LibroID:          004,
			Titulo:           "Meditaciones",
			Autor:            "Marco Aurelio",
			FechaPublicacion: "2023",
			Genero:           "Filosofía",
			URL:              "www.libros.com/meditaciones",
		},
		{
			LibroID:          005,
			Titulo:           "Sobre la brevedad de la vida",
			Autor:            "Lucio A. Séneca",
			FechaPublicacion: "2023",
			Genero:           "Filosofía",
			URL:              "www.libros.com/brevedad_vida",
		},
	}

}
