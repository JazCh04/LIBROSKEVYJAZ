/* Autores: Jazmin Chillagana & Kevin López
Fecha de creacion: 11/18/2024
Descripcion: Sistema de Gestión de Libros Electrónicos
*/

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

//Definicion de Estructuras

// Administrador
type Administrador struct {
	administradorID int       `json:"id"`
	nombre          string    `json:"nombre"`
	mail            string    `json:"mail"`
	contrasena      string    `json:"contrasena"`
	rol             string    `json:"rol"`
	fechaCreacion   time.Time `json:"fecha_creacion"`
	ultimoAcceso    time.Time `json:"ultimo_acceso"`
}

// Usuario
type Usuario struct {
	usuarioID  int    `json:"id"`
	nombre     string `json:"nombre"`
	mail       string `json:"mail"`
	contrasena string `json:"contrasena"`
	rol        string `json:"rol"`
}

// Inventario
type Inventario struct {
	inventarioId int  `json:"id"`
	libroID      int  `json:"libro_id"`
	disponible   bool `json:"disponible"`
}

// Libro
type Libro struct {
	libroID          int       `json:"id"`
	titulo           string    `json:"titulo"`
	autor            string    `json:"autor"`
	fechaPublicacion time.Time `json:"fecha_publicacion"`
	genero           string    `json:"genero"`
	url              string    `json:"url"`
}

