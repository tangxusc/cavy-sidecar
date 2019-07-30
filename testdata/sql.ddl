DROP TABLE IF EXISTS `events`;
CREATE TABLE `events`
(
    `id`          varchar(11) NOT NULL,
    `agg_type`    varchar(40) NOT NULL,
    `agg_id`      varchar(40) NOT NULL,
    `create_time` datetime    NOT NULL,
    `data`        tinytext    NOT NULL,
    `status`      integer     NOT NULL,
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