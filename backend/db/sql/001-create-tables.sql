-- ローカルのテストデータの作成 --
create table IF not exists `user`
(
    `id`               INT(20) AUTO_INCREMENT,              -- ID
    `uuid`             VARCHAR(20) NOT NULL UNIQUE,         -- ユニークID
    `name`             VARCHAR(100) NOT NULL UNIQUE,        -- ユーザーネーム
    `group`            VARCHAR(100) NOT NULL,               -- ユーザーが属しているグループ
    `work`             VARCHAR(10000) NOT NULL,             -- 役職
    `created_at`       Datetime DEFAULT NULL,
    `updated_at`       Datetime DEFAULT NULL,
    PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;


