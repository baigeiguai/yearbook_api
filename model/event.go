package model

/* CREATE TABLE `event` (
	`id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id generator',
	`title` varchar(256) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'event title',
	`uri` varchar(256) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'uri',
	`publish_time` varchar(256) COMMENT 'publish_time',
	PRIMARY KEY (`id`),
	UNIQUE KEY `uniq_uri` (`uri`)
  ) DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='yearbook_event'
*/
// Event  yearbook_event
type Event struct {
	ID          int64  `gorm:"column:id" db:"id" json:"id" form:"id"`                                         //  id generator
	Title       string `gorm:"column:title" db:"title" json:"title" form:"title"`                             //  event title
	Uri         string `gorm:"column:uri" db:"uri" json:"uri" form:"uri"`                                     //  uri
	PublishTime string `gorm:"column:publish_time" db:"publish_time" json:"publish_time" form:"publish_time"` //  publish_time
}

func (Event) TableName() string {
	return "event"
}
