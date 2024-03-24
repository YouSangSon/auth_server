package model

import "time"

type LottoResult struct {
	ID             int64     `gorm:"column:id;primaryKey"`
	BonusNo        int       `gorm:"column:bnus_no;not null"`
	DrawNo         int       `gorm:"column:drw_no;not null"`
	DrawNoDate     time.Time `gorm:"column:drw_no_date"`
	DrawtNo1       int       `gorm:"column:drwt_no1;not null"`
	DrawtNo2       int       `gorm:"column:drwt_no2;not null"`
	DrawtNo3       int       `gorm:"column:drwt_no3;not null"`
	DrawtNo4       int       `gorm:"column:drwt_no4;not null"`
	DrawtNo5       int       `gorm:"column:drwt_no5;not null"`
	DrawtNo6       int       `gorm:"column:drwt_no6;not null"`
	FirstAccumamnt int64     `gorm:"column:first_accumamnt;not null"`
	FirstPrzwnerCo int       `gorm:"column:first_przwner_co;not null"`
	FirstWinamnt   int64     `gorm:"column:first_winamnt;not null"`
	ReturnValue    string    `gorm:"column:return_value"`
	TotSellamnt    int64     `gorm:"column:tot_sellamnt;not null"`
}
