package handlers

import (
	"bufio"
	"encoding/json"
	"os"
	baseErrors "serv/errors"
	"serv/model"
	"sync"
)

type UserStore struct {
	users  []*model.UserDB
	mu     sync.RWMutex
	nextID uint
}

func NewUserStore() *UserStore {

	return &UserStore{
		mu:    sync.RWMutex{},
		users: []*model.UserDB{},
	}
}

func (us *UserStore) AddUser(in *model.UserDB) (uint, error) {
	us.mu.Lock()
	defer us.mu.Unlock()

	str := model.UserToString(*in)

	outFile := "users.txt"
	var writer *bufio.Writer
	file, err := os.OpenFile(outFile, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return 0, baseErrors.ErrServerError500
	}
	defer file.Close()
	writer = bufio.NewWriter(file)

	writer.WriteString("\n")
	writer.WriteString(str)
	writer.Flush()

	return in.ID, nil
}

func (us *UserStore) GetUsers() ([]*model.UserDB, error) {
	us.mu.RLock()
	defer us.mu.RUnlock()

	users := []*model.UserDB{}
	inpFile := "users.txt"
	file, err := os.Open(inpFile)
	if err != nil {
		return nil, baseErrors.ErrServerError500
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		txt := scanner.Text()
		dat := model.UserDB{}
		err := json.Unmarshal([]byte(txt), &dat)
		if err != nil {
			return nil, baseErrors.ErrServerError500
		}
		users = append(users, &dat)
	}
	file.Close()

	return users, nil
}
