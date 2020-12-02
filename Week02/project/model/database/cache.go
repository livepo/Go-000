package database

import (
    "project/err"
    "sync"
)

type Memory struct {
    Cache map[string]string
    Lock sync.Mutex
}

func NewMemory() *Memory {
    return &Memory{
        Cache: make(map[string]string),
        Lock: sync.Mutex{},
    }
}

func (m *Memory) Put(key string, value string) error {
    m.Lock.Lock()
    defer m.Lock.Unlock()
    if _, ok := m.Cache[key]; ok {
        return err.Duplicate
    }
    m.Cache[key] = value
    return nil
}

func (m *Memory) Get(key string) (string, error) {
    m.Lock.Lock()
    defer m.Lock.Unlock()
    if value, ok := m.Cache[key]; ok {
        return value, nil
    }
    return "", err.NotExist
}