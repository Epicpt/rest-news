-- +goose Up
-- +goose StatementBegin
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
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM "News" WHERE "Title" IN ('Breaking News', 'Tech Trends', 'Sports Update', 'Economy Insights', 'Health Tips');
-- +goose StatementEnd
