# Sistema de Gestión de Libros

Este es un sistema para gestionar libros, usuarios, inventarios y préstamos en una biblioteca. Permite realizar operaciones como modificar, visualizar y administrar estos datos de manera eficiente.

## Autores
- Jazmin Chillagana
- Kevin López

## Fecha de Creación
**11/18/2024**

## Fecha de Presentación
**12/18/2024**

---

## Objetivo
Facilitar la administración de bibliotecas mediante funcionalidades como:
- Gestión de libros, usuarios, inventarios y préstamos.
- Modificación y visualización dinámica de información.

---

## Estructuras
El sistema utiliza las siguientes estructuras para representar los componentes de una biblioteca:

- **Administrador**: Representa a los administradores del sistema.
- **Usuario**: Personas que pueden solicitar préstamos de libros.
- **Inventario**: Representa el inventario de libros disponibles.
- **Libro**: Contiene la información de los libros.
- **Préstamo**: Representa los préstamos realizados por los usuarios.

---

## Interfaces
Se definieron las siguientes interfaces para manejar funcionalidades clave:

- **Permisos**: Modificar e ingresar información.
- **Búsqueda**: Realizar búsquedas por ID y nombre.
- **Serialización**: Manejar la serialización y deserialización de datos en formatos JSON.

---

## Funcionalidades
### Creación Dinámica de Datos
Se utilizan slices ([]) para permitir la creación dinámica de datos para administradores, usuarios, libros, inventarios y préstamos, facilitando la adición de nuevos registros en el futuro.

### Métodos Getter y Setter
Permiten acceder y modificar las propiedades encapsuladas de las estructuras.

### Manejo de Errores
Incluye funciones para manejar errores durante la creación de administradores, usuarios, libros, inventarios y préstamos.

### Manejo de Archivos JSON
Funciones:
- saveToJSON: Guarda los datos estructurados en archivos JSON.
- loadFromJSON: Carga los datos desde archivos JSON, con manejo de errores en caso de fallas.

### Visualización de Datos
Se utiliza HTML para visualizar los datos mediante un servidor web.

---

## Generación del Servicio Web
El servidor web se ejecuta en el puerto 8080 y proporciona las siguientes rutas:

- **Página de inicio (/)**  
  - **Función**: homePage  
  - Sirve como página de bienvenida y utiliza una plantilla HTML para mostrar contenido.

- **Crear Administrador (/crear-admin)**  
  - Endpoint registrado (sin implementación en el código compartido).

- **Crear Usuario (/crear-user)**  
  - Endpoint registrado (sin implementación en el código compartido).

- **Crear Libro (/crear-book)**  
  - Endpoint registrado (sin implementación en el código compartido).

- **Visualizar Administradores (/visualizar-admin)**  
  - **Función**: visualizarAdministrador  
  - Carga los datos desde el archivo administradores.json y los devuelve en formato JSON.

- **Visualizar Usuarios (/visualizar-user)**  
  - **Función**: visualizarUsuario  
  - Carga los datos desde el archivo usuarios.json y los devuelve en formato JSON.

- **Visualizar Libros (/visualizar-libro)**  
  - **Función**: visualizarLibro  
  - Carga los datos desde el archivo libros.json y los devuelve en formato JSON.

- **Visualizar Inventario (/visualizar-inv)**  
  - **Función**: visualizarInventario
  - Carga los datos desde el archivo inventario.json y los devuelve en formato JSON.

- **Visualizar Préstamos (/visualizar-pres)**  
  - **Función**: visualizarPrestamos  
  - Carga los datos desde el archivo prestamos.json y los devuelve en formato JSON.

- **Buscar Libro (/buscar-libro)**  
  - **Función**: buscarLibro  
  - Permite buscar un libro por ID a través de un formulario HTML.  
  - Retorna los detalles del libro encontrado en formato JSON o un error si no se encuentra.

- **Página de Despedida (/away)**  
  - **Función**: awayPage 
  - Devuelve un mensaje simple de agradecimiento por visitar la biblioteca.

### Endpoints Comentados (No Implementados)
- /registrar-inventario: Probablemente para registrar nuevos inventarios.
- /solicitar-prestamo: Probablemente para registrar nuevos préstamos.
- /ver-disponibilidad: Probablemente para verificar la disponibilidad de libros.

---

## Ejecución del Servidor
Para ejecutar el servidor, usa el siguiente comando:
```bash
go run Sistema_Gestion_Libros.go