// Prestamo
type Prestamo struct {
	prestamoID      int       `json:"id"`
	libroID         int       `json:"libro_id"`
	usuarioID       int       `json:"usuario_id"`
	fechaReserva    time.Time `json:"fecha_reserva"`
	fechaDevolucion time.Time `json:"fecha_devolucion"`
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

// Aplicacion de metodos getter para poder acceder a las propiedades que estan encapsuladas

// Administrador
func (a *Administrador) GetNombre() string {
	return a.nombre
}
func (a *Administrador) GetMail() string {
	return a.mail
}
func (a *Administrador) GetRol() string {
	return a.rol
}
func (a *Administrador) GetFechaCreacion() time.Time {
	return a.fechaCreacion
}
func (a *Administrador) GetUltimoAcceso() time.Time {
	return a.ultimoAcceso
}

// Usuario
func (u *Usuario) GetNombre() string {
	return u.nombre
}
func (u *Usuario) GetMail() string {
	return u.mail
}

func (u *Usuario) GetRol() string {
	return u.rol
}

// Inventario
func (i *Inventario) GetInventario() int {
	return i.inventarioId
}
func (i *Inventario) GetLibroID() int {
	return i.libroID
}
func (i *Inventario) IsDisponible() bool {
	return i.disponible
}

// Libro
func (l *Libro) GetTitulo() string {
	return l.titulo
}
func (l *Libro) GetAutor() string {
	return l.autor
}
func (l *Libro) GetFechaPublicacion() time.Time {
	return l.fechaPublicacion
}
func (l *Libro) GetGenero() string {
	return l.genero
}
func (l *Libro) GetURL() string {
	return l.url
}

// Prestamo
func (p *Prestamo) GetLibroID() int {
	return p.libroID
}
func (p *Prestamo) GetUsuarioID() int {
	return p.usuarioID
}
func (p *Prestamo) GetFechaReserva() time.Time {
	return p.fechaReserva
}
func (p *Prestamo) GetFechaDevolucion() time.Time {
	return p.fechaDevolucion
}

// Aplicacion de metodo setter para poder modificar las propiedades que estan encapsuladas

// Adminsitrador
func (a *Administrador) SetNombre(nombre string) {
	a.nombre = nombre
}
func (a *Administrador) SetMail(mail string) {
	a.mail = mail
}
func (a *Administrador) SetRol(rol string) {
	a.rol = rol
}
func (a *Administrador) SetUltimoAcceso(t time.Time) {
	a.ultimoAcceso = t
}

// Usuario
func (u *Usuario) SetNombre(nombre string) {
	u.nombre = nombre
}
func (u *Usuario) SetMail(mail string) {
	u.mail = mail
}
func (u *Usuario) SetRol(rol string) {
	u.rol = rol
}

// Inventario
func (i *Inventario) SetDisponible(disponible bool) {
	i.disponible = disponible
}

// Libro
func (l *Libro) SetTirulo(titulo string) {
	l.titulo = titulo
}
func (l *Libro) SetAutor(autor string) {
	l.autor = autor
}
func (l *Libro) SetFechaPublicacion(fecha time.Time) {
	l.fechaPublicacion = fecha
}
func (l *Libro) SetGenero(genero string) {
	l.genero = genero
}
func (l *Libro) SetURL(url string) {
	l.url = url
}

// Prestamo
func (p *Prestamo) SetFechaDevolucion(fecha time.Time) {
	p.fechaDevolucion = fecha
}

// Manejo de errores en creacion de administradores
func nuevoAdministrador(id int, nombre, mail, contrasena, rol string) (*Administrador, error) {
	if id <= 0 || nombre == "" || mail == "" || contrasena == "" {
		return nil, errors.New("error en los datos para crear un administrador")
	}
	return &Administrador{
		administradorID: id,
		nombre:          nombre,
		mail:            mail,
		contrasena:      contrasena,
		rol:             rol,
		fechaCreacion:   time.Now(),
		ultimoAcceso:    time.Now(),
	}, nil
}

// Manejo de errores en creacion de usuarios
func nuevoUsuario(id int, nombre, mail, contrasena, rol string) (*Usuario, error) {
	if id <= 0 || nombre == "" || mail == "" || contrasena == "" {
		return nil, errors.New("error en los datos para crear un usuario")
	}
	return &Usuario{
		usuarioID:  id,
		nombre:     nombre,
		mail:       mail,
		contrasena: contrasena,
		rol:        rol,
	}, nil
}

// Manejo de errores en creacion de libros
func nuevoLibro(id int, titulo, autor string, fecha time.Time, genero, url string) (*Libro, error) {
	if id <= 0 || titulo == "" || autor == "" {
		return nil, errors.New("error en los datos para crear un libro")
	}
	return &Libro{
		libroID:          id,
		titulo:           titulo,
		autor:            autor,
		fechaPublicacion: fecha,
		genero:           genero,
		url:              url,
	}, nil
}

// Manejo de errores para registro de inventarios
func nuevoInventario(id, libroID int, disponible bool) (*Inventario, error) {
	if id <= 0 || libroID <= 0 {
		return nil, errors.New("error en los datos para registrar en inventario")
	}
	return &Inventario{
		inventarioId: id,
		libroID:      libroID,
		disponible:   disponible,
	}, nil
}

// Manejo de errores para registro de prestamos
func nuevoPrestamo(id, libroID, usuarioID int, fechaReserva, fechaDevolucion time.Time) (*Prestamo, error) {
	if id <= 0 || libroID <= 0 || usuarioID <= 0 {
		return nil, errors.New("error en los datos para registrar un préstamo")
	}
	return &Prestamo{
		prestamoID:      id,
		libroID:         libroID,
		usuarioID:       usuarioID,
		fechaReserva:    fechaReserva,
		fechaDevolucion: fechaDevolucion,
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

// Generamos funciones para paginas: de bienvenida y de despedida
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bienvenido a la Biblioteca de Jaz y Kev")
}
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
			administradorID: 100,
			nombre:          "Kevin Lopez",
			mail:            "kevin.lopez@correo.com",
			contrasena:      "contrasena100",
			rol:             "Administrador",
			fechaCreacion:   time.Now(),
			ultimoAcceso:    time.Now(),
		},
		{
			administradorID: 200,
			nombre:          "Jazmin Chillagana",
			mail:            "jazmin.chillagana@correo.com",
			contrasena:      "contrasena200",
			rol:             "Administrador",
			fechaCreacion:   time.Now(),
			ultimoAcceso:    time.Now(),
		},
	}

	/*Creacion de usuarios
	Utilizamos un slice [] para crear varios usuarios ya que constantemente se puede
	requerir crear mas en el futuro*/

	usuarios := []Usuario{
		{
			usuarioID:  001,
			nombre:     "Juan Perez",
			mail:       "juan.perez@correo.com",
			contrasena: "librosjuan1",
			rol:        "Usuario",
		},
		{
			usuarioID:  002,
			nombre:     "Maria Enriquez",
			mail:       "maria.enriquez@correo.com",
			contrasena: "mislibros123",
			rol:        "Usuario",
		},
		{
			usuarioID:  003,
			nombre:     "Pedro Alvarez",
			mail:       "pedro.alvarez@correo.com",
			contrasena: "miperro5",
			rol:        "Usuario",
		},
		{
			usuarioID:  004,
			nombre:     "Pablo Hernandez",
			mail:       "pablo.hernandez@correo.com",
			contrasena: "contra123",
			rol:        "Usuario",
		},
		{
			usuarioID:  005,
			nombre:     "Samantha Rivera",
			mail:       "samy.rivera@correo.com",
			contrasena: "riosol159",
			rol:        "Usuario",
		},
	}

	/*Creacion de libros
	Utilizamos un slice [] para crear varios libros ya que constantemente se puede
	requerir crear mas en el futuro*/

	libros := []*Libro{
		{
			libroID:          001,
			titulo:           "Cartas de un Estoico",
			autor:            "Lucio A. Séneca",
			fechaPublicacion: time.Date(2024, time.September, 21, 0, 0, 0, 0, time.UTC),
			genero:           "Filosofía",
			url:              "www.libros.com/cartas_estoico",
		},
		{
			libroID:          002,
			titulo:           "Los Discursos de Epicteto",
			autor:            "Epicteto",
			fechaPublicacion: time.Date(2024, time.September, 21, 0, 0, 0, 0, time.UTC),
			genero:           "Filosofía",
			url:              "www.libros.com/discursos_epicteto",
		},
		{
			libroID:          003,
			titulo:           "Manual de Epicteto",
			autor:            "Epicteto",
			fechaPublicacion: time.Date(1980, time.May, 20, 0, 0, 0, 0, time.UTC),
			genero:           "Filosofía",
			url:              "www.libros.com/manual_epicteto",
		},
		{
			libroID:          004,
			titulo:           "Meditaciones",
			autor:            "Marco Aurelio",
			fechaPublicacion: time.Date(2023, time.October, 20, 0, 0, 0, 0, time.UTC),
			genero:           "Filosofía",
			url:              "www.libros.com/meditaciones",
		},
		{
			libroID:          005,
			titulo:           "Sobre la brevedad de la vida",
			autor:            "Lucio A. Séneca",
			fechaPublicacion: time.Date(2024, time.September, 21, 0, 0, 0, 0, time.UTC),
			genero:           "Filosofía",
			url:              "www.libros.com/brevedad_vida",
		},
	}

	/*Creacion de inventario
	Utilizamos un slice [] para crear varios libros ya que constantemente se puede
	requerir crear mas en el futuro*/

	inventario := []*Inventario{
		{
			inventarioId: 001,
			libroID:      libros[0].libroID,
			disponible:   true,
		},
		{
			inventarioId: 002,
			libroID:      libros[1].libroID,
			disponible:   true,
		},
		{
			inventarioId: 003,
			libroID:      libros[2].libroID,
			disponible:   true,
		},
		{
			inventarioId: 004,
			libroID:      libros[3].libroID,
			disponible:   true,
		},
		{
			inventarioId: 005,
			libroID:      libros[4].libroID,
			disponible:   true,
		},
	}

	/*Creacion de prestamos
	Utilizamos un slice [] para crear varios prestamos ya que constantemente se puede
	requerir crear mas en el futuro*/

	prestamos := []*Prestamo{
		{
			prestamoID:      001,
			libroID:         libros[0].libroID,
			usuarioID:       usuarios[0].usuarioID,
			fechaReserva:    time.Now(),
			fechaDevolucion: time.Now().AddDate(0, 0, 5),
		},
		{
			prestamoID:      002,
			libroID:         libros[1].libroID,
			usuarioID:       usuarios[1].usuarioID,
			fechaReserva:    time.Now(),
			fechaDevolucion: time.Now().AddDate(0, 0, 5),
		},
		{
			prestamoID:      003,
			libroID:         libros[2].libroID,
			usuarioID:       usuarios[2].usuarioID,
			fechaReserva:    time.Now(),
			fechaDevolucion: time.Now().AddDate(0, 0, 5),
		},
		{
			prestamoID:      004,
			libroID:         libros[3].libroID,
			usuarioID:       usuarios[3].usuarioID,
			fechaReserva:    time.Now(),
			fechaDevolucion: time.Now().AddDate(0, 0, 5),
		},
		{
			prestamoID:      005,
			libroID:         libros[4].libroID,
			usuarioID:       usuarios[4].usuarioID,
			fechaReserva:    time.Now(),
			fechaDevolucion: time.Now().AddDate(0, 0, 5),
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
	http.HandleFunc("/visualizar-admin", visualizarAdministrador)
	http.HandleFunc("/visualizar-user", visualizarUsuario)
	http.HandleFunc("/visualizar-libro", visualizarLibro)
	http.HandleFunc("/visualizar-inv", visualizarInventario)
	http.HandleFunc("/visualizar-pres", visualizarPrestamos)
	/*http.HandleFunc("/crear-usuario", crearUsuario)
	http.HandleFunc("/crear-libro", crearLibro)
	http.HandleFunc("/registrar-inventario", registrarInventario)
	http.HandleFunc("/solicitar-prestamo", solicitarPrestamo)
	http.HandleFunc("/ver-disponibilidad", verDisponibilidad)*/
	http.HandleFunc("/away", awayPage)

	fmt.Println("Servidor iniciado en el puerto 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

	/*//Imprimir detalles de administradores
	for _, admin := range administradores {
		fmt.Printf("Administrador: %+v\n", admin)
	}

	//Imprimir detalles de usuarios
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
