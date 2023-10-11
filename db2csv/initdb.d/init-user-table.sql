CREATE DATABASE tmpdb;

\c tmpdb;

CREATE TABLE "user" (
  uid serial,
  name varchar(30) default '',
  phone varchar(20) default '',
  email varchar(30) default '',
  password varchar(100) default '',
  primary key (uid)
);

BEGIN;
insert into "user" (name, email, phone, password) values ('McGuigan', 'amcguigan0@china.com.cn', '426-483-0116', '$2a$04$VuqAAGdnXzeKsxZ0I/RcCOAPlQbMw1X/eei7xZNQ8EEEYzwUGjuvC');
insert into "user" (name, email, phone, password) values ('Ralphs', 'aralphs1@mashable.com', '538-553-4693', '$2a$04$ENSljrPRNWRUIGoSuERIP.nGggzF.s9YiQRA9Ha4Exrm/2M9vjbPS');
insert into "user" (name, email, phone, password) values ('Eich', 'deich2@weather.com', '757-726-7410', '$2a$04$OKcuCtvdnjvPEOOycka5jOzpaY2PhwyZN2em1T5nd5TS7JAMfKnO2');
insert into "user" (name, email, phone, password) values ('MacPhail', 'bmacphail3@cam.ac.uk', '241-240-2924', '$2a$04$7zcWtBUWpFMITxL/Mz.Qke/9dWw.JevH17vUyW1hQLZIbC7BiLdTu');
insert into "user" (name, email, phone, password) values ('Ixor', 'bixor4@timesonline.co.uk', '297-276-5471', '$2a$04$z4bs4G24yO9BKEHCo7kZRus/vQUM7X6PZA1xvVERwmUCqFfQiQHO2');
insert into "user" (name, email, phone, password) values ('Mumbeson', 'lmumbeson5@craigslist.org', '754-556-4652', '$2a$04$f592nwPlzVSEWBivOFXuouyN4pZ36fYL4rZW7OsUKsm8wLb7J2sCG');
insert into "user" (name, email, phone, password) values ('Demaine', 'mdemaine6@mlb.com', '924-334-4118', '$2a$04$mJ0g03F961BQCos/2BCDSueYlIDYHb8aBJcbgEwwqzDFnVNdzQRHK');
insert into "user" (name, email, phone, password) values ('Kubacki', 'ckubacki7@reverbnation.com', '464-956-2880', '$2a$04$izUKIdcDYIpSCLz6NSIgf.dhT4I51K/VqPRWpnSbq5CizzLmJ7hQe');
insert into "user" (name, email, phone, password) values ('Kirmond', 'tkirmond8@typepad.com', '765-323-1410', '$2a$04$LJGL/rTBB8naBgAyXsviKuEgUBMM8U/QkZF51P3UHLFe6eiwxdOOy');
insert into "user" (name, email, phone, password) values ('Loton', 'rloton9@buzzfeed.com', '413-125-1829', '$2a$04$YBLwdmibqBnkWmrVOZJHreJbkf/3hVW4ZGA9aQKiicrR0Ch6X98Uu');
COMMIT;
