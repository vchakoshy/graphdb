package graph

import (
	"encoding/gob"
	"errors"
	"fmt"
	"log"
	"os"
	"os/signal"
	"path"
	"sync"
	"syscall"
)

const AppVersion = "0.0.5"

var dumpFile = "dump.db"

// Graph
type Graph struct {
	follow       Follow // store user following
	lock         sync.RWMutex
	metrics      Metrics
	queryOptions QueryOptions
}

func NewGraph() *Graph {
	g := &Graph{
		follow:       NewFollow(),
		metrics:      Metrics{},
		queryOptions: QueryOptionsDefault(),
	}

	g.handleClose()
	err := g.Load(dumpFile)
	if err != nil {
		log.Println("error loading graph:", err)
	}

	return g
}

func (g *Graph) handleClose() {
	var gracefulStop = make(chan os.Signal)
	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)
	go func() {
		sig := <-gracefulStop
		fmt.Printf("caught sig: %+v", sig)
		err := g.Close()
		if err != nil {
			log.Println(err.Error())
		}

		os.Exit(0)
	}()
}

func (g *Graph) Close() error {
	log.Println("Closing Graph")
	log.Println("save data")
	return g.save(dumpFile)
}

func (g *Graph) SetQueryOptions(opts QueryOptions) *Graph {
	g.queryOptions = opts
	return g
}

func (g *Graph) GetMetrics() Metrics {
	m := g.metrics
	m.FollowCount = g.follow.CountAll()
	return m
}

func (g *Graph) GetFollows(from int64) ([]int64, error) {
	g.lock.Lock()
	defer g.lock.Unlock()
	if e, err := g.follow.List(from); err == nil {
		return e, nil
	}
	return []int64{}, fmt.Errorf("user not found: %d", from)
}

func (g *Graph) GetFriendsOfFriends(from int64) ([]int64, error) {
	g.lock.Lock()
	defer g.lock.Unlock()
	return g.follow.FofIds(from, 0, 10), nil
}

func (g *Graph) AddFollow(from, to int64) *Graph {
	if from == to {
		// we should raise error
		return g
	}
	g.lock.Lock()
	defer g.lock.Unlock()
	if g.follow.Exists(from, to) {
		return g
	}
	g.follow.Add(from, to)
	return g
}

func (g *Graph) SuggestByUser(user int64) *Graph {
	return g
}

func (g *Graph) RemoveFollow(from, to int64) *Graph {
	g.lock.Lock()
	g.follow.Remove(from, to)
	g.lock.Unlock()

	return g
}

func (g *Graph) getDumpFilePath(f string) string {
	ed := os.Getenv("DATA_DIR")
	if ed != "" {
		return path.Join(ed, f)
	}
	return path.Join("data", f)
}

func (g *Graph) save(name string) error {
	name = g.getDumpFilePath(name)
	fp, err := os.Create(name)
	if err != nil {
		return err
	}
	enc := gob.NewEncoder(fp)
	defer func() {
		if x := recover(); x != nil {
			err = errors.New("error registering item types with Gob library")
		}
	}()
	g.lock.Lock()
	defer g.lock.Unlock()

	err = enc.Encode(&g.follow)
	if err != nil {
		fp.Close()
		return err
	}
	return fp.Close()
}

func (g *Graph) Load(name string) error {
	name = g.getDumpFilePath(name)
	fp, err := os.Open(name)
	if err != nil {
		return err
	}
	dec := gob.NewDecoder(fp)
	items := g.follow
	err = dec.Decode(&items)
	if err == nil {
		g.lock.Lock()
		defer g.lock.Unlock()
		g.follow = items
	}
	if err != nil {
		fp.Close()
		return err
	}
	return fp.Close()
}
