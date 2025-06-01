package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	models "github.com/patogihub/pryctgo/todo-api/modules"
	storage "github.com/patogihub/pryctgo/todo-api/storages"
)

/*
CreateTask(c *gin.Context)

GetTasks(c *gin.Context)

UpdateTask(c *gin.Context)

DeleteTask(c *gin.Context)
*/

func GetTasks(c *gin.Context) {
	// TODO: devolver lista de tareas

	tasks := storage.Tasks
	c.JSON(http.StatusOK, gin.H{"message:": tasks})

}

func CreateTask(c *gin.Context) {

	// TODO: crear tarea desde JSON -> se le pasa un json y arma la tarea el json tiene 3 datos id(int) titulo(string) y estado(bool)
	c.JSON(http.StatusCreated, gin.H{"message": "Crear tarea"})
	var newTask models.Task

	// Vinculamos el cuerpo del request al struct Task
	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
		return
	}

	storage.LastID++
	newTask.ID = storage.LastID
	newTask.Completed = false // por defecto, nueva tarea no está completada
	// agregamos al final en la lista
	storage.Tasks = append(storage.Tasks, newTask)

	// Devolvemos la tarea creada
	c.JSON(http.StatusCreated, newTask)
}

func UpdateTask(c *gin.Context) {
	// TODO: actualizar tarea por toma un id modifica esa tarea o su estado
	c.JSON(http.StatusOK, gin.H{"message": "Actualizar tarea"})

	idp := c.Param("id") //-> esto lo saca desde el url
	var id int
	if _, err := fmt.Sscanf(idp, "%d", &id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var updateData models.Task
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
		return
	}

	for i, task := range storage.Tasks {
		if task.ID == id {
			// Actualiza los campos si vienen en el JSON
			if updateData.Title != "" {
				storage.Tasks[i].Title = updateData.Title
			}
			storage.Tasks[i].Completed = updateData.Completed
			c.JSON(http.StatusOK, storage.Tasks[i])
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Tarea no encontrada"})

}

func DeleteTask(c *gin.Context) {
	// TODO: eliminar tarea por toma un id elimina esa tarea
	c.JSON(http.StatusOK, gin.H{"message": "Eliminar tarea"})

	idp := c.Param("id") //-> esto lo saca desde el url
	var id int
	if _, err := fmt.Sscanf(idp, "%d", &id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	for i, task := range storage.Tasks {
		if task.ID == id {
			storage.Tasks = append(storage.Tasks[:i], storage.Tasks[i+1:]...) //""...""  hace referencia a que se repite hasta el final del arreglo
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Tarea no encontrada"})
}
