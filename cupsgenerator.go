package cupsgenerator

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

var controlNumbers = []string{"T", "R", "W", "A", "G", "M", "Y", "F", "P", "D", "X", "B", "N", "J", "Z", "S", "Q", "V", "H", "L", "C", "K", "E"}

func genReeId() int {
	rand.Seed(time.Now().UnixNano())
	random := rand.Intn(9999)
	return random
}

func genDistId() int {
	rand.Seed(time.Now().UnixNano())
	random := rand.Intn(999999999999)
	return random
}

func Generate() string {
	reeId := genReeId()
	distId := genDistId()

	output := fmt.Sprintf("%d%d", reeId, distId)
	i, _ := strconv.Atoi(output)

	control := i % 529
	division := control / 23
	mod := control % 23

	cups := fmt.Sprintf("%s%d%d%s%s", "ES", reeId, distId, controlNumbers[division], controlNumbers[mod])

	return cups
}
