SELECT 
SUM(T1.`合計金額合計`+T1.`優待券販売（金額）`-T1.`未清算金額合計`+T1.`過不足`+T1.`手売（金額）`) AS `売上金額累計`
,SUM(T1.`入場券合計`) AS `入場者累計`
,SUM(CASE WHEN T1.`入場券合計` > 0 THEN 1 ELSE 0 END) AS `営業日数累計`
FROM (
SELECT `1号機現金金額`-`1号機精算金額` + `2号機現金金額`-`2号機精算金額` + `3号機現金金額`-`3号機精算金額` + `4号機現金金額`-`4号機精算金額` + `5号機現金金額`-`5号機精算金額` AS `合計金額合計`
,`優待券販売（金額）`
,`1号機未精算金額`+`2号機未精算金額`+`3号機未精算金額`+`4号機未精算金額`+`5号機未精算金額` AS 未清算金額合計
,`過不足`
,`手売（金額）`
,`大人入浴券枚数`+`大人入浴セット券枚数`+`小人入浴券枚数`+`回数券回収`+`招待券回収`+`優待券回収`+`6回数券回収`+`感謝祭招待券回収`+`ﾎﾟｲﾝﾄｶｰﾄﾞ大人回収`+`ﾎﾟｲﾝﾄｶｰﾄﾞﾞ小人回収`+`過去回数券回収` AS `入場券合計`
FROM `日次報告ﾃｰﾌﾞﾙ`
WHERE
    -- 日付 (date) must be greater than or equal to the start of the current fiscal year
    `日付` >= STR_TO_DATE(
        -- Calculate the Fiscal Year Start (FYS) date
        CONCAT(
            -- Determine the starting calendar year:
            YEAR(CURDATE()) - IF(
                -- If today is before Sept 21st of the current year (Jan 1 to Sep 20),
                -- the FYS started in the previous calendar year (subtract 1 from current year).
                MONTH(CURDATE()) < 9 OR (MONTH(CURDATE()) = 9 AND DAY(CURDATE()) < 21),
                1,
                -- Otherwise (Sept 21st to Dec 31st), the FYS started in the current calendar year (subtract 0).
                0
            ),
            '-09-21'
        ),
        '%Y-%m-%d'
    )
    ) AS T1;