WITH RECURSIVE date_series (calendar_date)
AS (
    SELECT DATE (?) AS calendar_date
    
    UNION ALL
    
    SELECT DATE_ADD(calendar_date, INTERVAL 1 DAY)
    FROM date_series
    WHERE calendar_date < ?
    ),
-- Second CTE to combine current year dates and previous year dates
all_dates (combined_date)
AS (
    -- Select the dates from the current year (the first CTE)
    SELECT calendar_date
    FROM date_series
    
    UNION ALL
    
    -- Select the corresponding dates from the previous year
    SELECT DATE_SUB(calendar_date, INTERVAL 1 YEAR)
    FROM date_series
    )
SELECT d.combined_date AS '日付',
    
+ n.`1号機現金金額` 
+ n.`2号機現金金額` 
+ n.`3号機現金金額` 
+ n.`4号機現金金額` 
+ n.`5号機現金金額` 
- n.`1号機精算金額` 
- n.`2号機精算金額` 
- n.`3号機精算金額` 
- n.`4号機精算金額` 
- n.`5号機精算金額` 
- n.`1号機未精算枚数` 
- n.`2号機未精算枚数` 
- n.`3号機未精算枚数` 
- n.`4号機未精算枚数` 
- n.`5号機未精算枚数` 
+ n.`1号機QR金額` 
+ n.`2号機QR金額` 
+ n.`3号機QR金額` 
+ n.`4号機QR金額` 
+ n.`5号機QR金額` 
- n.`1号機QR精算金額` 
- n.`2号機QR精算金額` 
- n.`3号機QR精算金額` 
- n.`4号機QR精算金額` 
- n.`5号機QR精算金額` 
+ n.`1号機電子マネ金額` 
+ n.`2号機電子マネ金額` 
+ n.`3号機電子マネ金額` 
+ n.`4号機電子マネ金額` 
+ n.`5号機電子マネ金額` 
- n.`1号機電子マネ精算金額` 
- n.`2号機電子マネ精算金額` 
- n.`3号機電子マネ精算金額` 
- n.`4号機電子マネ精算金額` 
- n.`5号機電子マネ精算金額` 
+ n.`1号機クレジット金額` 
+ n.`2号機クレジット金額` 
+ n.`3号機クレジット金額` 
+ n.`4号機クレジット金額` 
+ n.`5号機クレジット金額` 
- n.`1号機クレジット精算金額` 
- n.`2号機クレジット精算金額` 
- n.`3号機クレジット精算金額` 
- n.`4号機クレジット精算金額` 
- n.`5号機クレジット精算金額` 
+ `n`.`優待券販売（金額）` 
+ n.リアレジ金額
- n.`1号機電子マネ精算金額`
- n.`2号機電子マネ精算金額`
- n.`3号機電子マネ精算金額`
- n.`4号機電子マネ精算金額`
- n.`5号機電子マネ精算金額`
AS 'total',
    n.`大人入浴券枚数` 
+ n.`大人入浴セット券枚数` 
+ n.`小人入浴券枚数` 
+ n.`6回数券回収` 
+ n.`回数券回収` 
+ n.`招待券回収` 
+ n.`優待券回収` 
+ n.`感謝祭招待券回収` 
+ n.`ﾎﾟｲﾝﾄｶｰﾄﾞ大人回収` 
+ n.`ﾎﾟｲﾝﾄｶｰﾄﾞﾞ小人回収` 
+ n.`過去回数券回収` AS visitors,
    CASE 
        WHEN k.日付 IS NULL
            THEN 0
        ELSE 1
        END AS 'closed',
    CASE 
        WHEN s.日付 IS NOT NULL
            THEN 1
        ELSE 0
        END AS 'isHoliday',
    CASE 
        WHEN s.日付 IS NOT NULL
            OR DAYOFWEEK(d.combined_date) = 1
            OR DAYOFWEEK(d.combined_date) = 7
            THEN 1
        ELSE 0
        END AS 'isHolidayOrWeekend'
FROM all_dates d
LEFT OUTER JOIN 日次報告ﾃｰﾌﾞﾙ n ON d.combined_date = n.日付
LEFT OUTER JOIN t_休館日 k ON k.日付 = `d`.combined_date
LEFT OUTER JOIN t_祝日 s ON s.日付 = `d`.combined_date
ORDER BY d.combined_date;