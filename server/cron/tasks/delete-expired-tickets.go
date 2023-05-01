package tasks

import (
	"time"

	"github.com/VashUber/coursework-go/server/db"
	"github.com/VashUber/coursework-go/server/models"
	"github.com/robfig/cron/v3"
)

func InitDeleteExpiredTicketCron(c *cron.Cron) {
	c.AddFunc("@daily", func() {
		db.Database.Where("expired_date <= ?", time.Now()).Delete(&models.Ticket{})
	})
}
