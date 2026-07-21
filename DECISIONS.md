# Décisions techniques

## SQLite

Vitalynq utilisera SQLite pour stocker les données localement.

Raisons :

- stockage local adapté à une application CLI mono-utilisateur ;
- fichier unique facile à sauvegarder ;
- pas de serveur à administrer ;
- bon support des transactions.

Le code utilisera `database/sql` pour rester proche de la bibliothèque standard.

Le driver SQLite prévu est `modernc.org/sqlite`, car il fonctionne sans CGO et simplifie l'installation locale.

Cette décision pourra être réévaluée si une contrainte technique concrète apparaît.