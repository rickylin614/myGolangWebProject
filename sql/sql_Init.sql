
--會員
DROP TABLE user;


CREATE TABLE user (
  id bigint NOT NULL AUTO_INCREMENT,
  name varchar(100) DEFAULT '',
  pwd varchar(100) DEFAULT '0',
  session_id varchar(100) DEFAULT NULL,
  login_time timestamp(6) NULL DEFAULT NULL,
  created_at timestamp(6) NULL DEFAULT NULL,
  updated_at timestamp(6) NULL DEFAULT NULL,
  deleted_at timestamp(6) NULL DEFAULT NULL,
  PRIMARY KEY (id),
  UNIQUE KEY user_UN (name),
  KEY idx_user_id (id),
  KEY idx_user_deleted_at (deleted_at)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci