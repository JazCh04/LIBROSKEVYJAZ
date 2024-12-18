/* Autores: Jazmin Chillagana & Kevin López
Fecha de creacion: 11/18/2024
Fecha de presentación: 12/18/2024
Descripcion: Sistema de Gestión de Libros Electrónicos
*/

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

//Definicion de Estructuras

// Administrador
type Administrador struct {
	AdministradorID int       `json:"id"`
	Nombre          string    `json:"nombre"`
	Mail            string    `json:"mail"`
	Contrasena      string    `json:"contrasena"`
	Rol             string    `json:"rol"`
	FechaCreacion   time.Time `json:"fecha_creacion"`
	UltimoAcceso    time.Time `json:"ultimo_acceso"`
}

// Usuario
type Usuario struct {
	UsuarioID  int    `json:"id"`
	Nombre     string `json:"nombre"`
	Mail       string `json:"mail"`
	Contrasena string `json:"contrasena"`
	Rol        string `json:"rol"`
}

// Inventario
type Inventario struct {
	InventarioId int  `json:"id"`
	LibroID      int  `json:"libro_id"`
	Disponible   bool `json:"disponible"`
}

// Libro
type Libro struct {
	LibroID          int    `json:"id"`
	Titulo           string `json:"titulo"`
	Autor            string `json:"autor"`
	FechaPublicacion string `json:"fecha_publicacion"`
	Genero           string `json:"genero"`
	Url              string `json:"url"`
}

// Prestamo
type Prestamo struct {
	PrestamoID      int       `json:"id"`
	LibroID         int       `json:"libro_id"`
	UsuarioID       int       `json:"usuario_id"`
	FechaReserva    time.Time `json:"fecha_reserva"`
	FechaDevolucion time.Time `json:"fecha_devolucion"`
}

// Respuesta de JSON
type Respuesta struct {
	Mensaje string `json:"mensaje"`
}

// Definicion de interfaces

// Interfaz para verificar permisos
type Permisos interface {
	Prestar() bool
	Devolver() bool
	AdministrarUsuario() bool
}

// Interfaz para realizar busquedas
type Busqueda interface {
	BuscarID(id int) (interface{}, error)
	BuscarNombre(nombre string) ([]interface{}, error)
}

// Interfaz para realizar serializacion
type Serializacion interface {
	MarshalJSON() ([]byte, error)
	UnmarshalJSON([]byte) error
}

// Código en HTML para visualizar las funcionalidades del sistema

// Código HTML para la página inicial
var homeTemplate = template.Must(template.New("Home").Parse(`
<!DOCTYPE html>
<html lang="es">
<head>
	<meta charset="UTF-8"> 
	<title>Home Page</title> 
</head> 
<body>
	<h1>Bienvenido al Sistema de Gestión de Libros de Jaz y Kev</h1>
	<h2>Creación: </h2> 
	<ul>
		<li><a href="/crear-admin">Crear Administrador</a></li> 
		<li><a href="/crear-user">Crear Usuario</a></li>
		<li><a href="/crear-book">Crear Libro</a></li>
	</ul>
	<h2>Visualización: </h2> 
	<ul>
		<li><a href="/visualizar-inv">Validar Inventario</a></li> 
		<li><a href="/visualizar-pres">Validar Préstamo</a></li> 
		<li><a href="/visualizar-admin">Visualizar Administrador</a></li> 
		<li><a href="/visualizar-libro">Visualizar Libro</a></li> 
		<li><a href="/visualizar-user">Visualizar Usuario</a></li>
	</ul>
	<h2>Búsqueda: </h2> 
	<ul>
		<li><a href="/buscar-libro">Buscar Libro por ID</a></li>
	</ul>
	<h2>Validar Permisos: </h2> 
	<ul>
		<li><a href="/validar-permisos">Consultar permisos</a></li>
	</ul>
	<footer>
	<p>Vuelve pronto</p>
	</footer>
</body>
</html>
`))

