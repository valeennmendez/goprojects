package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

// Se crea una funcion para conectarse a la base de datos.
func ConectionDB() *sql.DB {
	Driver := "mysql"
	Usuario := "root"
	Contraseña := ""
	Nombre := "datospacientes"

	conectionDB, err := sql.Open(Driver, Usuario+":"+Contraseña+"@tcp(127.0.0.1)/"+Nombre)

	if err != nil {
		panic(err.Error())
	}

	return conectionDB
}

var router = gin.Default() //Se crea una variable que contiene un enturador de GIN.

type Paciente struct {
	Id     int
	Nombre string
	Correo string
}

func Index(c *gin.Context) {

	conection := ConectionDB()

	registros, err := conection.Query("SELECT * FROM pacientes")

	if err != nil {
		panic(err.Error())
	}

	paciente := Paciente{}
	arrayPacientes := []Paciente{}

	for registros.Next() {
		var id int
		var nombre, correo string

		err := registros.Scan(&id, &nombre, &correo)

		if err != nil {
			panic(err.Error())
		}

		paciente.Id = id
		paciente.Nombre = nombre
		paciente.Correo = correo

		arrayPacientes = append(arrayPacientes, paciente)
	}

	fmt.Println(arrayPacientes)

	c.HTML(http.StatusOK, "index", arrayPacientes)

}

func Agregar(c *gin.Context) {
	//Se carga el template del formulario para el ingreso de pacientes.
	c.HTML(http.StatusOK, "insert", nil)
}

func Insert(c *gin.Context) {

	conection := ConectionDB()

	if c.Request.Method == "POST" {
		id := c.Request.FormValue("id")
		nombre := c.Request.FormValue("nombre")
		correo := c.Request.FormValue("correo")

		registro, err := conection.Prepare("INSERT INTO pacientes(id,nombre,correo) VALUES (?,?,?)")

		if err != nil {
			panic(err.Error())
		}

		registro.Exec(id, nombre, correo)
	}

	c.Redirect(http.StatusFound, "/index")
}

func Delete(c *gin.Context) {
	conection := ConectionDB()

	idPaciente := c.Request.URL.Query().Get("id")
	fmt.Println(idPaciente)

	//Se borra la fila de base de datos del id ingresado.
	deleteRegistros, err := conection.Prepare("DELETE FROM pacientes WHERE id=?")

	if err != nil {
		panic(err.Error())
	}

	//Ejecuta la sentencia de codigo MySql
	deleteRegistros.Exec(idPaciente)
	c.Redirect(http.StatusFound, "/index")
}

func Edit(c *gin.Context) {
	conection := ConectionDB()
	idPaciente := c.Request.URL.Query().Get("id")
	fmt.Println(idPaciente)

	editRegistro, err := conection.Query("SELECT * FROM pacientes WHERE id=?", idPaciente)

	if err != nil {
		panic(err.Error())
	}

	pacientes := Paciente{}

	for editRegistro.Next() {
		var id int
		var nombre, correo string

		err := editRegistro.Scan(&id, &nombre, &correo)

		if err != nil {
			panic(err.Error())
		}

		pacientes.Id = id
		pacientes.Nombre = nombre
		pacientes.Correo = correo

	}

	c.HTML(http.StatusOK, "edit", pacientes)
}

func Actualizar(c *gin.Context) {
	conection := ConectionDB()

	if c.Request.Method == "POST" {
		id := c.Request.FormValue("id")
		nombre := c.Request.FormValue("nombre")
		correo := c.Request.FormValue("correo")

		fmt.Println(id)
		fmt.Println(nombre)
		fmt.Println(correo)

		modificarRegistros, err := conection.Prepare("UPDATE pacientes SET nombre=?,correo=? WHERE id=?")

		if err != nil {
			panic(err.Error())
		}

		modificarRegistros.Exec(nombre, correo, id)

		c.Redirect(http.StatusFound, "/index")

	}
}

func LoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login", nil)
}

// Se crea una estructura Empleado para almacenar los datos
// encontrados en la base de datos.
type Empleado struct {
	ID       int
	Name     string
	LastName string
	Email    string
	Password string
}

