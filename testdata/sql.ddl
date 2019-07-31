DROP TABLE IF EXISTS `events`;
CREATE TABLE `events`
(
    `id`             varchar(11) NOT NULL,
    `event_type`     varchar(11) NOT NULL,
    `agg_type`       varchar(40) NOT NULL,
    `agg_id`         varchar(40) NOT NULL,
    `create_time`    datetime    NOT NULL,
    `data`           tinytext    NOT NULL,
    `handler_status` integer     NOT NULL,
    `mq_status`      integer     NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8;


DROP TABLE IF EXISTS `snapshot`;
CREATE TABLE `snapshot`
(
    `id`          varchar(11) NOT NULL,
    `agg_type`    varchar(40) NOT NULL,
    `agg_id`      varchar(40) NOT NULL,
    `create_time` datetime    NOT NULL,
    `data`        tinytext    NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8;