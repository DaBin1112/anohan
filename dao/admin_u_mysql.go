package dao

import (
	"anohan/model"
	"database/sql"
	"github.com/donnie4w/go-logger/logger"
)

const (
	_findLoginSQL = "select id,ctime,mtime,username,email,type,password,photo,ip,phone,state from admin_u where phone = ?"
)

func (d *Dao) GetUserByPhone(phone string) (r *model.AdminUModel, err error) {
	r = &model.AdminUModel{}
	row := d.db.QueryRow(_findLoginSQL, phone)
	if err = row.Scan(
		&r.Id,
		&r.Ctime,
		&r.Mtime,
		&r.Username,
		&r.Email,
		&r.Type,
		&r.Password,
		&r.Photo,
		&r.Ip,
		&r.Phone,
		&r.State,
	); err != nil {
		if err == sql.ErrNoRows {
			r = nil
			err = nil
		}
		logger.Info(err)
	}
	return
}
