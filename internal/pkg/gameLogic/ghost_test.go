package gameLogic

import (
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/app/config"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/utils/log"
	"github.com/sirupsen/logrus"
	"testing"
)

func InitConfig() {
	configPath := "/etc/5factorial/"
	err := config.Init(configPath)
	if err != nil {
		logrus.Fatal(err.Error())
	}

	log.InitLogs()
}

func TestNewRandomGhost(t *testing.T) {
	for i := 0; i < 10; i++ {
		g := NewRandomGhost()
		if g.Damage != config.Get().GameConfig.DefaultDamage {
			t.Error("#", i, "damage not fit")
		}

		if g.Speed < 0 {
			if g.Speed != -config.Get().GameConfig.DefaultMovementSpeed {
				t.Error("#", i, "speed not fit, want:", -config.Get().GameConfig.DefaultMovementSpeed,
					", have:", g.Speed)
			}

			if g.X != config.Get().GameConfig.DefaultLeftPosition-config.Get().GameConfig.DefaultSpriteWidth {
				t.Error("#", i, "x not fit, want:",
					config.Get().GameConfig.DefaultLeftPosition-config.Get().GameConfig.DefaultSpriteWidth,
					", have:", g.X)

			}
		} else {
			if g.Speed != config.Get().GameConfig.DefaultMovementSpeed {
				t.Error("#", i, "speed not fit, want:", config.Get().GameConfig.DefaultMovementSpeed,
					", have:", g.Speed)
			}

			if g.X != config.Get().GameConfig.DefaultRightPosition+config.Get().GameConfig.DefaultSpriteWidth {
				t.Error("#", i, "x not fit, want:",
					config.Get().GameConfig.DefaultRightPosition+config.Get().GameConfig.DefaultSpriteWidth,
					", have:", g.X)
			}

		}
	}
}

var casesNewGhost = []struct {
	startPosition int
	damage        uint32
	speed         int
	symbolsLen    int
}{
	{
		startPosition: 1,
		damage:        1,
		speed:         0,
		symbolsLen:    1,
	},
	{
		startPosition: -1,
		damage:        1,
		speed:         -1,
		symbolsLen:    1,
	},
}

func TestNewGhost(t *testing.T) {
	for i, val := range casesNewGhost {
		res := NewGhost(val.startPosition, val.damage, val.speed, val.symbolsLen)

		if res.X != val.startPosition || res.Damage != val.damage || res.Speed != val.speed {
			t.Error("#", i, "RES expected:", val, "have:", res)
			continue
		}
	}
}

func TestNewRightGhost(t *testing.T) {
	for i := 0; i < 10; i++ {
		g := NewRightGhost()
		if g.Damage != config.Get().GameConfig.DefaultDamage {
			t.Error("#", i, "damage not fit")
		}

		if g.Speed != -config.Get().GameConfig.DefaultMovementSpeed {
			t.Error("#", i, "speed not fit, want:", -config.Get().GameConfig.DefaultMovementSpeed,
				", have:", g.Speed)
		}

		if g.X != config.Get().GameConfig.DefaultRightPosition+config.Get().GameConfig.DefaultSpriteWidth {
			t.Error("#", i, "x not fit, want:",
				config.Get().GameConfig.DefaultRightPosition+config.Get().GameConfig.DefaultSpriteWidth,
				", have:", g.X)

		}
	}
}

func TestNewLeftGhost(t *testing.T) {
	for i := 0; i < 10; i++ {
		g := NewLeftGhost()
		if g.Damage != config.Get().GameConfig.DefaultDamage {
			t.Error("#", i, "damage not fit")
		}

		if g.Speed != config.Get().GameConfig.DefaultMovementSpeed {
			t.Error("#", i, "speed not fit, want:", config.Get().GameConfig.DefaultMovementSpeed,
				", have:", g.Speed)
		}

		if g.X != config.Get().GameConfig.DefaultLeftPosition-config.Get().GameConfig.DefaultSpriteWidth {
			t.Error("#", i, "x not fit, want:",
				config.Get().GameConfig.DefaultLeftPosition-config.Get().GameConfig.DefaultSpriteWidth,
				", have:", g.X)

		}
	}
}

func TestGhost_Move(t *testing.T) {
	for i := 0; i < 10; i++ {
		g := NewLeftGhost()
		before := g.X
		g.Move()
		if g.X != before+g.Speed {
			t.Error("#", i, "not moved")
		}
	}
}

func TestNewGhostStack(t *testing.T) {
	if len(NewGhostStack().Items) != 0 {
		t.Error("#", 0, "not empty queue")
	}
}

func TestGhostQueue_PushBack(t *testing.T) {
	nq := NewGhostStack()
	gh := NewRandomGhost()

	nq.PushBack(gh)
	if nq.Items[len(nq.Items)-1].Speed != gh.Speed {
		t.Error("not valid speed")
	}
	if nq.Items[len(nq.Items)-1].X != gh.X {
		t.Error("not valid X")
	}
	if nq.Items[len(nq.Items)-1].Damage != gh.Damage {
		t.Error("not valid damage")
	}
}

func TestGhostQueue_PopBack(t *testing.T) {
	nq := NewGhostStack()
	gh := NewRandomGhost()
	ghL := NewRandomGhost()
	nq.PushBack(gh)
	nq.PushBack(ghL)

	sizeBefore := len(nq.Items) - 1
	nq.PopBack()
	if len(nq.Items)-1 == sizeBefore {
		t.Error("element not popped")
	}
}

func TestGhostQueue_AddNewGhost(t *testing.T) {
	nq := NewGhostStack()
	nq.AddNewGhost()

	if len(nq.Items) == 0 {
		t.Error("item wasnt generate")
	}

	nq.PopBack()

	nq.PushBack(NewLeftGhost())
	nq.AddNewGhost()
	if nq.Items[1].Speed < 0 {
		t.Error("item generated wrong ghost, wanted right, got left")
	}

	nq.PopBack()
	nq.PopBack()

	nq.PushBack(NewRightGhost())
	nq.AddNewGhost()
	if nq.Items[1].Speed > 0 {
		t.Error("item generated wrong ghost, wanted left, got right")
	}
}

func TestGhostQueue_MoveAllGhosts(t *testing.T) {
	InitConfig()
	// зависит от дефолтных конфигов
	nq := NewGhostStack()

	nq.MoveAllGhosts()
	nq.PopFront()

	nq.AddNewGhost()
	nq.AddNewGhost()

	for i := 0; i < 10; i++ {
		if nq.MoveAllGhosts() {
			nq.PopFront()
		}
	}

	if len(nq.Items) != 0 {
		t.Error("check default vars")
	}
}

func TestGhostQueue_Len(t *testing.T) {
	nq := NewGhostStack()
	nq.AddNewGhost()
	nq.AddNewGhost()

	if nq.Len() != 2 {
		t.Error("wrong size of items, have:", nq.Len(), "exp:", 2)
	}
}

func TestGhostQueue_PopSymbol(t *testing.T) {
	InitConfig()
	// зависит от дефолтных конфигов
	nq := NewGhostStack()

	nq.MoveAllGhosts()
	nq.PopFront()

	nq.AddNewGhost()
	nq.AddNewGhost()

	score := nq.PopSymbol(LR)
	if score%config.Get().GameConfig.ScoreMatchSymbol != 0 {
		t.Error("have score != n*20")
	}
}
