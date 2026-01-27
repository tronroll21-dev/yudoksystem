package models

import (
	"log"
	"sort"
)

type SalesDetail struct {
	ID                  int     `json:"ID"`
	Bumon               string  `json:"部門"`
	CategoryName        string  `json:"カテゴリー名"`
	CategoryName_Sanuki string  `json:"カテゴリー名讃岐"`
	MenuCode            int     `json:"メニューコード"`
	MenuName            string  `json:"メニュー名"`
	Tanka               int     `json:"単価"`
	UriageKingaku       int     `json:"売上金額"`
	MakanaiKubun        string  `json:"まかない区分"`
	UriageMaisuu        int     `json:"売上枚数"`
	Hiduke              []uint8 `json:"日付"`
}

type SoldProduct struct {
	Bumon         int    `json:"bumon"`
	BumonName     string `json:"bumonName"`
	UriageKingaku int    `json:"uriageKingaku"`
	HanbaiMaisuu  int    `json:"hanbaiMaisuu"`
	ProductID     int    `json:"productID"`
	MenuName      string `json:"menuName"`
	TekiyouKakaku int    `json:"tekiyouKakaku"`
	MakanaiKubun  string `json:"makanaiKubun"`
	Date          string `json:"date"`
	Category      string `json:"category"`
}

// Sale represents a row in your sales table.
// type Sale struct {
// 	部門          string
// 	カテゴリー名      string
// 	ProductName string
// 	売上金額        float64
// }

// GroupedSale represents the hierarchical output structure.
type GroupedSale struct {
	Name      string
	Total     int
	SubGroups []GroupedSale
	Products  map[int]SalesDetail
}

// processSales takes the flat sales data, sorts it, and groups it hierarchically.
func processSales(sales []SalesDetail) GroupedSale {
	// 1. Sort the data: 部門 > カテゴリー名 > 売上金額 (descending)
	sort.Slice(sales, func(i, j int) bool {
		if sales[i].Bumon != sales[j].Bumon {
			return sales[i].Bumon < sales[j].Bumon
		}
		if sales[i].CategoryName != sales[j].CategoryName {
			return sales[i].CategoryName < sales[j].CategoryName
		}
		// Assuming "product number" means sorting by revenue (or a unique ID if available).
		// Sorting by 売上金額 in descending order to highlight top products.
		return sales[i].UriageKingaku > sales[j].UriageKingaku
	})

	// 2. Group the data: 部門 -> CategoryName -> Products
	categoryGroups := make(map[string]map[string]map[int]SalesDetail)
	var grandTotal int

	for _, s := range sales {

		catName := s.Bumon
		subCatName := s.CategoryName
		menuCode := s.MenuCode
		grandTotal += s.UriageKingaku

		// Initialize Category Map if needed
		if _, ok := categoryGroups[catName]; !ok {
			categoryGroups[catName] = make(map[string]map[int]SalesDetail)
		}

		// Initialize SubCategory GroupedSale if needed
		if _, ok := categoryGroups[catName][subCatName]; !ok {
			categoryGroups[catName][subCatName] = make(map[int]SalesDetail)
		}

		// Initialize Product if needed
		if _, ok := categoryGroups[catName][subCatName][menuCode]; !ok {
			categoryGroups[catName][subCatName][menuCode] = SalesDetail{
				ID:                  s.ID,
				Bumon:               s.Bumon,
				CategoryName:        s.CategoryName,
				CategoryName_Sanuki: s.CategoryName_Sanuki,
				MenuCode:            s.MenuCode,
				MenuName:            s.MenuName,
				Tanka:               s.Tanka,
				UriageMaisuu:        s.UriageMaisuu,
				MakanaiKubun:        s.MakanaiKubun,
				UriageKingaku:       0,
			}
		}

		product := categoryGroups[catName][subCatName][menuCode]
		product.UriageKingaku += s.UriageKingaku
		categoryGroups[catName][subCatName][menuCode] = product
	}

	//log.Printf("Category Groups: %+v", categoryGroups)

	// 3. Build the hierarchical GroupedSale structure
	var groupedCategories []GroupedSale
	for catName, subCatMap := range categoryGroups {
		var categoryTotal int
		var groupedSubCategories []GroupedSale

		for subCatName, productList := range subCatMap {
			var subCategoryTotal int
			for _, product := range productList {
				subCategoryTotal += product.UriageKingaku
			}

			// カテゴリー名 Group
			groupedSubCategories = append(groupedSubCategories, GroupedSale{
				Name:     subCatName,
				Total:    subCategoryTotal,
				Products: productList,
			})
			categoryTotal += subCategoryTotal
		}

		// Sort subcategories alphabetically
		sort.Slice(groupedSubCategories, func(i, j int) bool {
			return groupedSubCategories[i].Name < groupedSubCategories[j].Name
		})

		// 部門 Group
		groupedCategories = append(groupedCategories, GroupedSale{
			Name:      catName,
			Total:     categoryTotal,
			SubGroups: groupedSubCategories,
		})
	}

	// Sort categories alphabetically
	sort.Slice(groupedCategories, func(i, j int) bool {
		return groupedCategories[i].Name < groupedCategories[j].Name
	})

	// Grand Total Group
	return GroupedSale{
		Name:      "総計",
		Total:     grandTotal,
		SubGroups: groupedCategories, // Categories are treated as sub-groups of the Grand Total
	}
}

