package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type AnimalSpecies int

const (
	SHEEP AnimalSpecies = iota
	CHICKEN
	WOLF
	COW
	LION
	HUNTER
)

type Animal struct {
	Species   AnimalSpecies
	Move      int
	Gender    string
	IsAlive   bool
	xLocation int
	yLocation int
	Range     int
}

func (a *Animal) Process() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	newPosition := r.Intn(4)
	if newPosition == 2 && a.yLocation >= (500-a.Move) {
		newPosition = r.Intn(3)
		if newPosition == 2 {
			newPosition = 3
		}

	} else if newPosition == 3 && a.yLocation <= (a.Move-1) {
		newPosition = r.Intn(3)
	}
	if newPosition == 0 && a.xLocation >= (500-a.Move) {
		newPosition = r.Intn(3)
		if newPosition == 0 {
			newPosition = 3
		}

	} else if newPosition == 1 && a.xLocation <= (a.Move-1) {
		newPosition = r.Intn(3)
		if newPosition == 1 {
			newPosition = 3
		}

	}
	switch newPosition {
	case 0:
		a.xLocation += a.Move
		break
	case 1:
		a.xLocation -= a.Move
		break
	case 2:
		a.yLocation += a.Move
		break
	case 3:
		a.yLocation -= a.Move
		break

	}
}

type Hunter struct {
	Animal
}

func NewHunter() *Hunter {
	return &Hunter{
		Animal{
			IsAlive: true,
			Gender:  "unknown",
			Species: HUNTER,
			Range:   8,
			Move:    1,
		},
	}
}

func main() {
	var Animals []*Animal
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	generateLives(&Animals, SHEEP, 15, "female")
	generateLives(&Animals, SHEEP, 15, "male")
	generateLives(&Animals, COW, 5, "female")
	generateLives(&Animals, COW, 5, "male")
	generateLives(&Animals, CHICKEN, 10, "female")
	generateLives(&Animals, CHICKEN, 10, "male")
	generateLives(&Animals, WOLF, 5, "female")
	generateLives(&Animals, SHEEP, 5, "male")
	generateLives(&Animals, LION, 4, "female")
	generateLives(&Animals, LION, 4, "male")
	hunter := NewHunter()
	hunter.xLocation = random.Intn(500)
	hunter.yLocation = random.Intn(500)
	Animals = append(Animals, &hunter.Animal)

	for i := 0; i < 1000; i++ {

		for _, animal := range Animals {
			if animal.IsAlive {
				animal.Process()
				animalMove(animal, Animals)
			}
		}
	}

	printAnimals(Animals)
}
func printAnimals(Animals []*Animal) {
	livingAnimalCount := 0
	livingAnimalCountBySpecies := make(map[AnimalSpecies]int)

	for _, animal := range Animals {
		if animal.IsAlive {
			livingAnimalCount++
			livingAnimalCountBySpecies[animal.Species]++
		}
	}
	fmt.Printf("%d adet hayvan yaşıyor.\n", livingAnimalCount)
	for species, count := range livingAnimalCountBySpecies {
		switch species {
		case SHEEP:
			fmt.Printf("SHEEP: %d\n", count)
			break
		case COW:
			fmt.Printf("COW: %d\n", count)
			break
		case CHICKEN:
			fmt.Printf("CHICKEN: %d\n", count)
			break
		case WOLF:
			fmt.Printf("WOLF: %d\n", count)
			break
		case LION:
			fmt.Printf("LION: %d\n", count)
			break
		default:
			fmt.Printf("HUNTER: %d\n", count)
			break
		}

	}
}
func generateLive(species AnimalSpecies, animals *[]*Animal) {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	gender := "female"
	if random.Intn(2) == 1 {
		gender = "male"
	}

	generateLives(animals, species, 1, gender)
}
func generateLives(animals *[]*Animal, species AnimalSpecies, count int, gender string) {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < count; i++ {
		var animal *Animal
		switch species {
		case SHEEP:
			animal = &Animal{Species: SHEEP, Move: 2, Range: 0}
			break
		case COW:
			animal = &Animal{Species: COW, Move: 2, Range: 0}
			break
		case CHICKEN:
			animal = &Animal{Species: CHICKEN, Move: 1, Range: 0}
			break
		case WOLF:
			animal = &Animal{Species: WOLF, Move: 3, Range: 4}
			break
		case LION:
			animal = &Animal{Species: LION, Move: 4, Range: 5}
			break
		}

		animal.Gender = gender
		animal.yLocation = random.Intn(500)
		animal.xLocation = random.Intn(500)
		animal.IsAlive = true
		*animals = append(*animals, animal)
	}
}
func findAnimalsWithinDistance(animal *Animal, Animals []*Animal, distanceThreshold float64) []*Animal {
	var nearbyAnimals []*Animal
	for _, a := range Animals {
		if a != animal && a.IsAlive && distance(animal, a) <= distanceThreshold {
			nearbyAnimals = append(nearbyAnimals, a)
		}

	}
	return nearbyAnimals
}

func distance(a, b *Animal) float64 {
	return math.Sqrt(math.Pow(float64(a.xLocation-b.xLocation), 2) + math.Pow(float64(a.yLocation-b.yLocation), 2))
}

func animalMove(a *Animal, Animals []*Animal) {
	switch a.Species {
	case WOLF:
		wolfHunts := findAnimalsWithinDistance(a, Animals, float64(a.Range))
		for _, animal := range wolfHunts {
			if animal.Species == SHEEP || animal.Species == CHICKEN {
				animal.IsAlive = false
			}
		}
		break
	case HUNTER:
		huntedAnimals := findAnimalsWithinDistance(a, Animals, float64(a.Range))
		for _, animal := range huntedAnimals {
			animal.IsAlive = false
		}
		break
	case LION:
		lionHunts := findAnimalsWithinDistance(a, Animals, float64(a.Range))
		for _, animal := range lionHunts {
			if animal.Species == SHEEP || animal.Species == CHICKEN {
				animal.IsAlive = false
			}
		}
		break
	}

	nearAnimals := findAnimalsWithinDistance(a, Animals, 3)
	for _, animal := range nearAnimals {
		if animal.Species == a.Species && animal.Gender != a.Gender {
			generateLive(a.Species, &Animals)
		}
	}
}
