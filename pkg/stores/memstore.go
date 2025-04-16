package stores

import "sync"

type AccountStore struct {
	accounts map[string]int
	mutex sync.Mutex
}

func NewAccountStore() *AccountStore {
	return &AccountStore{
		accounts: make(map[string]int),
	}
}

func (ms *AccountStore) AddBalance(account string, value int) {
	ms.mutex.Lock()
	defer ms.mutex.Unlock()
	ms.accounts[account] += value
}

func (ms *AccountStore) SubtractBalance(account string, value int) {
	ms.mutex.Lock()
	defer ms.mutex.Unlock()
	ms.accounts[account] -= value
}

func (ms *AccountStore) GetBalance(account string) (int, bool) {
	ms.mutex.Lock()
	defer ms.mutex.Unlock()
	value, exists := ms.accounts[account]
	return value, exists
}

func (ms *AccountStore) Reset() {
	ms.mutex.Lock()
	defer ms.mutex.Unlock()
	ms.accounts = make(map[string]int)
}