func GetMenubetsuUriage(start_date string, end_date string) (*GroupedSale, error) {

	// Placeholder: Replace with actual DB query logic.

	// Define a slice to hold the query results.
	var sales []SalesDetail

	// Execute the SQL query to fetch all data from the t_メニュー別売上 table.
	rows, err := db.Query(
		"SELECT `ID`, `部門`, `カテゴリー名`, `カテゴリー名讃岐`, `メニューコード`, `メニュー名`, `単価`, `売上金額`, `まかない区分`, `売上枚数`, `日付` FROM `t_メニュー別売上` "+
			"WHERE `日付` BETWEEN ? AND ? ORDER BY `部門`, `カテゴリー名`, `メニューコード`",
		start_date, end_date)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Iterate over the rows and populate the sales slice.
	for rows.Next() {
		var salesDetail SalesDetail
		// Assuming the table columns match the GroupedSale structure.
		if err := rows.Scan(
			&salesDetail.ID,
			&salesDetail.Bumon,
			&salesDetail.CategoryName,
			&salesDetail.CategoryName_Sanuki,
			&salesDetail.MenuCode,
			&salesDetail.MenuName,
			&salesDetail.Tanka,
			&salesDetail.UriageKingaku,
			&salesDetail.MakanaiKubun,
			&salesDetail.UriageMaisuu,
			&salesDetail.Hiduke,
		); err != nil {
			return nil, err
		}
		sales = append(sales, salesDetail)

	}

	processedSales := processSales(sales)

	// Check for any errors encountered during iteration.
	if err := rows.Err(); err != nil {
		return nil, err
	}

	// Marshal the processedSales to JSON and return
	return &processedSales, nil

}

func SaveMenubetsuUriage(date string, products []SoldProduct) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec("DELETE FROM t_メニュー別売上 WHERE 日付 = ?", date)
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare("INSERT INTO t_メニュー別売上 (部門, カテゴリー名, メニューコード, メニュー名, 単価, 売上金額, まかない区分, 売上枚数, 日付) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, p := range products {
		log.Printf("Inserting product: %+v", p)
		_, err := stmt.Exec(p.Bumon, p.Category, p.ProductID, p.MenuName, p.TekiyouKakaku, p.UriageKingaku, p.MakanaiKubun, p.HanbaiMaisuu, date)
		if err != nil {
			log.Printf("Failed to insert product: %+v, error: %v", p, err)
			return err
		}
	}

	return tx.Commit()
}
