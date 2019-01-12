ALTER TABLE request drop foreign key request_ibfk_1;
ALTER TABLE request ENGINE=MyISAM;
ALTER TABLE request add constraint binid FOREIGN KEY (bin) REFERENCES bin(id);
ALTER TABLE request DROP PRIMARY KEY, ADD PRIMARY KEY (bin,id);
