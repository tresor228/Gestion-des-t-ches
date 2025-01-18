package main

import "fmt"

// Sauvegarde des tâches
func SauvegardeTache(ID int) {
	for i, task := range tasks {
		if task.ID == ID {
			tasks[i].Completed = true
			fmt.Println("Tâche sauvegardé!")
		}
	}
}
