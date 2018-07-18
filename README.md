## Dépôt contenant un projet go/ et un conteneur docker pour le lancer

### context

Se projet s'appel smartincrement et permet de prendre en main le langage go en créant un increment intelligent qui est fonction
du temps écoulé entre plusieurs appel du programme.

### Utilisation

Cet incrément est utilisé pour rendre plus ergonomique le changement de volume sous linux / i3.

Des appuis répétés permettent d'incrémenter par une grande valeur,
tandis que des appuis séparés par plus de 2 secondes permettent un réglage plus fin.

### Execution

Un example d'utilisation qui requière `pactl` pour la gestion du son est présent dans le dossier `./demo`

Le binding i3 à placer dans ~/.config/i3/config est le suivant :

```
bindcode 122 exec --no-startup-id "/bin/perso/soundManager.sh down"
bindcode 123 exec --no-startup-id "/bin/perso/soundManager.sh up"
```