// Codigo HTML para la pagina de creacion de admin
var createAdmin = template.Must(template.New("create").Parse(`
<!DOCTYPE html>
<html lang="es">
<head>
    <meta charset="UTF-8">
    <title>Crear Administrador</title>
</head>
<body>
    <h1>Crear Nuevo Administrador</h1>
    <form action="/crear-admin" method="post">
        <label for="id">ID:</label>
        <input type="number" id="id" name="id" required><br>
        <label for="nombre">Nombre:</label>
        <input type="text" id="nombre" name="nombre" required><br>
        <label for="mail">Correo:</label>
        <input type="email" id="mail" name="mail" required><br>
        <label for="contrasena">Contraseña:</label>
        <input type="password" id="contrasena" name="contrasena" required><br>
        <label for="rol">Rol:</label>
        <input type="text" id="rol" name="rol" required><br>
        <button type="submit">Crear</button>
    </form>
    <footer>
        <p>Vuelve pronto</p>
    </footer>
</body>
</html>
`))

// Codigo HTML para la pagina de creacion de usuarios
var createUser = template.Must(template.New("create").Parse(`
<!DOCTYPE html>
<html lang="es">
<head>
    <meta charset="UTF-8">
    <title>Crear Usuario</title>
</head>
<body>
    <h1>Crear Nuevo Usuario</h1>
    <form action="/crear-user" method="post">
        <label for="id">ID:</label>
        <input type="number" id="id" name="id" required><br>
        <label for="nombre">Nombre:</label>
        <input type="text" id="nombre" name="nombre" required><br>
        <label for="mail">Correo:</label>
        <input type="email" id="mail" name="mail" required><br>
        <label for="contrasena">Contraseña:</label>
        <input type="password" id="contrasena" name="contrasena" required><br>
        <label for="rol">Rol:</label>
        <input type="text" id="rol" name="rol" required><br>
        <button type="submit">Crear</button>
    </form>
    <footer>
        <p>Vuelve pronto</p>
    </footer>
</body>
</html>
`))

// Codigo HTML para la pagina de creacion de libros
var createBook = template.Must(template.New("create").Parse(`
<!DOCTYPE html>
<html lang="es">
<head>
    <meta charset="UTF-8">
    <title>Crear Libro</title>
</head>
<body>
	<h1>Crear Nuevo Libro</h1> 
	<form action="/crear-book" method="post"> 
		<label for="id">ID:</label> 
		<input type="number" id="id" name="id" required><br>
		<label for="titulo">Título:</label> 
		<input type="text" id="titulo" name="titulo" required><br> 
		<label for="autor">Autor:</label> 
		<input type="text" id="autor" name="autor" required><br> 
		<label for="fechaPublicacion">Fecha de Publicación (YYYY - MONTH):</label> 
		<input type="text" id="fechaPublicacion" name="fechaPublicacion" required><br> 
		<label for="genero">Género:</label> 
		<input type="text" id="genero" name="genero" required><br> 
		<label for="url">URL:</label> 
		<input type="text" id="url" name="url" required><br> 
		<button type="submit">Crear</button>
    </form>
    <footer>
        <p>Vuelve pronto</p>
    </footer>
</body>
</html>
`))

// Codigo HTML para la pagina de busqueda
var searchTemplate = template.Must(template.New("busqueda").Parse(`
<!DOCTYPE html>
<html lang="es">
<head>
	<meta charset="UTF-8">
	<title>Buscar Libro</title>
</head>
<body>
	<h1>Buscar Libro por ID</h1>
	<form action="/buscar-libro" method="post">
		<label for="libroID">Ingrese el ID del libro:</label>
		<input type="number" id="libroID" name="libroID" required>
		<button type="submit">Buscar</button>
	</form>
	<footer>
		<p>Vuelve pronto</p>
	</footer>
</body>
</html>
`))

// Codigo HTML para la validacion de permisos
var verPermisos = template.Must(template.New("permisos").Parse(`
<!DOCTYPE html>
<html lang="es">
<head>
	<meta charset="UTF-8">
	<title>Validar permisos</title>
</head>
<body>
	<h1>Consultar permisos</h1>
	<form action="/validar-permisos" method="post">
		<label for="tipo">Seleccione el tipo de usuario:</label>
		<select name="tipo" id="tipo">
			<option value="administrador">Administrador</option>
			<option value="usuario">Usuario</option>
		</select>
		<br><br>
		<button type="submit">Consultar</button>
	</form>
	<footer>
		<p>Vuelve pronto</p>
	</footer>
</body>
</html>
`))

// Código HTML para la página de despedida
var away = template.Must(template.New("Away").Parse(`
<!DOCTYPE html>
<html lang="es">
<head>
	<meta charset="UTF-8"> 
	<title>Away Page</title> 
</head> 
<body>
	<h1>Gracias por visitar el Sistema de Gestión de Libros de Jaz y Kev</h1>
	<footer>
	<p>Vuelve pronto</p>
	</footer>
</body>
</html>
`))

