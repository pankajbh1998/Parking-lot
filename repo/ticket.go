package repo

import (
	"Parking-lot/models"
	"math/rand"
)

type Ticket struct {
	mp map[int64]*models.Ticket
	mpr map[string]*models.Ticket
}

func NewTicket()*Ticket{
	return &Ticket{
		mp: make(map[int64]*models.Ticket),
		mpr: make(map[string]*models.Ticket),
	}
}

func (t *Ticket)GetNewTicket(ticket *models.Ticket) int64 {
	ticket.ID = rand.Int63()
	t.mp[ticket.ID] = ticket
	t.mpr[ticket.ConvertToStr()] = ticket
	return ticket.ID
}

func (t *Ticket)GetTicket(id int64)*models.Ticket {
	val, ok := t.mp[id]
	if !ok {
		return nil
	}
	return val
}

func (t *Ticket)GetTicketByTicketFormat(id string)*models.Ticket {
	val, ok := t.mpr[id]
	if !ok {
		return nil
	}
	return val
}

func (t *Ticket)RemoveTicket(id int64){
	delete(t.mp, id)
}
