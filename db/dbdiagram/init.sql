CREATE TABLE `posts` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `title` varchar(100) NOT NULL,
  `content` text NOT NULL,
  `category` varchar(100) NOT NULL,
  `created_date` timestamp,
  `updated_date` timestamp,
  `status` varchar(100) NOT NULL
);
