// ============================================================
// Types
// ============================================================
// ============================================================
// Standalone utility functions
// ============================================================
const isSameDay = (d1, d2) => {
    return d1.getFullYear() === d2.getFullYear() &&
        d1.getMonth() === d2.getMonth() &&
        d1.getDate() === d2.getDate();
};
const emptyPaymentStat = () => ({
    vendingMachineNo: 0,
    Machine1CashCount: 0, Machine1CashAmount: 0,
    Machine1SettleCount: 0, Machine1SettleAmount: 0,
    Machine1QrCount: 0, Machine1QrAmount: 0,
    Machine1QrSettleCount: 0, Machine1QrSettleAmount: 0,
    Machine1ECount: 0, Machine1EAmount: 0,
    Machine1ESettleCount: 0, Machine1ESettleAmount: 0,
    Machine1CCount: 0, Machine1CAmount: 0,
    Machine1CSettleCount: 0, Machine1CSettleAmount: 0,
    Machine2CashCount: 0, Machine2CashAmount: 0,
    Machine2SettleCount: 0, Machine2SettleAmount: 0,
    Machine2QrCount: 0, Machine2QrAmount: 0,
    Machine2QrSettleCount: 0, Machine2QrSettleAmount: 0,
    Machine2ECount: 0, Machine2EAmount: 0,
    Machine2ESettleCount: 0, Machine2ESettleAmount: 0,
    Machine2CCount: 0, Machine2CAmount: 0,
    Machine2CSettleCount: 0, Machine2CSettleAmount: 0,
    Machine3CashCount: 0, Machine3CashAmount: 0,
    Machine3SettleCount: 0, Machine3SettleAmount: 0,
    Machine3QrCount: 0, Machine3QrAmount: 0,
    Machine3QrSettleCount: 0, Machine3QrSettleAmount: 0,
    Machine3ECount: 0, Machine3EAmount: 0,
    Machine3ESettleCount: 0, Machine3ESettleAmount: 0,
    Machine3CCount: 0, Machine3CAmount: 0,
    Machine3CSettleCount: 0, Machine3CSettleAmount: 0,
    Machine4CashCount: 0, Machine4CashAmount: 0,
    Machine4SettleCount: 0, Machine4SettleAmount: 0,
    Machine4QrCount: 0, Machine4QrAmount: 0,
    Machine4QrSettleCount: 0, Machine4QrSettleAmount: 0,
    Machine4ECount: 0, Machine4EAmount: 0,
    Machine4ESettleCount: 0, Machine4ESettleAmount: 0,
    Machine4CCount: 0, Machine4CAmount: 0,
    Machine4CSettleCount: 0, Machine4CSettleAmount: 0,
    Machine5CashCount: 0, Machine5CashAmount: 0,
    Machine5SettleCount: 0, Machine5SettleAmount: 0,
    Machine5QrCount: 0, Machine5QrAmount: 0,
    Machine5QrSettleCount: 0, Machine5QrSettleAmount: 0,
    Machine5ECount: 0, Machine5EAmount: 0,
    Machine5ESettleCount: 0, Machine5ESettleAmount: 0,
    Machine5CCount: 0, Machine5CAmount: 0,
    Machine5CSettleCount: 0, Machine5CSettleAmount: 0,
});
const emptysummary = () => ({
    AdultTicketCount: 0,
    AdultSetTicketCount: 0,
    ChildTicketCount: 0,
    InfantTicketCount: 0,
    SixTicketCount: 0,
    TicketCount: 0,
});
const mergeResults = (results) => {
    const merged = {
        paymentStats: emptyPaymentStat(),
        soldProducts: {},
        summary: emptysummary(),
    };
    for (const result of results) {
        // Merge paymentStats — all properties are numeric so we can iterate
        for (const key of Object.keys(result.paymentStats)) {
            if (key === 'vendingMachineNo')
                continue;
            merged.paymentStats[key] += result.paymentStats[key];
        }
        // Merge soldProducts
        for (const key in result.soldProducts) {
            if (!merged.soldProducts[key]) {
                merged.soldProducts[key] = { ...result.soldProducts[key] };
            }
            else {
                merged.soldProducts[key].uriageKingaku += result.soldProducts[key].uriageKingaku;
                merged.soldProducts[key].hanbaiMaisuu += result.soldProducts[key].hanbaiMaisuu;
            }
        }
        // Merge summary
        for (const key of Object.keys(result.summary)) {
            merged.summary[key] += result.summary[key];
        }
    }
    return merged;
};
async function processFileOnFrontend(file, processDateStr, colCategoryRngs) {
    return new Promise((resolve, reject) => {
        const reader = new FileReader();
        reader.onload = (event) => {
            try {
                const content = event.target.result;
                const lines = content.split('\n');
                if (lines.length < 2) {
                    throw new Error('ファイルが空であるか、ヘッダーしか含まれていません。');
                }
                const processDate = new Date(processDateStr);
                processDate.setHours(0, 0, 0, 0);
                const soldProducts = {};
                const paymentStats = emptyPaymentStat();
                const getCategoryName = (bumonID, menuCode) => {
                    for (const elem of colCategoryRngs) {
                        if (elem.bumon_id === bumonID &&
                            menuCode >= elem.han_i_kaishi &&
                            menuCode <= elem.han_i_shuryo)
                            return elem.category_name;
                    }
                    return '';
                };
                const updateSoldProducts = (intBumon, strBumon, hanbaiMaisuu, menuCode, tekiyouKakaku, menuMei, kessaiKingaku, factor) => {
                    const key = `${strBumon}_${String(menuCode).padStart(3, '0')}`;
                    if (soldProducts[key]) {
                        soldProducts[key].uriageKingaku += kessaiKingaku * factor;
                        soldProducts[key].hanbaiMaisuu += hanbaiMaisuu * factor;
                    }
                    else {
                        soldProducts[key] = {
                            bumon: intBumon,
                            bumonName: strBumon,
                            uriageKingaku: kessaiKingaku,
                            hanbaiMaisuu,
                            productID: menuCode,
                            menuName: menuMei,
                            tekiyouKakaku,
                            makanaiKubun: menuMei.includes('まかない') ? 'まかない' : '実売上',
                            date: processDateStr,
                            category: getCategoryName(strBumon, menuCode),
                        };
                    }
                };
                const GOUKI = 2;
                const TORIHIKI_KUBUN = 3;
                const TORIHIKIKAISHIHIDUKE = 6;
                const HANBAI_MAISUU = 16;
                const TEKIYOU_KAKAKU = 18;
                const KESSAI_KINGAKU = 19;
                const MENU_CODE = 21;
                const MENU_MEI = 22;
                const KAADO_KESSAI = 56;
                for (let i = 1; i < lines.length; i++) {
                    const line = lines[i].trim();
                    if (line === '' || !line.startsWith('2'))
                        continue;
                    const fields = line.split(',');
                    const fileDateStr = fields[TORIHIKIKAISHIHIDUKE];
                    const parsedFileDate = new Date(fileDateStr);
                    if (!isSameDay(parsedFileDate, processDate)) {
                        throw new Error(`ファイルの日付 ${fileDateStr} が選択された日付と一致しません。`);
                    }
                    const vendingmachineNo = parseInt(fields[GOUKI], 10);
                    const torihikiKubun = parseInt(fields[TORIHIKI_KUBUN], 10);
                    const kessaiKingaku = parseInt(fields[KESSAI_KINGAKU], 10);
                    const hanbaiMaisuu = parseInt(fields[HANBAI_MAISUU], 10);
                    const menuCode = parseInt(fields[MENU_CODE], 10);
                    const tekiyouKakaku = parseInt(fields[TEKIYOU_KAKAKU], 10);
                    const menuMei = fields[MENU_MEI];
                    const kaadoKessai = parseInt(fields[KAADO_KESSAI], 10);
                    let bumon = '';
                    let intBumon = 0;
                    switch (vendingmachineNo) {
                        case 1:
                        case 2:
                            intBumon = 1;
                            bumon = '入浴';
                            break;
                        case 3:
                        case 4:
                            intBumon = 2;
                            bumon = '飲食';
                            break;
                        case 5:
                            intBumon = 3;
                            bumon = 'エステ';
                            break;
                    }
                    const factor = torihikiKubun === 1 ? -1 : 1;
                    updateSoldProducts(intBumon, bumon, hanbaiMaisuu, menuCode, tekiyouKakaku, menuMei, kessaiKingaku, factor);
                    // Use the vending machine number as a prefix for PaymentStat keys
                    const m = `Machine${vendingmachineNo}`;
                    switch (kaadoKessai) {
                        case 0: // 現金
                            if (factor === 1) {
                                paymentStats[`${m}CashAmount`] += kessaiKingaku;
                                paymentStats[`${m}CashCount`] += hanbaiMaisuu;
                            }
                            else {
                                paymentStats[`${m}SettleAmount`] += kessaiKingaku;
                                paymentStats[`${m}SettleCount`] += hanbaiMaisuu;
                            }
                            break;
                        case 5: // QRコード
                            if (factor === 1) {
                                paymentStats[`${m}QrAmount`] += kessaiKingaku;
                                paymentStats[`${m}QrCount`] += hanbaiMaisuu;
                            }
                            else {
                                paymentStats[`${m}QrSettleAmount`] += kessaiKingaku;
                                paymentStats[`${m}QrSettleCount`] += hanbaiMaisuu;
                            }
                            break;
                        case 6: // クレジットカード
                            if (factor === 1) {
                                paymentStats[`${m}CAmount`] += kessaiKingaku;
                                paymentStats[`${m}CCount`] += hanbaiMaisuu;
                            }
                            else {
                                paymentStats[`${m}CSettleAmount`] += kessaiKingaku;
                                paymentStats[`${m}CSettleCount`] += hanbaiMaisuu;
                            }
                            break;
                        case 7: // 電子マネー
                            if (factor === 1) {
                                paymentStats[`${m}EAmount`] += kessaiKingaku;
                                paymentStats[`${m}ECount`] += hanbaiMaisuu;
                            }
                            else {
                                throw new Error('電子マネーの精算データは存在しないはずです。');
                            }
                            break;
                    }
                }
                const getHanbaiMaisuu = (key) => soldProducts[key]?.hanbaiMaisuu ?? 0;
                const summary = {
                    AdultTicketCount: getHanbaiMaisuu('入浴_051') + getHanbaiMaisuu('入浴_155'),
                    AdultSetTicketCount: getHanbaiMaisuu('入浴_054'),
                    ChildTicketCount: getHanbaiMaisuu('入浴_052') + getHanbaiMaisuu('入浴_156'),
                    InfantTicketCount: getHanbaiMaisuu('入浴_053') + getHanbaiMaisuu('入浴_157'),
                    SixTicketCount: getHanbaiMaisuu('入浴_057'),
                    TicketCount: getHanbaiMaisuu('入浴_055'),
                };
                resolve({ paymentStats, soldProducts, summary });
            }
            catch (e) {
                reject(e);
            }
        };
        reader.onerror = () => reject(new Error('ファイルの読み込み中にエラーが発生しました。'));
        reader.readAsText(file, 'Shift_JIS');
    });
}
// ============================================================
// Alpine.js component
// ============================================================
document.addEventListener('alpine:init', () => {
    const Alpine = window.Alpine;
    Alpine.data('dailyReport', () => ({
        formData: {
            date: '',
            files: null,
        },
        data: {
            Record: { ID: 0,
                DateString: '',
                WeatherCode: 0,
                StaffCode: 0,
                Machine1CashCount: 0,
                Machine1CashAmount: 0,
                Machine1SettleCount: 0,
                Machine1SettleAmount: 0,
                Machine2CashCount: 0,
                Machine2CashAmount: 0,
                Machine2SettleCount: 0,
                Machine2SettleAmount: 0,
                Machine3CashCount: 0,
                Machine3CashAmount: 0,
                Machine3SettleCount: 0,
                Machine3SettleAmount: 0,
                Machine4CashCount: 0,
                Machine4CashAmount: 0,
                Machine4SettleCount: 0,
                Machine4SettleAmount: 0,
                Machine5CashCount: 0,
                Machine5CashAmount: 0,
                Machine5SettleCount: 0,
                Machine5SettleAmount: 0,
                Machine1UnsettledCount: 0,
                Machine1UnsettledAmount: 0,
                Machine2UnsettledCount: 0,
                Machine2UnsettledAmount: 0,
                Machine3UnsettledCount: 0,
                Machine3UnsettledAmount: 0,
                Machine4UnsettledCount: 0,
                Machine4UnsettledAmount: 0,
                Machine5UnsettledCount: 0,
                Machine5UnsettledAmount: 0,
                Machine1QrCount: 0,
                Machine1QrAmount: 0,
                Machine2QrCount: 0,
                Machine2QrAmount: 0,
                Machine3QrCount: 0,
                Machine3QrAmount: 0,
                Machine4QrCount: 0,
                Machine4QrAmount: 0,
                Machine5QrCount: 0,
                Machine5QrAmount: 0,
                Machine1QrSettleCount: 0,
                Machine1QrSettleAmount: 0,
                Machine2QrSettleCount: 0,
                Machine2QrSettleAmount: 0,
                Machine3QrSettleCount: 0,
                Machine3QrSettleAmount: 0,
                Machine4QrSettleCount: 0,
                Machine4QrSettleAmount: 0,
                Machine5QrSettleCount: 0,
                Machine5QrSettleAmount: 0,
                Machine1ECount: 0,
                Machine1EAmount: 0,
                Machine2ECount: 0,
                Machine2EAmount: 0,
                Machine3ECount: 0,
                Machine3EAmount: 0,
                Machine4ECount: 0,
                Machine4EAmount: 0,
                Machine5ECount: 0,
                Machine5EAmount: 0,
                Machine1ESettleCount: 0,
                Machine1ESettleAmount: 0,
                Machine2ESettleCount: 0,
                Machine2ESettleAmount: 0,
                Machine3ESettleCount: 0,
                Machine3ESettleAmount: 0,
                Machine4ESettleCount: 0,
                Machine4ESettleAmount: 0,
                Machine5ESettleCount: 0,
                Machine5ESettleAmount: 0,
                Machine1CCount: 0,
                Machine1CAmount: 0,
                Machine2CCount: 0,
                Machine2CAmount: 0,
                Machine3CCount: 0,
                Machine3CAmount: 0,
                Machine4CCount: 0,
                Machine4CAmount: 0,
                Machine5CCount: 0,
                Machine5CAmount: 0,
                Machine1CSettleCount: 0,
                Machine1CSettleAmount: 0,
                Machine2CSettleCount: 0,
                Machine2CSettleAmount: 0,
                Machine3CSettleCount: 0,
                Machine3CSettleAmount: 0,
                Machine4CSettleCount: 0,
                Machine4CSettleAmount: 0,
                Machine5CSettleCount: 0,
                Machine5CSettleAmount: 0,
                AdultTicketCount: 0,
                AdultSetTicketCount: 0,
                ChildTicketCount: 0,
                InfantTicketCount: 0,
                TicketCount: 0,
                SixTicketCount: 0,
                MaleTicketCount: 0,
                FemaleTicketCount: 0,
                MaleTicketShare: '',
                FemaleTicketShare: '',
                InvitationTicketCount: 0,
                CourtesyTicketCount: 0,
                ThanksgivingTicketCount: 0,
                PointCardAdultCount: 0,
                PointCardChildCount: 0,
                TicketSalesCount: 0,
                OldTicketCount: 0,
                SalesNoStart: 0,
                SalesNoEnd: 0,
                Change: 0,
                PhoneFee: 0,
                CourtesySalesCount: 0,
                CourtesySalesAmount: 0,
                HonjitsuMitounyuuAmountUncertain: 0,
                HonjitsuMitounyuuAmountCertain: 0,
                Deficiency: 0,
                ZenjitsuMitounyuuAmount: 0,
                ReceiptTotalAmount: 0,
                Remarks: '',
                EscortMale: 0,
                EscortFemale: 0,
                EscortChild: 0,
                SixTicketSalesCount: 0,
                SixSalesNoStart: 0,
                SixSalesNoEnd: 0,
                SixMaleTicketCount: 0,
                SixFemaleTicketCount: 0,
                SCutMale: 0,
                SCutFemale: 0,
                SCutChild: 0,
                CouponCount: 0,
                RearRegisterAmount: 0,
                RearRegisterTicketAmount: 0,
                RearRegisterRelaxAmount: 0,
                ReportSpace: '' },
            Color: '',
            Found: false,
            Mode: '',
        },
        toast: {
            show: false,
            message: '',
            type: 'success',
        },
        tenkiOptions: [
            'なし', '晴れ', '曇り', '雨',
            '晴れのち曇り', '曇りのち雨', '雨のち曇り', '雪', '雨台風',
        ],
        showModal: false,
        selectedDate: '',
        currentMonth: new Date().getMonth(),
        currentYear: new Date().getFullYear(),
        loading: true,
        error: null,
        tantoushas: {},
        userId: null,
        api_url: '/api/sales-data?date=',
        api_tantousha_url: '/api/tantoushas',
        fileStatus: 'ファイルが選択されていません。',
        isLoading: false,
        errorMessage: '',
        successMessage: '',
        responseData: null,
        // ---- computed getters ----
        get tenkiIconUrl() {
            if (this.data?.Record?.WeatherCode !== undefined) {
                const code = this.data.Record.WeatherCode;
                if (code >= 1)
                    return `/assets/img/${this.tenkiOptions[code]}.svg`;
                return 'data:image/gif;base64,R0lGODlhAQABAIAAAP///wAAACH5BAEAAAAALAAAAAABAAEAAAICRAEAOw==';
            }
            return '';
        },
        get currentMonthYear() {
            return `${this.currentYear}年${this.currentMonth + 1}月`;
        },
        get calendarDays() {
            const days = [];
            const firstDay = new Date(this.currentYear, this.currentMonth, 1).getDay();
            const daysInMonth = new Date(this.currentYear, this.currentMonth + 1, 0).getDate();
            for (let i = 0; i < firstDay; i++)
                days.push({ date: null, label: '' });
            for (let i = 1; i <= daysInMonth; i++) {
                days.push({ date: new Date(this.currentYear, this.currentMonth, i), label: i });
            }
            return days;
        },
        get selectedDateDisplay() {
            if (!this.selectedDate)
                return '日付を選択';
            const [year, month, day] = this.selectedDate.split('-');
            return `${year}年${parseInt(month)}月${parseInt(day)}日`;
        },
        get curGetsudoRange() {
            if (!this.formData?.date)
                return '';
            const parts = String(this.formData.date).split('-');
            if (parts.length !== 3)
                return '';
            const y = parseInt(parts[0], 10);
            const m = parseInt(parts[1], 10) - 1;
            const d = parseInt(parts[2], 10);
            if (Number.isNaN(y) || Number.isNaN(m) || Number.isNaN(d))
                return '';
            const today = new Date(y, m, d);
            const year = today.getFullYear();
            const month = today.getMonth();
            const day = today.getDate();
            let target;
            if (day >= 21) {
                target = new Date(year, month, 21);
            }
            else {
                let prevMonth = month - 1;
                let yy = year;
                if (prevMonth < 0) {
                    prevMonth = 11;
                    yy = year - 1;
                }
                target = new Date(yy, prevMonth, 21);
            }
            const yyyy = target.getFullYear();
            const mm = String(target.getMonth() + 1).padStart(2, '0');
            const dd = String(target.getDate()).padStart(2, '0');
            const target20Date = new Date(yyyy, target.getMonth() + 1, 20);
            return {
                mostRecent21: `${yyyy}-${mm}-${dd}`,
                closest20th: `${target20Date.getFullYear()}-${String(target20Date.getMonth() + 1).padStart(2, '0')}-${String(target20Date.getDate()).padStart(2, '0')}`,
            };
        },
        get Totals() {
            if (!this.data?.Record)
                return {};
            const rec = this.data.Record;
            let cashCountTotal = 0, cashAmountTotal = 0, settleCountTotal = 0, settleAmountTotal = 0, unsettledCountTotal = 0, unsettledAmountTotal = 0, qrCountTotal = 0, qrAmountTotal = 0, qrSettleCountTotal = 0, qrSettleAmountTotal = 0, eCountTotal = 0, eAmountTotal = 0, eSettleCountTotal = 0, eSettleAmountTotal = 0, cCountTotal = 0, cAmountTotal = 0, cSettleCountTotal = 0, cSettleAmountTotal = 0, netCashCountTotal = 0, netCashAmountTotal = 0, netCashCountMinusUnsettledTotal = 0, netCashAmountMinusUnsettledTotal = 0, netECountTotal = 0, netEAmountTotal = 0, netCCountTotal = 0, netCAmountTotal = 0, netQrCountTotal = 0, netQrAmountTotal = 0;
            const totals = {};
            for (let i = 1; i <= 5; i++) {
                const n = (key) => Number(rec[key]) || 0;
                const cc = n(`Machine${i}CashCount`);
                const ca = n(`Machine${i}CashAmount`);
                const sc = n(`Machine${i}SettleCount`);
                const sa = n(`Machine${i}SettleAmount`);
                const uc = n(`Machine${i}UnsettledCount`);
                const ua = n(`Machine${i}UnsettledAmount`);
                const qrc = n(`Machine${i}QrCount`);
                const qra = n(`Machine${i}QrAmount`);
                const qrsc = n(`Machine${i}QrSettleCount`);
                const qrsa = n(`Machine${i}QrSettleAmount`);
                const ec = n(`Machine${i}ECount`);
                const ea = n(`Machine${i}EAmount`);
                const esc = n(`Machine${i}ESettleCount`);
                const esa = n(`Machine${i}ESettleAmount`);
                const ccc = n(`Machine${i}CCount`);
                const cca = n(`Machine${i}CAmount`);
                const csc = n(`Machine${i}CSettleCount`);
                const csa = n(`Machine${i}CSettleAmount`);
                cashCountTotal += cc;
                cashAmountTotal += ca;
                settleCountTotal += sc;
                settleAmountTotal += sa;
                unsettledCountTotal += uc;
                unsettledAmountTotal += ua;
                qrCountTotal += qrc;
                qrAmountTotal += qra;
                qrSettleCountTotal += qrsc;
                qrSettleAmountTotal += qrsa;
                eCountTotal += ec;
                eAmountTotal += ea;
                eSettleCountTotal += esc;
                eSettleAmountTotal += esa;
                cCountTotal += ccc;
                cAmountTotal += cca;
                cSettleCountTotal += csc;
                cSettleAmountTotal += csa;
                netCashCountTotal += cc - sc;
                netCashAmountTotal += ca - sa;
                netCashCountMinusUnsettledTotal += cc - sc - uc;
                netCashAmountMinusUnsettledTotal += ca - sa - ua;
                netECountTotal += ec - esc;
                netEAmountTotal += ea - esa;
                netCCountTotal += ccc - csc;
                netCAmountTotal += cca - csa;
                netQrCountTotal += qrc - qrsc;
                netQrAmountTotal += qra - qrsa;
                totals[`Machine${i}NetCashCount`] = (cc - sc).toLocaleString('ja-JP');
                totals[`Machine${i}NetCashAmount`] = (ca - sa).toLocaleString('ja-JP');
                totals[`Machine${i}NetCashCountMinusUnsettled`] = (cc - sc - uc).toLocaleString('ja-JP');
                totals[`Machine${i}NetCashAmountMinusUnsettled`] = (ca - sa - ua).toLocaleString('ja-JP');
                totals[`Machine${i}NetECount`] = (ec - esc).toLocaleString('ja-JP');
                totals[`Machine${i}NetEAmount`] = (ea - esa).toLocaleString('ja-JP');
                totals[`Machine${i}NetCCount`] = (ccc - csc).toLocaleString('ja-JP');
                totals[`Machine${i}NetCAmount`] = (cca - csa).toLocaleString('ja-JP');
                totals[`Machine${i}NetQrCount`] = (qrc - qrsc).toLocaleString('ja-JP');
                totals[`Machine${i}NetQrAmount`] = (qra - qrsa).toLocaleString('ja-JP');
                totals[`Machine${i}CashlessTotal`] = (qra - qrsa + ea - esa).toLocaleString('ja-JP');
            }
            totals.CashCountTotal = cashCountTotal.toLocaleString('ja-JP');
            totals.CashAmountTotal = cashAmountTotal.toLocaleString('ja-JP');
            totals.SettleCountTotal = settleCountTotal.toLocaleString('ja-JP');
            totals.SettleAmountTotal = settleAmountTotal.toLocaleString('ja-JP');
            totals.UnsettledCountTotal = unsettledCountTotal.toLocaleString('ja-JP');
            totals.UnsettledAmountTotal = unsettledAmountTotal.toLocaleString('ja-JP');
            totals.QrCountTotal = qrCountTotal.toLocaleString('ja-JP');
            totals.QrAmountTotal = qrAmountTotal.toLocaleString('ja-JP');
            totals.QrSettleCountTotal = qrSettleCountTotal.toLocaleString('ja-JP');
            totals.QrSettleAmountTotal = qrSettleAmountTotal.toLocaleString('ja-JP');
            totals.ECountTotal = eCountTotal.toLocaleString('ja-JP');
            totals.EAmountTotal = eAmountTotal.toLocaleString('ja-JP');
            totals.ESettleCountTotal = eSettleCountTotal.toLocaleString('ja-JP');
            totals.ESettleAmountTotal = eSettleAmountTotal.toLocaleString('ja-JP');
            totals.CCountTotal = cCountTotal.toLocaleString('ja-JP');
            totals.CAmountTotal = cAmountTotal.toLocaleString('ja-JP');
            totals.CSettleCountTotal = cSettleCountTotal.toLocaleString('ja-JP');
            totals.CSettleAmountTotal = cSettleAmountTotal.toLocaleString('ja-JP');
            totals.NetCashCountTotal = netCashCountTotal.toLocaleString('ja-JP');
            totals.NetCashAmountTotal = netCashAmountTotal.toLocaleString('ja-JP');
            totals.NetCashCountMinusUnsettledTotal = netCashCountMinusUnsettledTotal.toLocaleString('ja-JP');
            totals.NetCashAmountMinusUnsettledTotal = netCashAmountMinusUnsettledTotal.toLocaleString('ja-JP');
            totals.NetECountTotal = netECountTotal.toLocaleString('ja-JP');
            totals.NetEAmountTotal = netEAmountTotal.toLocaleString('ja-JP');
            totals.NetQrCountTotal = netQrCountTotal.toLocaleString('ja-JP');
            totals.NetQrAmountTotal = netQrAmountTotal.toLocaleString('ja-JP');
            if (this.data) {
                totals.TounyuuGoukei = (netCashAmountTotal +
                    Number(rec.Change) -
                    (Number(rec.Machine1UnsettledAmount) +
                        Number(rec.Machine2UnsettledAmount) +
                        Number(rec.Machine3UnsettledAmount) +
                        Number(rec.Machine4UnsettledAmount) +
                        Number(rec.Machine5UnsettledAmount)) +
                    Number(rec.PhoneFee) -
                    Number(rec.HonjitsuMitounyuuAmountUncertain) -
                    Number(rec.HonjitsuMitounyuuAmountCertain) +
                    Number(rec.Deficiency) +
                    Number(rec.ZenjitsuMitounyuuAmount)).toLocaleString('ja-JP');
            }
            return totals;
        },
        get TicketTotals() {
            if (!this.data?.Record)
                return {};
            const rec = this.data.Record;
            const maleTotal = Number(rec.MaleTicketCount) + Number(rec.SixMaleTicketCount);
            const femaleTotal = Number(rec.FemaleTicketCount) + Number(rec.SixFemaleTicketCount);
            const grandTotal = maleTotal + femaleTotal;
            return {
                TicketCountTotal: (Number(rec.MaleTicketCount) + Number(rec.FemaleTicketCount)).toLocaleString('ja-JP'),
                SixTicketCountTotal: (Number(rec.SixMaleTicketCount) + Number(rec.SixFemaleTicketCount)).toLocaleString('ja-JP'),
                MaleTicketCountTotal: maleTotal.toLocaleString('ja-JP'),
                FemaleTicketCountTotal: femaleTotal.toLocaleString('ja-JP'),
                TicketCountGrandTotal: grandTotal.toLocaleString('ja-JP'),
                MaleTicketShare: Math.round(maleTotal / grandTotal * 100),
                FemaleTicketShare: Math.round(femaleTotal / grandTotal * 100),
                SCutTotal: (Number(rec.SCutMale) + Number(rec.SCutFemale) + Number(rec.SCutChild)).toLocaleString('ja-JP'),
            };
        },
        // ---- methods ----
        prevNextDate(offset) {
            const date = new Date(this.selectedDate);
            const next_date = new Date(date.setDate(date.getDate() + offset));
            this.selectedDate = `${next_date.getFullYear()}-${String(next_date.getMonth() + 1).padStart(2, '0')}-${String(next_date.getDate()).padStart(2, '0')}`;
            this.$dispatch('update-date', { data: this.selectedDate });
            this.fetchData(this.selectedDate);
        },
        focusNextTabbable(element) {
            document.querySelector(`[tabindex="${element.tabIndex + 1}"]`)?.focus();
        },
        showToast(message, type = 'success') {
            this.toast.message = message;
            this.toast.type = type;
            this.toast.show = true;
            setTimeout(() => { this.toast.show = false; }, 2000);
        },
        isSelected(date) {
            if (!this.selectedDate)
                return false;
            const sel = new Date(this.selectedDate);
            return date.getFullYear() === sel.getFullYear() &&
                date.getMonth() === sel.getMonth() &&
                date.getDate() === sel.getDate();
        },
        isToday(date) {
            const today = new Date();
            return date.getFullYear() === today.getFullYear() &&
                date.getMonth() === today.getMonth() &&
                date.getDate() === today.getDate();
        },
        selectDate(date) {
            const year = date.getFullYear();
            const month = (date.getMonth() + 1).toString().padStart(2, '0');
            const day = date.getDate().toString().padStart(2, '0');
            this.selectedDate = `${year}-${month}-${day}`;
            this.showModal = false;
            this.$nextTick(() => {
                this.fetchData(this.selectedDate);
                if (this.data)
                    this.data.Record.DateString = this.selectedDate;
                this.$dispatch('update-date', { data: this.selectedDate });
            });
        },
        previousMonth() {
            if (this.currentMonth === 0) {
                this.currentMonth = 11;
                this.currentYear--;
            }
            else
                this.currentMonth--;
        },
        nextMonth() {
            if (this.currentMonth === 11) {
                this.currentMonth = 0;
                this.currentYear++;
            }
            else
                this.currentMonth++;
        },
        async init() {
            const today = new Date();
            const year = today.getFullYear();
            const month = (today.getMonth() + 1).toString().padStart(2, '0');
            const day = today.getDate().toString().padStart(2, '0');
            this.selectedDate = `${year}-${month}-${day}`;
            if (this.data)
                this.data.Record.DateString = this.selectedDate;
            this.$dispatch('update-date', { data: this.selectedDate });
            this.fetchTantoushas();
            await this.fetchCurrentUser();
            this.fetchData(this.selectedDate);
        },
        async fetchCurrentUser() {
            try {
                const response = await fetch('/api/user/me', { credentials: 'include' });
                if (!response.ok)
                    throw new Error('Not authenticated');
                const userData = await response.json();
                if (userData.username) {
                    const currentUser = Object.values(this.tantoushas).find((t) => t.name === userData.username);
                    if (currentUser?.id) {
                        this.userId = currentUser.id;
                    }
                    else {
                        throw new Error('User ID not found in tantoushas list');
                    }
                }
                else {
                    throw new Error('Username not found in response');
                }
            }
            catch (e) {
                console.error('Failed to fetch user data, redirecting to login.', e);
                window.location.href = '/login';
            }
        },
        async fetchTantoushas() {
            this.loading = true;
            this.error = null;
            try {
                const response = await fetch(this.api_tantousha_url, { credentials: 'include' });
                if (!response.ok)
                    throw new Error('ネットワーク応答が正常ではありませんでした');
                const raw = await response.json();
                this.tantoushas = JSON.parse(raw);
            }
            catch (e) {
                const err = e;
                this.error = `データの取得に失敗しました: ${err.message}`;
                console.error('Error fetching data:', e);
            }
            finally {
                this.loading = false;
            }
        },
        async fetchData(date) {
            this.loading = true;
            this.error = null;
            try {
                const response = await fetch(`${this.api_url}${date}`, { credentials: 'include' });
                if (!response.ok)
                    throw new Error('ネットワーク応答が正常ではありませんでした');
                const rawData = await response.json();
                if (rawData.Found && rawData.Record.DateString !== date) {
                    console.error('Fetched record date does not match requested date.');
                }
                if (!rawData.Found)
                    rawData.Record.DateString = date;
                this.data = rawData;
                if (!this.data.Found && this.data.Mode === '登録' && this.userId !== null) {
                    this.data.Record.StaffCode = this.userId;
                }
            }
            catch (e) {
                const err = e;
                this.error = `データの取得に失敗しました: ${err.message}`;
                console.error('Error fetching data:', e);
            }
            finally {
                this.loading = false;
            }
        },
        async saveRecord() {
            try {
                const recordToSend = JSON.parse(JSON.stringify(this.data.Record));
                const response = await fetch('/api/sales-data', {
                    method: 'POST',
                    credentials: 'include',
                    headers: { 'Content-Type': 'application/json' },
                    body: JSON.stringify({ Record: recordToSend, Found: this.data.Found }),
                });
                if (!response.ok) {
                    throw new Error(`サーバーエラー: ${response.status} ${response.statusText}`);
                }
                const responseData = await response.json();
                this.data = responseData;
                this.showToast('レコードが正常に保存されました！', 'success');
            }
            catch (error) {
                const err = error;
                console.error('レコードの保存中にエラーが発生しました:', err);
                this.showToast(`レコードの保存に失敗しました: ${err.message}`, 'error');
            }
        },
        handleFileChange(event) {
            this.formData.files = event.target.files;
            this.updateFileStatus();
            this.submitForm();
        },
        handleDrop(event) {
            this.formData.files = event.dataTransfer.files;
            this.updateFileStatus();
            this.submitForm();
        },
        updateFileStatus() {
            this.fileStatus = this.formData.files && this.formData.files.length > 0
                ? `${this.formData.files.length} ファイルが選択されました。`
                : 'ファイルが選択されていません。';
        },
        async submitForm() {
            this.isLoading = true;
            this.errorMessage = '';
            this.successMessage = '';
            this.responseData = null;
            if (!this.selectedDate) {
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
                // Fetch ranges once for all files
                let colCategoryRngs = [];
                try {
                    const rangesResponse = await fetch('/api/ranges');
                    if (!rangesResponse.ok)
                        throw new Error(`HTTP error! status: ${rangesResponse.status}`);
                    colCategoryRngs = await rangesResponse.json();
                }
                catch (error) {
                    console.error('Failed to fetch category ranges:', error);
                    throw error;
                }
                const results = await Promise.all(Array.from(this.formData.files).map((file) => processFileOnFrontend(file, this.selectedDate, colCategoryRngs) // ← passed in
                ));
                const combined = mergeResults(results);
                // Write paymentStats fields directly onto the Record
                Object.assign(this.data.Record, combined.paymentStats);
                // Write summary fields directly onto the Record
                Object.assign(this.data.Record, combined.summary);
                const machinesPresent = new Set(results.map((_, idx) => idx + 1));
                const allFiveMachinesPresent = this.formData.files.length === 5 &&
                    [1, 2, 3, 4, 5].every((k) => machinesPresent.has(k));
                if (allFiveMachinesPresent) {
                    try {
                        const payload = {
                            date: this.selectedDate,
                            soldProducts: Object.values(combined.soldProducts),
                        };
                        const response = await fetch('/api/menubetsu-uriage', {
                            method: 'POST',
                            headers: { 'Content-Type': 'application/json' },
                            body: JSON.stringify(payload),
                        });
                        if (!response.ok) {
                            const errorData = await response.json();
                            throw new Error(errorData.error ?? 'メニュー別売上データの更新に失敗しました。');
                        }
                        this.successMessage = 'データが正常に処理され、メニュー別売上データが更新されました！';
                    }
                    catch (error) {
                        const err = error;
                        console.error('Error sending sold products:', err);
                        this.errorMessage = (this.errorMessage ? this.errorMessage + '\n' : '') + err.message;
                    }
                }
                else if (this.formData.files.length !== 5) {
                    this.errorMessage = `５機分のデータファイルが揃っておらず、\nメニュー別売上データの保存はスキップされました。\n\n５機分のデータファイルが揃ったらアップロードを再度実行してください。`;
                }
                else {
                    this.errorMessage = '同一券売機のデータが複数含まれている様です。';
                }
                if (this.errorMessage === '')
                    this.$dispatch('close-menu');
            }
            catch (error) {
                const err = error;
                console.error('Processing error:', err);
                this.errorMessage = err.message ?? 'ファイルの処理中に予期せぬエラーが発生しました。';
            }
            finally {
                this.isLoading = false;
            }
        },
    }));
});
export {};
//# sourceMappingURL=dailyReport.js.map