package main

import "fmt"
import "strconv"
import "math/rand"

func main() {
fmt.Print("囚犯 -Bonjour Yin, vous êtes prisonnier et vous allez devoir vous liberez face à Yang- 囚犯 \n")
var yin,yang,nombrePasPorte int
var joueur string
fmt.Print("A chaque tour je vous dirai la case sur laquelle vous êtes, gagne celui qui se trouve sur la case -1 car il a franchis la porte \n")
fmt.Print("好运 -Bonne chance- 好运 \n")
fmt.Print("Vous pouvez sélectionner le nombre de pas qui vous sépares de la porte \n")
fmt.Scan(&nombrePasPorte)
    for nombrePasPorte < 16 {
    fmt.Print("Veuillez saisir le nombre de pas par rapport à la porte supérieur à 15 et inférieur à 100 :\n")
    fmt.Scan(&nombrePasPorte)
    }
    for nombrePasPorte > 99 {
    fmt.Print("Veuillez saisir le nombre de pas par rapport à la porte supérieur à 15 et inférieur à 100 : \n")
    fmt.Scan(&nombrePasPorte)
    }
fmt.Print("Yang, votre adversaire va sélectionner le nombre de pas maximum \n")
var nombrePasMaximum int
nombrePasMaximum = rand.Intn(8)
if nombrePasMaximum < 4{
    nombrePasMaximum = rand.Intn(8)
}
fmt.Print("Le nombre de pas maximum choisis par votre adversaire est :",nombrePasMaximum," pas \n")
fmt.Print("Yin vous avez le privilège de commencer cette partie \n")
for nombrePasPorte >= -1 {
    if nombrePasPorte > -1 {
        fmt.Print("C'est à vous de jouer \n")
        fmt.Print("Choisissez le nombre de case entre 1 et "+ strconv.Itoa(nombrePasMaximum) + " : \n")
        fmt.Scan(&yin)
        joueur="yin"
        for yin < 1 {
            fmt.Print("Choisissez le nombre de case entre 1 et "+strconv.Itoa(nombrePasMaximum) + " : \n")
            fmt.Scan(&yin)
        }
        for yin > nombrePasMaximum {
            fmt.Print("Choisissez le nombre de case entre 1 et "+strconv.Itoa(nombrePasMaximum) + " : \n")
            fmt.Scan(&yin)
        }
        nombrePasPorte=nombrePasPorte-yin
        fmt.Print("Vous vous trouvez désormais à la case ",nombrePasPorte, "\n")
    }
    if nombrePasPorte > -1 {
        fmt.Print("c'est désormais à yang de jouer \n")
        joueur="yang"
        yang = rand.Intn(nombrePasMaximum)
        if yang <= 0{
            yang = rand.Intn(nombrePasMaximum)
        }
        fmt.Print("Yang avance de ", yang, " pas \n")
        nombrePasPorte=nombrePasPorte-yang
        fmt.Print("Vous vous trouvez désormais à la case ",nombrePasPorte,"\n")
    }
    if nombrePasPorte <= -1 {
        fmt.Print("la porte est franchis, la partie est terminé,", joueur, " est le grand gagnant \n")
        break
         }
}
}