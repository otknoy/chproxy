package model

import "strings"

type Thread struct {
	Title   string
	ResList []Res
}

type Res struct {
	Name  string
	EMail string
	Date  string
	ID    string
	BE    string
	Text  string
}

func (t *Thread) ToDat() (dat string) {
	for i := 0; i < len(t.ResList); i++ {
		row := t.ResList[i].ToDatRow()

		if i == 0 {
			dat += row + t.Title + "\n"
			continue
		}

		dat += row + "\n"
	}

	return dat
}

func (r *Res) ToDatRow() string {
	text := strings.ReplaceAll(r.Text, "\n", " <br> ")

	be := ""
	if r.BE != "" {
		be = " BE:" + r.BE
	}

	return r.Name + "<>" + r.EMail + "<>" + r.Date + " ID:" + r.ID + be + "<>" + text + "<>"
}
