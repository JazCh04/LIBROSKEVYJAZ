/* Autores: Jazmin Chillagana & Kevin López
Fecha de creacion: 11/18/2024
Descripcion: Sistema de Gestión de Libros Electrónicos
*/

package main

import (
	"fmt"
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
	fmt.Println("¡Hola, mundo!")
}
