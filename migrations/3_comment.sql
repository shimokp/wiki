-- +migrate Up
CREATE TABLE `comments` (
  `comment_id` int(11) NOT NULL AUTO_INCREMENT,
  `article_id` int(11) NOT NULL,
    `user_id` int(11),
  `body` TEXT NOT NULL,
  `created` timestamp NOT NULL DEFAULT NOW(),
  `updated` timestamp NOT NULL DEFAULT NOW() ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`comment_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='list of comments';

-- +migrate Down
DROP TABLE comments;
