package main

import (
	"fmt"
	"github.com/abadojack/whatlanggo"
)

func main() {
	info := whatlanggo.Detect(`Chaque génération doit être prête à combattre le type de sectarisme et de haine qui a été affiché par les suprémacistes blancs à Charlottesville. » Sheryl Sandberg, la numéro deux de Facebook, n’a pas mâché pas ses mots dans un message, publié lundi 14 août sur le réseau social.
	Samedi, une femme a été tuée par un automobiliste qui a foncé dans une foule de manifestants venus protester contre un rassemblement de militants d’extrême droite à Charlottesville, en Virginie. La tenue de ce rassemblement, la mort de cette femme et la condamnation timide de Donald Trump, qui a immédiatement suivi ont généré une polémique aux Etats-Unis. Les entreprises de la Silicon Valley, qui se montrent de plus en plus politisées, ont elles aussi réagi de différentes manières.
	En savoir plus sur http://www.lemonde.fr/pixels/article/2017/08/15/messages-blocages-et-demission-apres-charlottesville-la-silicon-valley-reagit_5172607_4408996.html#vmyCHY7SqCZddysI.99`)
	
	fmt.Println("Language:", whatlanggo.LangToString(info.Lang), "Script:", whatlanggo.Scripts[info.Script])
}
