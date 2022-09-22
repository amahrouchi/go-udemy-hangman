package dictionary

import (
	"bufio"
	"math/rand"
	"os"
	"time"
)

var words = make([]string, 0, 50)

func Load(fileName string) error {
	f, err := os.Open(fileName)
	if err != nil {

	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		return err
	}

	return nil
}

func PickWord() string {
	// Pour utiliser de l'aléatoire en informatique, on utilise toujours une seed qui permet de servir de racine à l'algo aleatoire
	// Si on utilise toujours la mm seed, la sequence aleatoire sera toujours la mm
	// Ici pour rendre le jeu le plus aléatoire possible malgré le systeme incontournable de seed, on utilise le timestamp courant comme seed
	rand.Seed(time.Now().Unix())
	i := rand.Intn(len(words))

	return words[i]
}
