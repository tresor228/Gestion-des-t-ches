package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"time"
)

var tasks []Task

// Ajout des tâches
func AddTask(description, priority, tasksFile string) {
	id := len(tasks) + 1
	task := Task{
		ID:          id,
		Description: description,
		Priority:    priority,
		Completed:   false,
		CreatedAt:   time.Now(),
	}
	tasks = append(tasks, task)
	fmt.Println("Tâche ajoutée :", task.Description)
	SaveTasks(tasks, tasksFile)
}

// Liste des tâches
func ListTasks() {
	if len(tasks) == 0 {
		fmt.Println("Aucune tâche disponible.")
		return
	}

	fmt.Println("Liste des tâches :")
	for _, task := range tasks {
		status := "Non terminée"
		if task.Completed {
			status = "Terminée"
		}
		fmt.Printf("[%d] %s (Priorité : %s, Statut : %s)\n", task.ID, task.Description, task.Priority, status)
	}
}

// Suppression des tâches
func DeleteTask(id int, tasksFile string) {
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			fmt.Println("Tâche supprimée :", task.Description)
			SaveTasks(tasks, tasksFile)
			return
		}
	}
	fmt.Println("Tâche non trouvée avec l'ID :", id)
}

// Sauvegarde des tâches
func SaveTasks(tasks []Task, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(tasks)
}

// Chargement des tâches
func LoadTasks(filename string) ([]Task, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var tasks []Task
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&tasks)
	return tasks, err
}

// Fin des tâches
func CompleteTask(id int, tasksFile string) {
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Completed = true
			fmt.Println("Tâche terminée :", task.Description)
			SaveTasks(tasks, tasksFile)
			return
		}
	}
	fmt.Println("Tâche non trouvée avec l'ID :", id)
}
func main() {
	// Charger les tâches depuis un fichier au démarrage
	tasksFile := "tasks.json"
	var err error
	tasks, err = LoadTasks(tasksFile)
	if err != nil {
		fmt.Println("Aucune tâche trouvée ou erreur de lecture :", err)
		tasks = []Task{}
	}

	// Définir les commandes disponibles
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	addDescription := addCmd.String("desc", "", "Description de la tâche")
	addPriority := addCmd.String("priority", "moyen", "Priorité (petit, moyen, grand)")

	completeCmd := flag.NewFlagSet("complete", flag.ExitOnError)
	completeID := completeCmd.Int("id", -1, "ID de la tâche à marquer comme terminée")

	deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)
	deleteID := deleteCmd.Int("id", -1, "ID de la tâche à supprimer")

	// Vérifier la commande passée
	if len(os.Args) < 2 {
		fmt.Println("Commande manquante : add, list, complete, delete")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "add":
		addCmd.Parse(os.Args[2:])
		if *addDescription == "" {
			fmt.Println("La description est obligatoire pour ajouter une tâche.")
			os.Exit(1)
		}
		AddTask(*addDescription, *addPriority, tasksFile)

	case "list":
		ListTasks()

	case "complete":
		completeCmd.Parse(os.Args[2:])
		if *completeID == -1 {
			fmt.Println("L'ID est obligatoire pour marquer une tâche comme terminée.")
			os.Exit(1)
		}
		CompleteTask(*completeID, tasksFile)

	case "delete":
		deleteCmd.Parse(os.Args[2:])
		if *deleteID == -1 {
			fmt.Println("L'ID est obligatoire pour supprimer une tâche.")
			os.Exit(1)
		}
		DeleteTask(*deleteID, tasksFile)

	default:
		fmt.Println("Commande inconnue :", os.Args[1])
		os.Exit(1)
	}
}
