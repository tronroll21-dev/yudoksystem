/**
         * 2つのDateオブジェクトが同じ日付を指しているか比較します。
         * 時間は無視します。
         * @param {Date} d1
         * @param {Date} d2
         * @returns {boolean}
         */
const isSameDay = (d1, d2) => {
    return d1.getFullYear() === d2.getFullYear() &&
           d1.getMonth() === d2.getMonth() &&
           d1.getDate() === d2.getDate();
};

/**
 * この関数は、単一のTSVファイルを処理し、集計データを返します。
 * GoのバックエンドロジックをJavaScriptに移植したものです。
 * @param {File} file 処理するファイル
 * @param {string} processDateStr 処理対象日 (YYYY-MM-DD形式)
 * @returns {Promise<object>} 集計結果を含むPromise
 */
async function processFileOnFrontend(file, processDateStr) {
    let colCategoryRngs = [];
    try {
        const response = await fetch('/api/ranges');
        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }
        colCategoryRngs = await response.json();
    } catch (error) {
        console.error("Failed to fetch category ranges:", error);
        throw error;
    }

    return new Promise((resolve, reject) => {
        const reader = new FileReader();

        reader.onload = function(event) {
            try {
                const content = event.target.result;
                const lines = content.split('\n');
                if (lines.length < 2) {
                    throw new Error('ファイルが空であるか、ヘッダーしか含まれていません。');
                }

                // `processDateStr`から日付部分のみを考慮したDateオブジェクトを作成
                const processDate = new Date(processDateStr);
                processDate.setHours(0, 0, 0, 0); // タイムゾーンのずれを防ぐため、時刻をリセット

                const dictSoldProducts = {};
                const paymentStats = {};
                
                const getCategoryName = (bumonID, menuCode) => {
                    for (const elem of colCategoryRngs) {
                        if (elem.bumon_id === bumonID && menuCode >= elem.han_i_kaishi && menuCode <= elem.han_i_shuryo) {
                            return elem.category_name;
                        }
                    }
                    return "";
                };

                const updateSoldProducts = (intBumon, strBumon, hanbaiMaisuu, menuCode, tekiyouKakaku, menuMei, kessaiKingaku, factor) => {
                    const key = `${strBumon}_${String(menuCode).padStart(3, '0')}`;
                    if (dictSoldProducts[key]) {
                        dictSoldProducts[key].uriageKingaku += kessaiKingaku * factor;
                        dictSoldProducts[key].hanbaiMaisuu += hanbaiMaisuu * factor;
                    } else {
                        const newProduct = {
                            bumon: intBumon,
                            bumonName: strBumon,
                            uriageKingaku: kessaiKingaku,
                            hanbaiMaisuu: hanbaiMaisuu,
                            productID: menuCode,
                            menuName: menuMei,
                            tekiyouKakaku: tekiyouKakaku,
                            makanaiKubun: menuMei.includes("まかない") ? "まかない" : "実売上",
                            date: processDateStr,
                            category: getCategoryName(strBumon, menuCode),
                        };
                        dictSoldProducts[key] = newProduct;
                    }
                };
                
                // ヘッダーをスキップ
                for (let i = 1; i < lines.length; i++) {
                    const line = lines[i].trim();
                    if (line === '' || !line.startsWith("2")) {
                        continue;
                    }
                    
                    const fields = line.split(',');
                    
                    const GOUKI = 2;
                    const TORIHIKI_KUBUN = 3;
                    const TORIHIKIKAISHIHIDUKE = 6;
                    const HANBAI_MAISUU = 16;
                    const TEKIYOU_KAKAKU = 18;
                    const KESSAI_KINGAKU = 19;
                    const MENU_CODE = 21;
                    const MENU_MEI = 22;
                    const KAADO_KESSAI = 56;
                    
                    const fileDateStr = fields[TORIHIKIKAISHIHIDUKE];
                    const parsedFileDate = new Date(fileDateStr);
                    if (!isSameDay(parsedFileDate, processDate)) {
                        throw new Error(`ファイルの日付 ${fileDateStr} が選択された日付と一致しません。`);
                    }
                    
                    const vendingmachineNo = parseInt(fields[GOUKI]);
                    const torihikiKubun = parseInt(fields[TORIHIKI_KUBUN]);
                    const kessaiKingaku = parseInt(fields[KESSAI_KINGAKU]);
                    const hanbaiMaisuu = parseInt(fields[HANBAI_MAISUU]);
                    const menuCode = parseInt(fields[MENU_CODE]);
                    const tekiyouKakaku = parseInt(fields[TEKIYOU_KAKAKU]);
                    const menuMei = fields[MENU_MEI];
                    const kaadoKessai = parseInt(fields[KAADO_KESSAI]);
                    
                    let bumon = "";
                    let intBumon = 0;
                    switch (vendingmachineNo) {
                        case 1: case 2: intBumon = 1; bumon = "入浴"; break;
                        case 3: case 4: intBumon = 2; bumon = "飲食"; break;
                        case 5: intBumon = 3; bumon = "エステ"; break;
                    }
                    
                    const factor = (torihikiKubun === 1) ? -1 : 1;
                    updateSoldProducts(intBumon, bumon, hanbaiMaisuu, menuCode, tekiyouKakaku, menuMei, kessaiKingaku, factor);
                    
                    if (!paymentStats[vendingmachineNo]) {
                        paymentStats[vendingmachineNo] = { vendingMachineNo: vendingmachineNo, GenkinGrossMaisuu: 0, GenkinGrossKingaku: 0, GenkinSeisanMaisuu: 0, GenkinSeisanKingaku: 0, QrcodeGrossMaisuu: 0, QrcodeGrossKingaku: 0, QrcodeSeisanMaisuu: 0, QrcodeSeisanKingaku: 0, DenshimaneeGrossMaisuu: 0, DenshimaneeGrossKingaku: 0 };
                    }
                    
                    switch (kaadoKessai) {
                        case 0: // 現金
                            if (factor === 1) {
                                paymentStats[vendingmachineNo].GenkinGrossKingaku += kessaiKingaku;
                                paymentStats[vendingmachineNo].GenkinGrossMaisuu += hanbaiMaisuu;
                            } else {
                                paymentStats[vendingmachineNo].GenkinSeisanKingaku += kessaiKingaku;
                                paymentStats[vendingmachineNo].GenkinSeisanMaisuu += hanbaiMaisuu;
                            }
                            break;
                        case 5: // QRコード
                            if (factor === 1) {
                                paymentStats[vendingmachineNo].QrcodeGrossKingaku += kessaiKingaku;
                                paymentStats[vendingmachineNo].QrcodeGrossMaisuu += hanbaiMaisuu;
                            } else {
                                paymentStats[vendingmachineNo].QrcodeSeisanKingaku += kessaiKingaku;
                                paymentStats[vendingmachineNo].QrcodeSeisanMaisuu += hanbaiMaisuu;
                            }
                            break;
                        case 7: // 電子マネー
                            if (factor === 1) {
                                paymentStats[vendingmachineNo].DenshimaneeGrossKingaku += kessaiKingaku;
                                paymentStats[vendingmachineNo].DenshimaneeGrossMaisuu += hanbaiMaisuu;
                            } else {
                                throw new Error("電子マネーの精算データは存在しないはずです。");
                            }
                            break;
                    }
                }

                // 最終的な集計をまとめる
                const summary = {};
                const getHanbaiMaisuu = (key) => dictSoldProducts[key] ? dictSoldProducts[key].hanbaiMaisuu : 0;
                summary["大人入浴券枚数"] = getHanbaiMaisuu("入浴_051") + getHanbaiMaisuu("入浴_155");
                summary["大人入浴セット券枚数"] = getHanbaiMaisuu("入浴_054");
                summary["小人入浴券枚数"] = getHanbaiMaisuu("入浴_052") + getHanbaiMaisuu("入浴_156");
                summary["感謝祭招待券回収"] = getHanbaiMaisuu("入浴_053") + getHanbaiMaisuu("入浴_157");
                summary["6回数券売数"] = getHanbaiMaisuu("入浴_057");
                summary["回数券売数"] = getHanbaiMaisuu("入浴_055");
                
                resolve({
                    paymentStats,
                    soldProducts: dictSoldProducts,
                    summary
                });
                
            } catch (e) {
                reject(e);
            }
        };

        reader.onerror = function() {
            reject(new Error("ファイルの読み込み中にエラーが発生しました。"));
        };

        reader.readAsText(file, 'Shift_JIS'); // TSVファイルがShift-JISの場合を想定
    });
}


