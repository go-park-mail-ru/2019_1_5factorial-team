package gameLogic

import (
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/app/config"
	"testing"
)

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
		if g.X != before + g.Speed {
			t.Error("#", i, "not moved")
		}
	}
}
