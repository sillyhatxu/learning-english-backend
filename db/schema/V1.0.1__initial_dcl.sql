CREATE USER sillyhat_english identified BY 'sillyhat_english';
CREATE SCHEMA `sillyhat_english` DEFAULT CHARACTER SET utf8 COLLATE utf8_bin ;
GRANT ALL ON sillyhat_english.* TO sillyhat_english;
GRANT ALL ON sillyhat_english.* TO sillyhat;