// Aplicacion de metodos getter para poder acceder a las propiedades que estan encapsuladas

// Administrador
func (a *Administrador) GetNombre() string {
	return a.Nombre
}
func (a *Administrador) GetMail() string {
	return a.Mail
}
func (a *Administrador) GetRol() string {
	return a.Rol
}
func (a *Administrador) GetFechaCreacion() time.Time {
	return a.FechaCreacion
}
func (a *Administrador) GetUltimoAcceso() time.Time {
	return a.UltimoAcceso
}

// Usuario
func (u *Usuario) GetNombre() string {
	return u.Nombre
}
func (u *Usuario) GetMail() string {
	return u.Mail
}
func (u *Usuario) GetRol() string {
	return u.Rol
}

// Inventario
func (i *Inventario) GetInventario() int {
	return i.InventarioId
}
func (i *Inventario) GetLibroID() int {
	return i.LibroID
}
func (i *Inventario) IsDisponible() bool {
	return i.Disponible
}

// Libro
func (l *Libro) GetTitulo() string {
	return l.Titulo
}
func (l *Libro) GetAutor() string {
	return l.Autor
}
func (l *Libro) GetFechaPublicacion() string {
	return l.FechaPublicacion
}
func (l *Libro) GetGenero() string {
	return l.Genero
}
func (l *Libro) GetURL() string {
	return l.Url
}

// Prestamo
func (p *Prestamo) GetLibroID() int {
	return p.LibroID
}
func (p *Prestamo) GetUsuarioID() int {
	return p.UsuarioID
}
func (p *Prestamo) GetFechaReserva() time.Time {
	return p.FechaReserva
}
func (p *Prestamo) GetFechaDevolucion() time.Time {
	return p.FechaDevolucion
}

// Aplicacion de metodo setter para poder modificar las propiedades que estan encapsuladas

// Adminsitrador
func (a *Administrador) SetNombre(nombre string) {
	a.Nombre = nombre
}
func (a *Administrador) SetMail(mail string) {
	a.Mail = mail
}
func (a *Administrador) SetRol(rol string) {
	a.Rol = rol
}
func (a *Administrador) SetUltimoAcceso(t time.Time) {
	a.UltimoAcceso = t
}

// Usuario
func (u *Usuario) SetNombre(nombre string) {
	u.Nombre = nombre
}
func (u *Usuario) SetMail(mail string) {
	u.Mail = mail
}
func (u *Usuario) SetRol(rol string) {
	u.Rol = rol
}

// Inventario
func (i *Inventario) SetDisponible(disponible bool) {
	i.Disponible = disponible
}

// Libro
func (l *Libro) SetTirulo(titulo string) {
	l.Titulo = titulo
}
func (l *Libro) SetAutor(autor string) {
	l.Autor = autor
}
func (l *Libro) SetFechaPublicacion(fecha string) {
	l.FechaPublicacion = fecha
}
func (l *Libro) SetGenero(genero string) {
	l.Genero = genero
}
func (l *Libro) SetURL(url string) {
	l.Url = url
}

// Prestamo
func (p *Prestamo) SetFechaDevolucion(fecha time.Time) {
	p.FechaDevolucion = fecha
}

// Implementacion de validacion de permisos con la Interface "Permisos"

// Administrador
func (a *Administrador) Prestar() bool {
	return true
}
func (a *Administrador) Devolver() bool {
	return true
}
func (a *Administrador) AdministrarUsuario() bool {
	return true
}

// Usuario
func (u *Usuario) Prestar() bool {
	return true
}
func (u *Usuario) Devolver() bool {
	return true
}
func (u *Usuario) AdministrarUsuario() bool {
	return false
}

//Funcion para validar permisos

