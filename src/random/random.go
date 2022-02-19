package random

import (
	"math/rand"
	"sync"
)

const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
const charlen = 62

type Rnd struct {
	lock   sync.Mutex
	random *rand.Rand
}

func New(seed int64) *Rnd {
	return &Rnd{random: rand.New(rand.NewSource(seed))}
}

func (rnd *Rnd) String(size int) string {
	rnd.lock.Lock()
	defer rnd.lock.Unlock()

	result := make([]byte, 0, size)

	for i := 0; i < size; i++ {
		result = append(result, chars[rnd.random.Int31n(charlen)])
	}

	return string(result)
}
