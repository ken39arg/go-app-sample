
CREATE TABLE isubata_user (
    id              BIGINT         NOT NULL COMMENT 'katsubushi',
    name            VARBINARY(100) NOT NULL COMMENT 'アカウント名',
    salt            VARBINARY(255) NOT NULL COMMENT 'salt',
    pass            VARBINARY(255) NOT NULL COMMENT 'パスワード',
    display_name    VARCHAR(100)   NOT NULL COMMENT '表示名',
    avatar_icon     VARCHAR(100)   NOT NULL COMMENT 'アイコンURL',
    created_at      DATETIME       NOT NULL,
    UNIQUE name_uiq (name),
    PRIMARY KEY (id)
)
ENGINE=InnoDB
DEFAULT CHARACTER SET utf8mb4
COMMENT='ISUBATA ユーザー';

CREATE TABLE isubata_channel (
    id              BIGINT         NOT NULL COMMENT 'katsubushi',
    name            VARBINARY(100) NOT NULL COMMENT 'ルーム名',
    updated_at      DATETIME NOT NULL,
    created_at      DATETIME       NOT NULL,
    UNIQUE name_uiq (name),
    PRIMARY KEY (id)
)
ENGINE=InnoDB
DEFAULT CHARACTER SET utf8mb4
COMMENT='ISUBATA チャンネル';

CREATE TABLE isubata_message (
    id              BIGINT         NOT NULL COMMENT 'katsubushi',
    channel_id      BIGINT         NOT NULL COMMENT '=isubata_channel.id',
    user_id         BIGINT         NOT NULL COMMENT '=isubata_user.id',
    content         VARCHAR(191)   NOT NULL COMMENT '本文',
    created_at      DATETIME       NOT NULL,
    INDEX  channel_id_idx(channel_id),
    PRIMARY KEY (id)
)
ENGINE=InnoDB
DEFAULT CHARACTER SET utf8mb4
COMMENT='ISUBATA メッセージ';