func consultarPermisos(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		if err := verPermisos.Execute(w, nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error al procesar el formulario", http.StatusBadRequest)
		return
	}

	// Obtener el tipo de usuario
	tipo := r.FormValue("tipo")
	var permisos Permisos

	switch tipo {
	case "administrador":
		permisos = &Administrador{}
	case "usuario":
		permisos = &Usuario{}
	default:
		http.Error(w, "Tipo de usuario inválido", http.StatusBadRequest)
		return
	}

	// Crear la respuesta JSON con los permisos
	respuesta := map[string]bool{
		"Prestar":            permisos.Prestar(),
		"Devolver":           permisos.Devolver(),
		"AdministrarUsuario": permisos.AdministrarUsuario(),
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(respuesta); err != nil {
		http.Error(w, "Error al codificar la respuesta", http.StatusInternalServerError)
	}
}

// Manejo de errores en creacion de administradores
func nuevoAdministrador(id int, nombre, mail, contrasena, rol string) (*Administrador, error) {
	if id <= 0 || nombre == "" || mail == "" || contrasena == "" {
		return nil, errors.New("error en los datos para crear un administrador")
	}
	return &Administrador{
		AdministradorID: id,
		Nombre:          nombre,
		Mail:            mail,
		Contrasena:      contrasena,
		Rol:             rol,
		FechaCreacion:   time.Now(),
		UltimoAcceso:    time.Now(),
	}, nil
}

func crearAdminstrador(w http.ResponseWriter, r *http.Request) {
	respuesta := Respuesta{"Crear Administrador"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(respuesta)
}

// Creamos la estructura adminsitrador con slice para guardar los nuevos administradores
type Listadoadmin struct {
	Administradores []*Administrador
}

func crearAdministrador(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		if err := createAdmin.Execute(w, nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Error al procesar el formulario", http.StatusBadRequest)
			return
		}

		id, err := strconv.Atoi(r.FormValue("id"))
		if err != nil {
			http.Error(w, "El ID debe ser un número entero", http.StatusBadRequest)
			return
		}

		nombre := r.FormValue("nombre")
		mail := r.FormValue("mail")
		contrasena := r.FormValue("contrasena")
		rol := r.FormValue("rol")

		admin, err := nuevoAdministrador(id, nombre, mail, contrasena, rol)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		listadoadmin.Administradores = append(listadoadmin.Administradores, admin)

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(admin); err != nil {
			http.Error(w, "Error al codificar la respuesta", http.StatusInternalServerError)
		}
	}
}

var listadoadmin Listadoadmin

// Manejo de errores en creacion de usuarios
func nuevoUsuario(id int, nombre, mail, contrasena, rol string) (*Usuario, error) {
	if id <= 0 || nombre == "" || mail == "" || contrasena == "" {
		return nil, errors.New("error en los datos para crear un usuario")
	}
	return &Usuario{
		UsuarioID:  id,
		Nombre:     nombre,
		Mail:       mail,
		Contrasena: contrasena,
		Rol:        rol,
	}, nil
}

// Creamos la estructura adminsitrador con slice para guardar los nuevos usuarios
type Listadouser struct {
	Usuarios []*Usuario
}

func crearUsuario(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		if err := createUser.Execute(w, nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Error al procesar el formulario", http.StatusBadRequest)
			return
		}

		id, err := strconv.Atoi(r.FormValue("id"))
		if err != nil {
			http.Error(w, "El ID debe ser un número entero", http.StatusBadRequest)
			return
		}

		nombre := r.FormValue("nombre")
		mail := r.FormValue("mail")
		contrasena := r.FormValue("contrasena")
		rol := r.FormValue("rol")

		user, err := nuevoUsuario(id, nombre, mail, contrasena, rol)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		listadouser.Usuarios = append(listadouser.Usuarios, user)

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(user); err != nil {
			http.Error(w, "Error al codificar la respuesta", http.StatusInternalServerError)
		}
	}
}

var listadouser Listadouser

// Manejo de errores en creacion de libros
func nuevoLibro(id int, titulo, autor string, fecha string, genero, url string) (*Libro, error) {
	if id <= 0 || titulo == "" || autor == "" {
		return nil, errors.New("error en los datos para crear un libro")
	}
	return &Libro{
		LibroID:          id,
		Titulo:           titulo,
		Autor:            autor,
		FechaPublicacion: fecha,
		Genero:           genero,
		Url:              url,
	}, nil
}

// Creaamos la estructura adminsitrador con slice para guardar los nuevos libros
type Listadobook struct {
	Libros []*Libro
}

