-- Структура таблицы `News`
CREATE TABLE IF NOT EXISTS "News" (
  "Id" BIGSERIAL PRIMARY KEY,
  "Title" TEXT NOT NULL,
  "Content" TEXT NOT NULL
);

-- Структура таблицы `NewsCategories`
CREATE TABLE IF NOT EXISTS "NewsCategories" (
  "NewsId" BIGINT NOT NULL,
  "CategoryId" BIGINT NOT NULL,
  PRIMARY KEY ("NewsId", "CategoryId")
);

-- Наполняем таблицу News
INSERT INTO "News" ("Title", "Content") VALUES
('Breaking News', 'This is the latest breaking news content.'),
('Tech Trends', 'Latest trends in the tech industry.'),
('Sports Update', 'Today’s sports news and highlights.'),
('Economy Insights', 'Detailed insights on the global economy.'),
('Health Tips', 'Top health tips for a better lifestyle.');

-- Наполняем таблицу NewsCategories
INSERT INTO "NewsCategories" ("NewsId", "CategoryId") VALUES
(1, 1),
(1, 2),
(2, 3),
(3, 1),
(3, 4),
(4, 2),
(5, 5);