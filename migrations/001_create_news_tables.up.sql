-- Структура таблицы `News`
CREATE TABLE "News" (
  "Id" BIGSERIAL PRIMARY KEY,
  "Title" TEXT NOT NULL,
  "Content" TEXT NOT NULL
);

-- Структура таблицы `NewsCategories`
CREATE TABLE "NewsCategories" (
  "NewsId" BIGINT NOT NULL,
  "CategoryId" BIGINT NOT NULL,
  PRIMARY KEY ("NewsId", "CategoryId")
);