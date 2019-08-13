package services

import "github.com/liuchamp/godemo/models"

type TestService struct {
}

func (ts *TestService) GetOne() models.Test {
	t := models.Test{}
	tc := models.Test{}
	s := models.Sim{}
	s.Name = "dsaingiasdf"
	s.Mod = "1"
	t.Name = "王老五"
	t.Oks = "dsadsfa"
	t.Smod = s
	tc.Name = "王老五"
	tc.Oks = "dsadsfa"
	tc.Smod = s
	t.Next = &tc
	return t
}
func (ts *TestService) GetMune() models.Menu {
	m := models.Menu{}
	return m
}

func (ts *TestService) GetSim() models.Sim {
	m := models.Sim{}
	m.Name = "屈的写的"
	m.Mod = "sdaf"
	m.BenKis = "sdafdsfadsaf"
	return m
}
