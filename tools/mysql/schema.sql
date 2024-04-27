CREATE TABLE `user`
(
    `user_id`       BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'ユーザー識別子',
    `user_name`     varchar(20) NOT NULL COMMENT 'ユーザー名',
    `password` VARCHAR(80) NOT NULL COMMENT 'パスワードハッシュ',
    `role`     VARCHAR(80) NOT NULL COMMENT 'ロール',
    `created`  DATETIME(6) NOT NULL COMMENT 'レコード作成日時',
    `modified` DATETIME(6) NOT NULL COMMENT 'レコード修正日時',
    PRIMARY KEY (`user_id`),
    UNIQUE KEY `uix_name` (`user_name`) USING BTREE
) Engine=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='ユーザー';

CREATE TABLE `post`
(
    `post_id`  BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'ポストの識別子',
    `user_id`  BIGINT UNSIGNED NOT NULL COMMENT 'ポストを投稿したユーザーの識別子',
    `title`    VARCHAR(128) NOT NULL COMMENT 'ポストのタイトル',
    `comment`  TEXT  NOT NULL COMMENT 'ポストの詳細',
    `created`  DATETIME(6) NOT NULL COMMENT 'レコード作成日時',
    `modified` DATETIME(6) NOT NULL COMMENT 'レコード修正日時',
    PRIMARY KEY (`post_id`),
    CONSTRAINT `fk_user_id`
        FOREIGN KEY (`user_id`) REFERENCES `user` (`user_id`)
            ON DELETE RESTRICT ON UPDATE RESTRICT
) Engine=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='ポスト';

CREATE TABLE `reply`
(
    `reply_id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'リプライの識別子',
    `post_id`  BIGINT UNSIGNED NOT NULL COMMENT '対応するポストの識別子',
    `user_id`  BIGINT UNSIGNED NOT NULL COMMENT 'リプライを投稿したユーザーの識別子',
    `title`    VARCHAR(128) NOT NULL COMMENT 'リプライのタイトル',
    `comment`  TEXT  NOT NULL COMMENT 'リプライの詳細',
    `created`  DATETIME(6) NOT NULL COMMENT 'レコード作成日時',
    `modified` DATETIME(6) NOT NULL COMMENT 'レコード修正日時',
    PRIMARY KEY (`reply_id`),
    CONSTRAINT `fk_reply_post_id`
        FOREIGN KEY (`post_id`) REFERENCES `post` (`post_id`)
            ON DELETE RESTRICT ON UPDATE RESTRICT,
    CONSTRAINT `fk_reply_user_id`
        FOREIGN KEY (`user_id`) REFERENCES `user` (`user_id`)
            ON DELETE RESTRICT ON UPDATE RESTRICT
) Engine=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='リプライ';