func crearLibro(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		if err := createBook.Execute(w, nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Error al procesar el formulario", http.StatusBadRequest)
			return
		}

		id, err := strconv.Atoi(r.FormValue("id"))
		if err != nil {
			http.Error(w, "El ID debe ser un número entero", http.StatusBadRequest)
			return
		}

		titulo := r.FormValue("titulo")
		autor := r.FormValue("autor")
		fechaPublicacion := r.FormValue("fechaPublicacion")
		genero := r.FormValue("genero")
		url := r.FormValue("url")

		book, err := nuevoLibro(id, titulo, autor, fechaPublicacion, genero, url)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		listadobook.Libros = append(listadobook.Libros, book)

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(book); err != nil {
			http.Error(w, "Error al codificar la respuesta", http.StatusInternalServerError)
		}
	}
}

var listadobook Listadobook

// Busqueda de libros
type Libreria struct {
	Libros []*Libro
}

func (lib *Libreria) BuscarID(id int) (interface{}, error) {
	for _, libro := range lib.Libros {
		if libro.LibroID == id {
			return libro, nil
		}
	}
	return nil, errors.New("libro no encontrado con el ID digitado")
}

func (lib *Libreria) BuscarNombre(nombre string) ([]interface{}, error) {
	var resultados []interface{}
	for _, libro := range lib.Libros {
		if libro.Titulo == nombre {
			resultados = append(resultados, libro)
		}
	}
	if len(resultados) == 0 {
		return nil, errors.New("no existen libros con ese nombre")
	}
	return resultados, nil
}

func buscarLibro(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		if err := searchTemplate.Execute(w, nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Error al procesar el formulario", http.StatusBadRequest)
			return
		}

		libroID, err := strconv.Atoi(r.FormValue("libroID"))
		if err != nil {
			http.Error(w, "El ID del libro debe ser un número entero", http.StatusBadRequest)
			return
		}

		libro, err := libreria.BuscarID(libroID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(libro); err != nil {
			http.Error(w, "Error al codificar la respuesta", http.StatusInternalServerError)
		}
	}
}

var libreria *Libreria

// Manejo de errores para registro de inventarios
func nuevoInventario(id, libroID int, disponible bool) (*Inventario, error) {
	if id <= 0 || libroID <= 0 {
		return nil, errors.New("error en los datos para registrar en inventario")
	}
	return &Inventario{
		InventarioId: id,
		LibroID:      libroID,
		Disponible:   disponible,
	}, nil
}

// Manejo de errores para registro de prestamos
func nuevoPrestamo(id, libroID, usuarioID int, fechaReserva, fechaDevolucion time.Time) (*Prestamo, error) {
	if id <= 0 || libroID <= 0 || usuarioID <= 0 {
		return nil, errors.New("error en los datos para registrar un préstamo")
	}
	return &Prestamo{
		PrestamoID:      id,
		LibroID:         libroID,
		UsuarioID:       usuarioID,
		FechaReserva:    fechaReserva,
		FechaDevolucion: fechaDevolucion,
	}, nil
}

// Funciones para guardar y cargar en archivos JSON la información incluyendo manejo de errores
func saveToJSON(data interface{}, filename string) error {
	bytes, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, bytes, 0644)
}
func loadFromJSON(filename string, v interface{}) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, v)
}

