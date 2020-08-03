CREATE TABLE article(
   id BIGINT(20) PRIMARY KEY AUTO_INCREMENT,
   created_at DATETIME,
   updated_at DATETIME,
   deleted TINYINT,
   deleted_time	DATETIME,
   title VARCHAR(255) NOT NULL,
   folder_id BIGINT(20),FOREIGN KEY(folder_id) REFERENCES folder(id),
   mk_value TEXT,
   tags VARCHAR(255)
)
ENGINE=INNODB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;


CREATE TABLE folder(
   id BIGINT(20) PRIMARY KEY AUTO_INCREMENT,
   created_at DATETIME,
   updated_at DATETIME,
   deleted TINYINT,
   deleted_time	DATETIME,
   title VARCHAR(255) NOT NULL,
   folder_id BIGINT(20),FOREIGN KEY(folder_id) REFERENCES folder(id) 	
)
ENGINE=INNODB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;


CREATE TABLE my_book(
   id BIGINT(20) PRIMARY KEY AUTO_INCREMENT,
   title VARCHAR(255),
   writer VARCHAR(255),
   img_url VARCHAR(255),
   STATUS VARCHAR(255),
   COUNT INT(11),
   updated_at DATETIME	
)
ENGINE=INNODB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;