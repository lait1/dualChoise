-- +migrate Up

CREATE TABLE `categories`
(
    `id`      int(11)     NOT NULL,
    `name`    varchar(64) NOT NULL,
    `preview` varchar(60) NOT NULL,
    `created` date        NOT NULL
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8;

CREATE TABLE `options`
(
    `id`       int(11)          NOT NULL,
    `name`     varchar(64)      NOT NULL,
    `preview`  varchar(64)      NOT NULL,
    `priority` int(11) UNSIGNED NOT NULL
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8;

CREATE TABLE `quiz`
(
    `id`          int(11)     NOT NULL,
    `name`        varchar(64) NOT NULL,
    `category_id` int(11)     NOT NULL,
    `preview`     varchar(64) DEFAULT NULL
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8;

CREATE TABLE `quiz_options`
(
    `quiz_id`   int(11)          NOT NULL,
    `option_id` int(11)          NOT NULL,
    `wins`      int(11) UNSIGNED NOT NULL DEFAULT '0'
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8;

CREATE TABLE `translations`
(
    `var_name` varchar(64) NOT NULL,
    `ru`       varchar(250) DEFAULT NULL,
    `en`       varchar(250) DEFAULT NULL,
    `es`       varchar(250) DEFAULT NULL
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8;

ALTER TABLE `categories`
    ADD PRIMARY KEY (`id`),
    ADD KEY `name` (`name`);

ALTER TABLE `options`
    ADD PRIMARY KEY (`id`),
    ADD KEY `name_var` (`name`) USING BTREE;

ALTER TABLE `quiz`
    ADD PRIMARY KEY (`id`),
    ADD KEY `category_id` (`category_id`),
    ADD KEY `name` (`name`);

ALTER TABLE `quiz_options`
    ADD KEY `option_id` (`option_id`),
    ADD KEY `quiz_id` (`quiz_id`);

ALTER TABLE `translations`
    ADD PRIMARY KEY (`var_name`);

ALTER TABLE `categories`
    MODIFY `id` int(11) NOT NULL AUTO_INCREMENT,
    AUTO_INCREMENT = 14;

ALTER TABLE `options`
    MODIFY `id` int(11) NOT NULL AUTO_INCREMENT,
    AUTO_INCREMENT = 9;

ALTER TABLE `quiz`
    MODIFY `id` int(11) NOT NULL AUTO_INCREMENT,
    AUTO_INCREMENT = 10;

ALTER TABLE `categories`
    ADD CONSTRAINT `categories_ibfk_1` FOREIGN KEY (`name`) REFERENCES `translations` (`var_name`) ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE `options`
    ADD CONSTRAINT `options_ibfk_1` FOREIGN KEY (`name`) REFERENCES `translations` (`var_name`) ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE `quiz`
    ADD CONSTRAINT `quiz_ibfk_1` FOREIGN KEY (`category_id`) REFERENCES `categories` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
    ADD CONSTRAINT `quiz_ibfk_2` FOREIGN KEY (`name`) REFERENCES `translations` (`var_name`) ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE `quiz_options`
    ADD CONSTRAINT `quiz_options_ibfk_1` FOREIGN KEY (`option_id`) REFERENCES `options` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
    ADD CONSTRAINT `quiz_options_ibfk_2` FOREIGN KEY (`quiz_id`) REFERENCES `quiz` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

-- +migrate Down

DROP TABLE IF EXISTS `categories`;
DROP TABLE IF EXISTS `options`;
DROP TABLE IF EXISTS `quiz`;
DROP TABLE IF EXISTS `quiz_options`;
DROP TABLE IF EXISTS `translations`;

SELECT 1;
