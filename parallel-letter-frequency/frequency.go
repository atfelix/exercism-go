package letter

import "sync"

type FreqMap map[rune]int

var waitGroup sync.WaitGroup
var mutex sync.Mutex

func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func ConcurrentFrequency(phrases []string) FreqMap {
	m := FreqMap{}
	waitGroup.Add(len(phrases))
	for _, word :=  range phrases {
		go func(word string) {
			m.add(word)
			waitGroup.Done()
		}(word)
	}
	waitGroup.Wait()
	return m
}

func (freqMap FreqMap) add(word string) {
	for _, r := range word {
		go func(r rune) {
			mutex.Lock()
			freqMap.increment(r)
			mutex.Unlock()
		}(r)
	}
}

func (freqMap FreqMap) increment(r rune) {
	freqMap[r]++
}