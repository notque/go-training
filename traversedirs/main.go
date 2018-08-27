package main

import (
	"strings"
	"fmt"
	"os"
	"path/filepath"
	"io/ioutil"
	"encoding/json"
	//"github.com/davecgh/go-spew/spew"
	"sort"
)

func run(a *AllRuns) (*AllRuns) {
	searchDir := "C:/slay"

	fileList := make([]string, 0)
	e := filepath.Walk(searchDir, func(path string, f os.FileInfo, err error) error {
		fileList = append(fileList, path)
		return err
	})
	
	if e != nil {
		panic(e)
	}

	for _, file := range fileList {
		//fmt.Println(file)
		file, e := ioutil.ReadFile(file)
    	if e != nil {
        	fmt.Printf("File error: %v\n", e)
        	//os.Exit(1)
    	}
    	//fmt.Printf("%s\n", string(file))
		var c STSRun
   	 	err := json.Unmarshal(file, &c)
		if err != nil {
			fmt.Println(err)
		}	
		analyze(c, a)
		//spew.Dump(c)
	}

	return a
}

func analyze(c STSRun, a *AllRuns) {
	if c.IsDaily == false && c.IsAscensionMode == true {
		fmt.Printf("Analyzing Run: Reached Floor: %d\n", c.FloorReached )
		fmt.Printf("Killed By: %s\n", c.KilledBy)
		fmt.Printf("Purchased %v\n", c.ItemsPurchased)
		if c.Victory == true {
			fmt.Printf("Won game\n")
			for _, elem := range c.ItemsPurchased {
				a.ItemsPurchasedW[elem]++
			}
			for _, elem := range c.MasterDeck {
				replaced := strings.Replace(elem, "+1", "", -1)
				a.MasterDeckW[replaced]++
			}
		} else {
			fmt.Printf("Lost game\n")
			for _, elem := range c.ItemsPurchased {
				a.ItemsPurchasedL[elem]++
			}
			for _, elem := range c.MasterDeck {
				replaced := strings.Replace(elem, "+1", "", -1)
				a.MasterDeckL[replaced]++
			}
			a.KilledBy[c.KilledBy]++
		}
	}
}

func sortMap(m map[string]int) {
	type kv struct {
        Key   string
        Value int
    }

    var ss []kv
    for k, v := range m {
        ss = append(ss, kv{k, v})
    }

    sort.Slice(ss, func(i, j int) bool {
        return ss[i].Value > ss[j].Value
    })

    for _, kv := range ss {
        fmt.Printf("%s, %d\n", kv.Key, kv.Value)
    }
}

func main() {
	a := NewAllRuns()
	b := run(a)
	//fmt.Printf("Items Purchased in wins: %v\n", b.ItemsPurchasedW)
	//fmt.Printf("Items purchased in losses: %v\n", b.ItemsPurchasedL)
	//fmt.Printf("Killed by: %v\n", b.KilledBy)
	//fmt.Printf("Deck in Wins: %v\n", b.MasterDeckW)
	//fmt.Printf("Deck in Losses: %v\n", b.MasterDeckL)
	//sortMap(b.MasterDeckW)
	//sortMap(b.MasterDeckL)
	sortMap(b.KilledBy)
}

func NewAllRuns() *AllRuns {
	return &AllRuns{ 
		KilledBy: make(map[string]int),
		ItemsPurchasedW: make(map[string]int),
		ItemsPurchasedL: make(map[string]int),
		MasterDeckW: make(map[string]int),
		MasterDeckL: make(map[string]int),	
	}
}

type AllRuns struct {
	KilledBy map[string]int
	ItemsPurchasedW map[string]int
	ItemsPurchasedL map[string]int
	MasterDeckW map[string]int
	MasterDeckL map[string]int
}

type STSRun struct {
	PathPerFloor        []interface{} `json:"path_per_floor"`
	CharacterChosen     string        `json:"character_chosen"`
	ItemsPurchased      []string      `json:"items_purchased"`
	GoldPerFloor        []int         `json:"gold_per_floor"`
	FloorReached        int           `json:"floor_reached"`
	CampfireRested      int           `json:"campfire_rested"`
	Playtime            int           `json:"playtime"`
	CurrentHpPerFloor   []int         `json:"current_hp_per_floor"`
	ItemsPurged         []string      `json:"items_purged"`
	Gold                int           `json:"gold"`
	Score               int           `json:"score"`
	PlayID              string        `json:"play_id"`
	LocalTime           string        `json:"local_time"`
	IsProd              bool          `json:"is_prod"`
	IsDaily             bool          `json:"is_daily"`
	ChoseSeed           bool          `json:"chose_seed"`
	IsAscensionMode     bool          `json:"is_ascension_mode"`
	CampfireUpgraded    int           `json:"campfire_upgraded"`
	Timestamp           int           `json:"timestamp"`
	PathTaken           []string      `json:"path_taken"`
	BuildVersion        string        `json:"build_version"`
	SeedSourceTimestamp int64         `json:"seed_source_timestamp"`
	PurchasedPurges     int           `json:"purchased_purges"`
	Victory             bool          `json:"victory"`
	MasterDeck          []string      `json:"master_deck"`
	MaxHpPerFloor       []int         `json:"max_hp_per_floor"`
	Relics              []string      `json:"relics"`
	CardChoices         []struct {
		NotPicked []string `json:"not_picked"`
		Picked    string   `json:"picked"`
		Floor     int      `json:"floor"`
	} `json:"card_choices"`
	PlayerExperience  int   `json:"player_experience"`
	PotionsFloorUsage []int `json:"potions_floor_usage"`
	DamageTaken       []struct {
		Damage  int    `json:"damage"`
		Enemies string `json:"enemies"`
		Floor   int    `json:"floor"`
		Turns   int    `json:"turns"`
	} `json:"damage_taken"`
	EventChoices []struct {
		EventName    string `json:"event_name"`
		PlayerChoice string `json:"player_choice"`
		Floor        int    `json:"floor"`
		DamageTaken  int    `json:"damage_taken"`
	} `json:"event_choices"`
	BossRelics []struct {
		NotPicked []string `json:"not_picked"`
		Picked    string   `json:"picked"`
	} `json:"boss_relics"`
	PotionsFloorSpawned []int  `json:"potions_floor_spawned"`
	SeedPlayed          string `json:"seed_played"`
	KilledBy            string `json:"killed_by"`
	AscensionLevel      int    `json:"ascension_level"`
	IsTrial             bool   `json:"is_trial"`
}