func Login(c *gin.Context) {
	conexion := ConectionDB()

	if c.Request.Method == "POST" {
		//Se toma y se guardan en variables los datos ingresados en los
		//campos (inputs) del formulario.
		email := c.Request.FormValue("email")
		password := c.Request.FormValue("password")

		if strings.Trim(email, " ") == "" || strings.Trim(password, " ") == "" {
			fmt.Println("Los campos no pueden estar vacios.")
			c.HTML(http.StatusBadRequest, "login", gin.H{"error": "Los campos no pueden estar vacios."})
			return
		}

		//Selecciona la fila del Email ingresado en el formulario.
		registros, err := conexion.Query("SELECT name, lastname, email, password, id FROM empleados WHERE email=?", email)
		if err != nil {
			panic(err.Error())
		}
		defer registros.Close()

		//Se crea una variable con los campos de la estructura Empleado
		//para almacenar los datos.
		empleados := Empleado{}
		loginError := false

		for registros.Next() {
			var id int
			var name, lastname, email, password string

			//Se guardan en las variables los datos de la base de datos.
			err = registros.Scan(&name, &lastname, &email, &password, &id)
			if err != nil {
				panic(err.Error())
			}

			//Se almacenan los datos del empleado.
			empleados.Name = name
			empleados.LastName = lastname
			empleados.Email = email
			empleados.Password = password
			empleados.ID = id
		}

		if empleados.Email == "" {
			loginError = true
		} else {
			//Compara el hash que se encuentra en la base de datos con el hash ingresado en el login.
			errf := bcrypt.CompareHashAndPassword([]byte(empleados.Password), []byte(password))
			if errf != nil {
				loginError = true
			}
		}

		if loginError {
			fmt.Println("Credenciales inválidas.")
			c.HTML(http.StatusUnauthorized, "login", gin.H{"error": "¡Credenciales inválidas!"})
		} else {
			// Login exitoso, se redirigie a la página index.
			c.Redirect(http.StatusFound, "/index")
		}
	} else {
		c.HTML(http.StatusOK, "login", nil)
	}
}

func RegisterPage(c *gin.Context) {
	//Se carga el template del formulario de registro de empleados.
	c.HTML(http.StatusOK, "register", nil)
}

func Register(c *gin.Context) {
	conexion := ConectionDB()

	if c.Request.Method == "POST" {
		//Se toman y almacenan en variables los datos ingresados en los
		//campos(inputs) del formulario.
		name := c.Request.FormValue("name")
		lastname := c.Request.FormValue("lastname")
		email := c.Request.FormValue("email")
		password := c.Request.FormValue("password")
		confirmpassword := c.Request.FormValue("confirmpassword")

		if name == "" && lastname == "" && email == "" && password == "" && confirmpassword == "" {
			c.HTML(http.StatusBadRequest, "register", gin.H{"error": "Los campos NO pueden estar vacíos."})
			return
		}

		//La contraseña ingresada en el input password debe ser igual a la de confirmpassword
		//en caso contrario se imprime un error en la pagina.
		if password != confirmpassword {
			c.HTML(http.StatusBadRequest, "register", gin.H{"error": "Las contraseñas NO coinciden."})
			return
		}

		//Se hashea la contraseña ingresada.
		hashPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

		if err != nil {
			panic(err.Error())
		}

		//Creacion del formato de la fecha de registro.
		date := time.Now()
		createDate := date.Format("2006-01-02 15:04:05")

		register, err := conexion.Prepare("INSERT INTO empleados(name,lastname,email,password,date) VALUE (?,?,?,?,?)")

		if err != nil {
			panic(err.Error())
		}

		register.Exec(name, lastname, email, hashPass, createDate)
	}

	c.Redirect(http.StatusMovedPermanently, "/")

}

func main() {

	//Handlers.
	router.GET("/agregar", Agregar)
	router.POST("/insert", Insert)
	router.GET("/delete", Delete)
	router.GET("/edit", Edit)
	router.POST("/actualizar", Actualizar)

	router.GET("/", LoginPage)
	router.POST("/login", Login)
	router.GET("/index", Index)

	router.GET("/registerpage", RegisterPage)
	router.POST("/register", Register)

	router.SetHTMLTemplate(template.Must(template.ParseGlob("templates/*")))
	router.Static("/static", "./static")

	fmt.Println("Servidor corriendo...")
	router.Run(":8080")
}
