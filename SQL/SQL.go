SELECT                что мы хотим вывести
    `id`,
    `name`,
    `reg_date`
FROM `user`           откуда хотим вывести
WHERE                 какие именно параметры поиска
    `name` LIKE '% Иван%' AND
    `reg_date` >= '2017-03-01' AND
    `reg_date` < '2017-12-01';