// Funcion para visualizar el archivo json con los administradores
func visualizarAdministrador(w http.ResponseWriter, r *http.Request) {
	var administradores []*Administrador
	if err := loadFromJSON("administradores.json", &administradores); err != nil {
		http.Error(w, "Error al cargar el archivo", http.StatusInternalServerError)
		return
	}

	jsonBytes, err := json.Marshal(administradores)
	if err != nil {
		http.Error(w, "Error al serializar los administradores", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
}

// Funcion para visualizar el archivo json con los usuarios
func visualizarUsuario(w http.ResponseWriter, r *http.Request) {
	var usuarios []*Usuario
	if err := loadFromJSON("usuarios.json", &usuarios); err != nil {
		http.Error(w, "Error al cargar el archivo", http.StatusInternalServerError)
		return
	}

	jsonBytes, err := json.Marshal(usuarios)
	if err != nil {
		http.Error(w, "Error al serializar los usuarios", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
}

// Funcion para visualizar el archivo json con los libros
func visualizarLibro(w http.ResponseWriter, r *http.Request) {
	var libros []*Libro
	if err := loadFromJSON("libros.json", &libros); err != nil {
		http.Error(w, "Error al cargar el archivo", http.StatusInternalServerError)
		return
	}

	jsonBytes, err := json.Marshal(libros)
	if err != nil {
		http.Error(w, "Error al serializar los libros", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
}

// Funcion para visualizar el archivo json con el inventario
func visualizarInventario(w http.ResponseWriter, r *http.Request) {
	var inventario []*Inventario
	if err := loadFromJSON("inventario.json", &inventario); err != nil {
		http.Error(w, "Error al cargar el archivo", http.StatusInternalServerError)
		return
	}

	jsonBytes, err := json.Marshal(inventario)
	if err != nil {
		http.Error(w, "Error al serializar el inventario", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
}

// Funcion para visualizar el archivo json con los prestamos
func visualizarPrestamos(w http.ResponseWriter, r *http.Request) {
	var prestamos []*Prestamo
	if err := loadFromJSON("prestamos.json", &prestamos); err != nil {
		http.Error(w, "Error al cargar el archivo", http.StatusInternalServerError)
		return
	}

	jsonBytes, err := json.Marshal(prestamos)
	if err != nil {
		http.Error(w, "Error al serializar los prestamos", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
}

// Funcion para pagina de bienvenida y de despedida
func homePage(w http.ResponseWriter, r *http.Request) {
	if err := homeTemplate.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Funcion para pagina de despedida
func awayPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Gracias por visitar nuestra Biblioteca")
}

// Funcion Principal

func main() {

	/*Creacion de administradores
	Utilizamos un slice [] para crear varios administradores y pueda ser dinamico
	en caso de que se requiera crear mas en el futuro*/

	administradores := []*Administrador{
		{
			AdministradorID: 100,
			Nombre:          "Kevin Lopez",
			Mail:            "kevin.lopez@correo.com",
			Contrasena:      "contrasena100",
			Rol:             "Administrador",
			FechaCreacion:   time.Now(),
			UltimoAcceso:    time.Now(),
		},
		{
			AdministradorID: 200,
			Nombre:          "Jazmin Chillagana",
			Mail:            "jazmin.chillagana@correo.com",
			Contrasena:      "contrasena200",
			Rol:             "Administrador",
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

	libros := []*Libro{
		{
			LibroID:          001,
			Titulo:           "Cartas de un Estoico",
			Autor:            "Lucio A. Séneca",
			FechaPublicacion: "2024 September",
			Genero:           "Filosofía",
			Url:              "www.libros.com/cartas_estoico",
		},
		{
			LibroID:          002,
			Titulo:           "Los Discursos de Epicteto",
			Autor:            "Epicteto",
			FechaPublicacion: "2024 September",
			Genero:           "Filosofía",
			Url:              "www.libros.com/discursos_epicteto",
		},
		{
			LibroID:          003,
			Titulo:           "Manual de Epicteto",
			Autor:            "Epicteto",
			FechaPublicacion: "1980 May",
			Genero:           "Filosofía",
			Url:              "www.libros.com/manual_epicteto",
		},
		{
			LibroID:          004,
			Titulo:           "Meditaciones",
			Autor:            "Marco Aurelio",
			FechaPublicacion: "2023 October",
			Genero:           "Filosofía",
			Url:              "www.libros.com/meditaciones",
		},
		{
			LibroID:          005,
			Titulo:           "Sobre la brevedad de la vida",
			Autor:            "Lucio A. Séneca",
			FechaPublicacion: "2024 September",
			Genero:           "Filosofía",
			Url:              "www.libros.com/brevedad_vida",
		},
	}

	libreria = &Libreria{Libros: libros}

	/*Creacion de inventario
	Utilizamos un slice [] para crear varios libros ya que constantemente se puede
	requerir crear mas en el futuro*/

	inventario := []*Inventario{
		{
			InventarioId: 001,
			LibroID:      libros[0].LibroID,
			Disponible:   true,
		},
		{
			InventarioId: 002,
			LibroID:      libros[1].LibroID,
			Disponible:   true,
		},
		{
			InventarioId: 003,
			LibroID:      libros[2].LibroID,
			Disponible:   true,
		},
		{
			InventarioId: 004,
			LibroID:      libros[3].LibroID,
			Disponible:   true,
		},
		{
			InventarioId: 005,
			LibroID:      libros[4].LibroID,
			Disponible:   true,
		},
	}

	/*Creacion de prestamos
	Utilizamos un slice [] para crear varios prestamos ya que constantemente se puede
	requerir crear mas en el futuro*/

	prestamos := []*Prestamo{
		{
			PrestamoID:      001,
			LibroID:         libros[0].LibroID,
			UsuarioID:       usuarios[0].UsuarioID,
			FechaReserva:    time.Now(),
			FechaDevolucion: time.Now().AddDate(0, 0, 5),
		},
		{
			PrestamoID:      002,
			LibroID:         libros[1].LibroID,
			UsuarioID:       usuarios[1].UsuarioID,
			FechaReserva:    time.Now(),
			FechaDevolucion: time.Now().AddDate(0, 0, 5),
		},
		{
			PrestamoID:      003,
			LibroID:         libros[2].LibroID,
			UsuarioID:       usuarios[2].UsuarioID,
			FechaReserva:    time.Now(),
			FechaDevolucion: time.Now().AddDate(0, 0, 5),
		},
		{
			PrestamoID:      004,
			LibroID:         libros[3].LibroID,
			UsuarioID:       usuarios[3].UsuarioID,
			FechaReserva:    time.Now(),
			FechaDevolucion: time.Now().AddDate(0, 0, 5),
		},
		{
			PrestamoID:      005,
			LibroID:         libros[4].LibroID,
			UsuarioID:       usuarios[4].UsuarioID,
			FechaReserva:    time.Now(),
			FechaDevolucion: time.Now().AddDate(0, 0, 5),
		},
	}

	//Funcion para guardar la informacion en los archivos JSON
	//Administradores
	if err := saveToJSON(administradores, "administradores.json"); err != nil {
		fmt.Println("Error al guardar los administradores:", err)
	}

	//Usuarios
	if err := saveToJSON(usuarios, "usuarios.json"); err != nil {
		fmt.Println("Error al guardar los usuarios:", err)
	}

	//Libros
	if err := saveToJSON(libros, "libros.json"); err != nil {
		fmt.Println("Error al guardar los registros de libros:", err)
	}

	//Inventarios
	if err := saveToJSON(inventario, "inventario.json"); err != nil {
		fmt.Println("Error al guardar el inventario:", err)
	}

	//Prestamos
	if err := saveToJSON(prestamos, "prestamos.json"); err != nil {
		fmt.Println("Error al guardar los prestamos:", err)
	}

	//Generamos el servicio web para ver nuestras funcionalidades

	http.HandleFunc("/", homePage)
	http.HandleFunc("/crear-admin", crearAdministrador)
	http.HandleFunc("/crear-user", crearUsuario)
	http.HandleFunc("/crear-book", crearLibro)
	http.HandleFunc("/visualizar-admin", visualizarAdministrador)
	http.HandleFunc("/visualizar-user", visualizarUsuario)
	http.HandleFunc("/visualizar-libro", visualizarLibro)
	http.HandleFunc("/visualizar-inv", visualizarInventario)
	http.HandleFunc("/visualizar-pres", visualizarPrestamos)
	http.HandleFunc("/buscar-libro", buscarLibro)
	http.HandleFunc("/validar-permisos", consultarPermisos)
	http.HandleFunc("/away", awayPage)

	fmt.Println("Servidor iniciado en el puerto 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

	//Imprimir detalles de administradores
	for _, admin := range administradores {
		fmt.Printf("Administrador: %+v\n", admin)
	}

	/*//Imprimir detalles de usuarios
	for _, usuario := range usuarios {
		fmt.Printf("Usuario: %+v\n", usuario)
	}

	// Imprimir detalles de inventario
	for _, item := range inventario {
		fmt.Printf("Inventario: %+v\n", item)
	}

	// Imprimir detalles de prestamos
	for _, prestamos := range prestamos {
		fmt.Printf("Prestamos: %+v\n", prestamos)
	}

	// Prueba de manejo de errores intentando crear un prestamo sin datos
	admin, err := nuevoAdministrador(0, "", "", "", "")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Préstamo registrado: %+v\n", admin)
	}

	user, err := nuevoUsuario(0, "", "", "", "")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Préstamo registrado: %+v\n", user)
	}

	libro, err := nuevoLibro(0, "", "", time.Date(2024, time.September, 21, 0, 0, 0, 0, time.UTC), "", "")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Préstamo registrado: %+v\n", libro)
	}

	invent, err := nuevoInventario(0, libros[0].libroID, false)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Préstamo registrado: %+v\n", invent)
	}

	prestamo, err := nuevoPrestamo(0, libros[0].libroID, usuarios[0].usuarioID, time.Now(), time.Now().AddDate(0, 0, 5))
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Préstamo registrado: %+v\n", prestamo)
	}*/
}
