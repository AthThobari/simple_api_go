ALTER TABLE posts DROP FOREIGN KEY fk_user_id_posts;

ALTER TABLE posts 
MODIFY COLUMN user_id INT