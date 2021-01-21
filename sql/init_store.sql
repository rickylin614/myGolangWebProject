CREATE TABLE bento.store (
  id bigint NOT NULL AUTO_INCREMENT,
  name varchar(100)  NOT NULL,
  phone_no varchar(11) NOT NULL,
  region int ,
  create_user varchar(100)  NOT NULL,
  update_user varchar(100)  ,
  created_at timestamp(6) NULL DEFAULT NULL,
  updated_at timestamp(6) NULL DEFAULT NULL,
  deleted_at timestamp(6) NULL DEFAULT NULL,
  PRIMARY KEY (id),
  KEY idx_store_id (id),
  KEY idx_store_deleted_at (deleted_at)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci


CREATE TABLE bento.food (
  id bigint NOT NULL AUTO_INCREMENT,
  name varchar(100) NOT NULL,
  store_id bigint DEFAULT NULL,
  store_name varchar(100) DEFAULT NULL,
  description varchar(100) DEFAULT NULL,
  create_user varchar(100)  NOT NULL,
  update_user varchar(100)  ,
  created_at timestamp(6) NULL DEFAULT NULL,
  updated_at timestamp(6) NULL DEFAULT NULL,
  deleted_at timestamp(6) NULL DEFAULT NULL,
  PRIMARY KEY (id),
  KEY idx_food_id (id),
  KEY idx_food_deleted_at (deleted_at)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

CREATE TABLE bento.OB_TB_ENUM_TYPE (
	ENUM_TYPE_OID			bigint NOT NULL COMMENT '系統參數類別識別號',
	CODE					varchar(100) NOT NULL COMMENT '代碼',
	NAME					varchar(100) NOT NULL COMMENT '名稱',
	create_user varchar(100)  NOT NULL COMMENT '建立資料使用者',
	update_user varchar(100)  COMMENT '修改資料使用者',
	created_at timestamp(6) NULL DEFAULT NULL COMMENT '建立資料日期',
	updated_at timestamp(6) NULL DEFAULT NULL COMMENT '修改資料日期',
	deleted_at timestamp(6) NULL DEFAULT NULL COMMENT '刪除資料日期',
	VERSION					int NOT NULL COMMENT '版本',
	PRIMARY KEY (ENUM_TYPE_OID),
	UNIQUE KEY EWP_UQ1_ENUM_TYPE (CODE),
	KEY idx_ENUM_TYPE_id (ENUM_TYPE_OID),
	KEY idx_ENUM_TYPE_deleted_at (deleted_at)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE bento.OB_TB_ENUMERATION (
	ENUM_OID				bigint NOT NULL COMMENT '系統參數類別識別號',
	ENUM_TYPE_OID			bigint NOT NULL COMMENT '系統參數類別識別號',
	CODE					varchar(100) NOT NULL COMMENT '代碼',
	NAME					varchar(100) NOT NULL COMMENT '名稱',
	create_user varchar(100)  NOT NULL COMMENT '建立資料使用者',
	update_user varchar(100)  COMMENT '修改資料使用者',
	created_at timestamp(6) NULL DEFAULT NULL COMMENT '建立資料日期',
	updated_at timestamp(6) NULL DEFAULT NULL COMMENT '修改資料日期',
	deleted_at timestamp(6) NULL DEFAULT NULL COMMENT '刪除資料日期',
	VERSION					int NOT NULL COMMENT '版本',
	PRIMARY KEY (ENUM_OID),
	UNIQUE KEY EWP_UQ1_ENUM (CODE),
	KEY idx_ENUM_id (ENUM_OID),
	KEY idx_ENUM_deleted_at (deleted_at)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;