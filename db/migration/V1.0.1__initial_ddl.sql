CREATE TABLE IF NOT EXISTS `user`
(
  `id`                 bigint(48)   NOT NULL AUTO_INCREMENT,
  `login_name`         varchar(50)  NOT NULL,
  `password`           varchar(50)  NOT NULL,
  `channel`            int                   DEFAULT NULL,
  `type`               int                   DEFAULT NULL,
  `status`             varchar(10),
  `created_time`       timestamp(3) NOT NULL DEFAULT current_timestamp(3),
  `last_modified_time` timestamp(3) NOT NULL DEFAULT current_timestamp(3) ON UPDATE current_timestamp(3),
  PRIMARY KEY (`id`),
  INDEX (`login_name`, `channel`, `type`),
  INDEX (`status`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;