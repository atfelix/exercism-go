package tournament

import (
	"fmt"
	"errors"
	"io"
	"io/ioutil"
	"sort"
	"strings"
)

type game struct {
	firstTeam string
	secondTeam string
	result string
}

func gameFrom(s string) game {
	list := strings.Split(s, ";")
	if len(list) != 3 {
		return game{}
	}
	return game{
		firstTeam: list[0],
		secondTeam: list[1],
		result: list[2],
	}
}

type team struct {
	name string
	wins, losses, draws uint
}

func (t team) String() string {
	return fmt.Sprintf("%-31s| %2d | %2d | %2d | %2d | %2d\n", t.name, t.matchesPlayed(), t.wins, t.draws, t.losses, t.points())
}

func (t *team) matchesPlayed() uint {
	return t.wins + t.draws + t.losses
}

func (t *team) points() uint {
	return 3 * t.wins + t.draws
}

type tournament struct {
	games []game
	standings map[string]*team
}

func New() *tournament {
	return &tournament{
		games: []game{},
		standings: map[string]*team{},
	}
}

func (tourney *tournament) addGame(g game) error {
	tourney.games = append(tourney.games, g)

	firstTeam, ok := tourney.standings[g.firstTeam]

	if !ok {
		tourney.standings[g.firstTeam] = &team{ name: g.firstTeam }
		firstTeam = tourney.standings[g.firstTeam]
	}

	secondTeam, ok := tourney.standings[g.secondTeam]

	if !ok {
		tourney.standings[g.secondTeam] = &team{ name: g.secondTeam }
		secondTeam = tourney.standings[g.secondTeam]
	}

	if g.result == "win" {
		firstTeam.wins++
		secondTeam.losses++
	} else if g.result == "loss" {
		firstTeam.losses++
		secondTeam.wins++
	} else if g.result == "draw" {
		firstTeam.draws++
		secondTeam.draws++
	} else {
		return errors.New("Invalid match result")
	}
	return nil
}

func (tourney tournament) table() []team {
	teams := []team{}

	for _, t := range tourney.standings {
		teams = append(teams, *t)
	}

	sort.Slice(teams, func(i, j int) bool {
		return (teams[i].points() > teams[j].points() ||
				(teams[i].points() == teams[j].points() && teams[i].name < teams[j].name))
	})

	return teams
}

func Tally(reader io.Reader, writer io.Writer) (err error) {
	bytes, _ := ioutil.ReadAll(reader)
	tourney := New()
	lines := strings.Split(string(bytes), "\n")
	for _, line := range lines {
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		if err := tourney.addGame(gameFrom(line)); err != nil {
			return err
		}
	}
	writer.Write([]byte(fmt.Sprintf("%-31s| %2s | %2s | %2s | %2s | %2s\n", "Team", "MP", "W", "D", "L", "P")))
	teamsSorted := tourney.table()
	for _, t := range teamsSorted {
		writer.Write([]byte(t.String()))
	}
	return nil
}