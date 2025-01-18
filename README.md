# Gestion des Tâches

## Description

Ce projet est une application simple de gestion des tâches écrite en **Go (Golang)**. Elle permet de gérer des tâches en effectuant les opérations suivantes :
- Ajouter une tâche
- Lister toutes les tâches
- Marquer une tâche comme terminée
- Supprimer une tâche
- Sauvegarder et charger les données dans un fichier JSON

Le projet est conçu pour être une démonstration pratique des bases de la programmation en Go, notamment la gestion des fichiers, l'utilisation des structures, et les interfaces utilisateur en ligne de commande.

---

## Fonctionnalités

1. **Ajouter une tâche** : Permet d'ajouter une nouvelle tâche avec un titre et une description.
2. **Lister les tâches** : Affiche toutes les tâches existantes avec leur statut.
3. **Mettre à jour une tâche** : Modifier le statut d'une tâche (par exemple, la marquer comme terminée).
4. **Supprimer une tâche** : Permet de retirer une tâche spécifique de la liste.
5. **Sauvegarder les données** : Les tâches sont sauvegardées dans un fichier JSON pour persistance.
6. **Charger les données** : Les tâches existantes sont automatiquement chargées au démarrage.

---

## Structure du projet

```
task-manager/
├── main.go            # Point d'entrée de l'application
├── task.go            # Définitions des structures et méthodes pour les tâches
├── storage.go         # Gestion des fichiers JSON (sauvegarde et chargement)
├── README.md          # Documentation du projet
├── tasks.json         # Fichier de stockage des données des tâches
└── go.mod             # Gestion des dépendances et modules
```

---

## Installation

1. Assurez-vous que **Go** est installé sur votre machine. Vous pouvez le télécharger ici : [https://go.dev/](https://go.dev/)
2. Clonez ce dépôt sur votre machine locale :
   ```bash
   git clone https://github.com/tresor228/Mini-Gestion-des-T-ches.git
   ```
3. Accédez au répertoire du projet :
   ```bash
   cd Mini-Gestion-des-T-ches
   ```
4. Exécutez le programme :
   ```bash
   go run main.go
   ```

---

## Utilisation

Voici un aperçu des commandes disponibles dans l'application :

1. **Ajouter une tâche** :
   L'utilisateur entre le titre et la description d'une nouvelle tâche.
2. **Lister les tâches** :
   Toutes les tâches sont affichées, avec leurs détails.
3. **Mettre à jour une tâche** :
   L'utilisateur peut modifier le statut d'une tâche en fournissant son ID.
4. **Supprimer une tâche** :
   L'utilisateur peut supprimer une tâche spécifique en fonction de son ID.
---

## Contributeurs

- **Kodjo B. Trésor ALADE**  
  Étudiant en informatique, passionné par le développement d'applications et l'intelligence artificielle.  
  GitHub : [Tresor228](https://github.com/tresor228)

---

## Licence

Ce projet est sous licence MIT. Vous êtes libre de l'utiliser, de le modifier et de le distribuer, sous réserve de mentionner l'auteur original.
