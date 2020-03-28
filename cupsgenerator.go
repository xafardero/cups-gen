package cupsgenerator

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

var controlNumbers = []string{"T", "R", "W", "A", "G", "M", "Y", "F", "P", "D", "X", "B", "N", "J", "Z", "S", "Q", "V", "H", "L", "C", "K", "E"}

func genReeID() (int, string) {
	rand.Seed(time.Now().UnixNano())
	random := rand.Intn(9999)

	randomString := strconv.Itoa(random)

	if len(randomString) < 4 {
		repeat := 4 - len(randomString)
		randomString = strings.Repeat("0", repeat) + randomString
	}

	return random, randomString
}

func genDistID() (int, string) {
	rand.Seed(time.Now().UnixNano())
	random := rand.Intn(999999999999)

	randomString := strconv.Itoa(random)

	if len(randomString) < 12 {
		repeat := 12 - len(randomString)
		randomString = strings.Repeat("0", repeat) + randomString
	}

	return random, randomString
}

func Generate() string {
	reeID, reeIDString := genReeID()
	distID, distIDString := genDistID()

	output := fmt.Sprintf("%d%d", reeID, distID)
	i, _ := strconv.Atoi(output)

	control := i % 529
	division := control / 23
	mod := control % 23

	cups := fmt.Sprintf("%s%s%s%s%s", "ES", reeIDString, distIDString, controlNumbers[division], controlNumbers[mod])

	return cups
}
