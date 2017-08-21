package main

import (
	"fmt"

	"github.com/abadojack/whatlanggo"
	"github.com/jdkato/prose/summarize"
)

func main() {

	text := `Chaque génération doit être prête à combattre
	 le type de sectarisme et de haine qui a été affiché par les
	  suprémacistes blancs à Charlottesville. » Sheryl Sandberg, la numéro
	   deux de Facebook, n’a pas mâché pas ses mots dans un message, 
	   publié lundi 14 août sur le réseau social.
	Samedi, une femme a été tuée par un automobiliste qui a 
	foncé dans une foule de manifestants venus protester contre un rassemblement de militants d’extrême droite à Charlottesville, en Virginie. La tenue de ce rassemblement, la mort de cette femme et la condamnation timide de Donald Trump, qui a immédiatement suivi ont généré une polémique aux Etats-Unis. Les entreprises de la Silicon Valley, qui se montrent de plus en plus politisées, ont elles aussi réagi de différentes manières.`
	info := whatlanggo.Detect(text)

	fmt.Println("Language:", whatlanggo.LangToString(info.Lang), "Script:", whatlanggo.Scripts[info.Script])

	doc := summarize.NewDocument(text)
	fmt.Println(doc.SMOG(), doc.FleschKincaid())
	fmt.Printf("%v", doc.Summary(1))

}