document.addEventListener('alpine:init', () => {
Alpine.data('uploader', () => ({
        formData: {
            date: '',
            files: null,
        },
        get curGetsudoRange() {
            if (!this.formData || !this.formData.date) return '';
            const parts = String(this.formData.date).split('-');
            if (parts.length !== 3) return '';
            const y = parseInt(parts[0], 10);
            const m = parseInt(parts[1], 10) - 1; // 0-based month
            const d = parseInt(parts[2], 10);
            if (Number.isNaN(y) || Number.isNaN(m) || Number.isNaN(d)) return '';

            const today = new Date(y, m, d);
            const year = today.getFullYear();
            const month = today.getMonth(); // 0-based
            const day = today.getDate();
            let target;
            if (day >= 21) {
                target = new Date(year, month, 21);
            } else {
                let prevMonth = month - 1;
                let yy = year;
                if (prevMonth < 0) { prevMonth = 11; yy = year - 1; }
                target = new Date(yy, prevMonth, 21);
            }
            const yyyy = target.getFullYear();
            const mm = String(target.getMonth() + 1).padStart(2, '0');
            const dd = String(target.getDate()).padStart(2, '0');
            const target20Date = new Date(yyyy, target.getMonth() + 1, 20);
            return {mostRecent21:`${yyyy}-${mm}-${dd}`, closest20th: target20Date.getFullYear() + '-' + String(target20Date.getMonth() + 1).padStart(2, '0') + '-' + String(target20Date.getDate()).padStart(2, '0')    };
        },
        get curCalendarMonth() {
            if (!this.formData || !this.formData.date) return '';
            const parts = String(this.formData.date).split('-');
            if (parts.length !== 3) return '';
            const y = parseInt(parts[0], 10);
            const m = parseInt(parts[1], 10) - 1;
            const d = parseInt(parts[2], 10);
            if (Number.isNaN(y) || Number.isNaN(m) || Number.isNaN(d)) return '';   
            const today = new Date(y, m, d);
            const year = today.getFullYear();
            const month = today.getMonth();
            const firstDay = new Date(year, month, 1);
            const lastDay = new Date(year, month + 1, 0);
            const yyyy = firstDay.getFullYear();
            const mm = String(firstDay.getMonth() + 1).padStart(2, '0');
            const dd1 = String(firstDay.getDate()).padStart(2, '0');
            const dd2 = String(lastDay.getDate()).padStart(2, '0');
            return {startDate:`${yyyy}-${mm}-${dd1}`, endDate:`${yyyy}-${mm}-${dd2}`};
        },
        fileStatus: 'ファイルが選択されていません。',
        isLoading: false,
        errorMessage: '',
        successMessage: '',
        responseData: null,

        updateDate(newDate) {
            console.log('Updating date to:', newDate);
            this.formData.date = newDate;
        },

        handleFileChange(event) {
            this.formData.files = event.target.files;
            this.updateFileStatus();
            this.submitForm();
        },
        handleDrop(event) {
            debugger;
            this.formData.files = event.dataTransfer.files;
            this.updateFileStatus();
            this.submitForm();
        },
        updateFileStatus() {
            if (this.formData.files && this.formData.files.length > 0) {
                this.fileStatus = `${this.formData.files.length} ファイルが選択されました。`;
            } else {
                this.fileStatus = 'ファイルが選択されていません。';
            }
        },
        
        async submitForm() {
            this.isLoading = true;
            this.errorMessage = '';
            this.successMessage = '';
            this.responseData = null;

            if (!this.formData.date) {
                this.errorMessage = '日付を選択してください。';
                this.isLoading = false;
                return;
            }

            if (!this.formData.files || this.formData.files.length === 0) {
                this.errorMessage = 'ファイルをアップロードしてください。';
                this.isLoading = false;
                return;
            }
            
            try {
                let combinedData = {
                    paymentStats: {},
                    soldProducts: {},
                    summary: {
                        "大人入浴券枚数": 0,
                        "大人入浴セット券枚数": 0,
                        "小人入浴券枚数": 0,
                        "感謝祭招待券回収": 0,
                        "6回数券売数": 0,
                        "回数券売数": 0,
                    }
                };
                
                for (const file of this.formData.files) {
                    const result = await processFileOnFrontend(file, this.formData.date);
                    
                    // 各ファイルの集計結果を統合
                    for (const key in result.paymentStats) {
                        if (!combinedData.paymentStats[key]) {
                            combinedData.paymentStats[key] = result.paymentStats[key];
                        } else {
                            // 既存のpaymentStatsに新しい値を加算
                            const existing = combinedData.paymentStats[key];
                            const newStats = result.paymentStats[key];
                            existing.GenkinGrossMaisuu += newStats.GenkinGrossMaisuu;
                            existing.GenkinGrossKingaku += newStats.GenkinGrossKingaku;
                            existing.GenkinSeisanMaisuu += newStats.GenkinSeisanMaisuu;
                            existing.GenkinSeisanKingaku += newStats.GenkinSeisanKingaku;
                            existing.QrcodeGrossMaisuu += newStats.QrcodeGrossMaisuu;
                            existing.QrcodeGrossKingaku += newStats.QrcodeGrossKingaku;
                            existing.QrcodeSeisanMaisuu += newStats.QrcodeSeisanMaisuu;
                            existing.QrcodeSeisanKingaku += newStats.QrcodeSeisanKingaku;
                            existing.DenshimaneeGrossMaisuu += newStats.DenshimaneeGrossMaisuu;
                            existing.DenshimaneeGrossKingaku += newStats.DenshimaneeGrossKingaku;
                        }
                    }
                    
                    for (const key in result.soldProducts) {
                        if (!combinedData.soldProducts[key]) {
                            combinedData.soldProducts[key] = result.soldProducts[key];
                        } else {
                            const existing = combinedData.soldProducts[key];
                            const newProduct = result.soldProducts[key];
                            existing.uriageKingaku += newProduct.uriageKingaku;
                            existing.hanbaiMaisuu += newProduct.hanbaiMaisuu;
                        }
                    }
                    
                    for (const key in result.summary) {
                        combinedData.summary[key] += result.summary[key];
                    }
                }

                this.$dispatch('data-updated', { data: combinedData });

                const paymentStatKeys = Object.keys(combinedData.paymentStats).map(Number);
                const areKeysValid = paymentStatKeys.length > 0
                && paymentStatKeys.some(key => key == 1)
                && paymentStatKeys.some(key => key == 2)
                && paymentStatKeys.some(key => key == 3)
                && paymentStatKeys.some(key => key == 4)
                && paymentStatKeys.some(key => key == 5);

                if (areKeysValid && this.formData.files.length === 5) {
                    try {
                        const payload = {
                            date: this.formData.date,
                            soldProducts: Object.values(combinedData.soldProducts)
                        };

                        const response = await fetch('/api/menubetsu-uriage', {
                            method: 'POST',
                            headers: {
                                'Content-Type': 'application/json',
                            },
                            body: JSON.stringify(payload)
                        });

                        if (!response.ok) {
                            const errorData = await response.json();
                            throw new Error(errorData.error || 'メニュー別売上データの更新に失敗しました。');
                        }
                        this.successMessage = 'データが正常に処理され、メニュー別売上データが更新されました！';
                    } catch (error) {
                        console.error('Error sending sold products:', error);
                        this.errorMessage = (this.errorMessage ? this.errorMessage + '\n' : '') + error.message;
                    }
                }  else if (this.formData.files.length !== 5) {
                    this.errorMessage = `５機分のデータファイルが揃っておらず、\n
                    メニュー別売上データの保存はスキップされました。\n\n
                    ５機分のデータファイルが揃ったらアップロードを再度実行してください。`;
                }
                else {
                    this.errorMessage = '同一券売機のデータが複数含まれている様です。';
                }
                
                if (this.errorMessage === '') {
                    this.$dispatch('close-menu');
                }

            } catch (error) {
                console.error('Processing error:', error);
                this.errorMessage = error.message || 'ファイルの処理中に予期せぬエラーが発生しました。';
            } finally {
                this.isLoading = false;
            }
        }
    }))
});