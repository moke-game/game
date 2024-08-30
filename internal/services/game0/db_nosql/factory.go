package db_nosql

import (
	"errors"

	"github.com/gstones/moke-kit/orm/nerrors"
	"github.com/gstones/moke-kit/orm/nosql/diface"
	"go.uber.org/zap"

	"github.com/moke-game/game/internal/services/game0/db_nosql/game0"
)

type Database struct {
	logger *zap.Logger
	coll   diface.ICollection
}

func OpenDatabase(l *zap.Logger, coll diface.ICollection) Database {
	return Database{
		logger: l,
		coll:   coll,
	}
}

func (db *Database) LoadOrCreateDemo(id string) (*game0.Dao, error) {
	if dm, err := game0.NewDemoModel(id, db.coll); err != nil {
		return nil, err
	} else if err = dm.Load(); errors.Is(err, nerrors.ErrNotFound) {
		if dm, err = game0.NewDemoModel(id, db.coll); err != nil {
			return nil, err
		} else if err := dm.InitDefault(); err != nil {
			return nil, err
		} else if err = dm.Create(); err != nil {
			if err = dm.Load(); err != nil {
				return nil, err
			} else {
				return dm, nil
			}
		} else {
			return dm, nil
		}
	} else if err != nil {
		return nil, err
	} else {
		return dm, nil
	}
}
