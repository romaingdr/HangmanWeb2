# Jeu du Pendu web

Ce projet du Jeu du Pendu sur le Web est une implémentation simple du célèbre jeu du Pendu en utilisant Go pour le backend et des templates HTML pour le frontend.

# Fonctionnalités

- Jouabilité: Devinez des lettres pour découvrir un mot caché en un nombre limité d'essais.
- Niveaux de Difficulté: Choisissez parmi des ensembles de mots faciles, moyens ou difficiles.
- Interface Web: Interagissez avec le jeu via un navigateur web.


# Installation et Utilisation

- Prérequis: Langage de programmation Go.
# Installation:
Clonez ou téléchargez le projet.

# Exécution:
Accédez au répertoire du projet.
Exécutez le fichier principal Go avec "go run main.go"

# Accès:

Ouvrez un navigateur web et allez sur http://localhost:8080/accueil.
Choisissez un niveau de difficulté et commencez à jouer !

# Structure des Fichiers
- main.go: Contient la logique principale pour exécuter le serveur de jeu.
- templates: Répertoire contenant des templates HTML pour différentes interfaces de jeu.
- assets: Répertoire contenant des ensembles de mots pour différents niveaux de difficulté.

# Dépendances
- net/http: Pour gérer les requêtes et réponses HTTP.
- html/template: Pour le rendu des templates HTML.
- bufio, os: Pour la manipulation de fichiers et la lecture des ensembles de mots.
- strings, math/rand, time: Pour diverses opérations sur les chaînes et les nombres aléatoires.

# Contributeurs
Goudergues Romain @romaingdr
Carre rayan @rayancarre