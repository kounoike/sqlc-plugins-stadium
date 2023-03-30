CREATE TABLE blog_comment (
    id int unsigned not null auto_increment,
    guid text not null,
    url text not null,
    contents text not null,
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP
);
