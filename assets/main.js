document.addEventListener('alpine:init', () => {
    Alpine.data('dailyReport', () => ({
        data: {
            Record: {
                ID: 0,
                Date: '',
                WeatherCode: 0,
                WeatherStatus: '',
                WeatherMark: '',
                StaffCode: 0,
                StaffName: '',
                Machine1CashCount: '0',
                Machine1CashAmount: '0',
                Machine1SettleCount: '0',
                Machine1SettleAmount: '0',
                Machine2CashCount: '0',
                Machine2CashAmount: '0',
                Machine2SettleCount: '0',
                Machine2SettleAmount: '0',
                Machine3CashCount: '0',
                Machine3CashAmount: '0',
                Machine3SettleCount: '0',
                Machine3SettleAmount: '0',
                Machine4CashCount: '0',
                Machine4CashAmount: '0',
                Machine4SettleCount: '0',
                Machine4SettleAmount: '0',
                Machine5CashCount: '0',
                Machine5CashAmount: '0',
                Machine5SettleCount: '0',
                Machine5SettleAmount: '0',
                Machine1UnsettledCount: '0',
                Machine1UnsettledAmount: '0',
                Machine2UnsettledCount: '0',
                Machine2UnsettledAmount: '0',
                Machine3UnsettledCount: '0',
                Machine3UnsettledAmount: '0',
                Machine4UnsettledCount: '0',
                Machine4UnsettledAmount: '0',
                Machine5UnsettledCount: '0',
                Machine5UnsettledAmount: '0',
                Machine1QrCount: '0',
                Machine1QrAmount: '0',
                Machine2QrCount: '0',
                Machine2QrAmount: '0',
                Machine3QrCount: '0',
                Machine3QrAmount: '0',
                Machine4QrCount: '0',
                Machine4QrAmount: '0',
                Machine5QrCount: '0',
                Machine5QrAmount: '0',
                Machine1QrSettleCount: '0',
                Machine1QrSettleAmount: '0',
                Machine2QrSettleCount: '0',
                Machine2QrSettleAmount: '0',
                Machine3QrSettleCount: '0',
                Machine3QrSettleAmount: '0',
                Machine4QrSettleCount: '0',
                Machine4QrSettleAmount: '0',
                Machine5QrSettleCount: '0',
                Machine5QrSettleAmount: '0',
                Machine1ECount: '0',
                Machine1EAmount: '0',
                Machine2ECount: '0',
                Machine2EAmount: '0',
                Machine3ECount: '0',
                Machine3EAmount: '0',
                Machine4ECount: '0',
                Machine4EAmount: '0',
                Machine5ECount: '0',
                Machine5EAmount: '0',
                Machine1ESettleCount: '0',
                Machine1ESettleAmount: '0',
                Machine2ESettleCount: '0',
                Machine2ESettleAmount: '0',
                Machine3ESettleCount: '0',
                Machine3ESettleAmount: '0',
                Machine4ESettleCount: '0',
                Machine4ESettleAmount: '0',
                Machine5ESettleCount: '0',
                Machine5ESettleAmount: '0',
                Machine1CCount: '0',
                Machine1CAmount: '0',
                Machine2CCount: '0',
                Machine2CAmount: '0',
                Machine3CCount: '0',
                Machine3CAmount: '0',
                Machine4CCount: '0',
                Machine4CAmount: '0',
                Machine5CCount: '0',
                Machine5CAmount: '0',
                Machine1CSettleCount: '0',
                Machine1CSettleAmount: '0',
                Machine2CSettleCount: '0',
                Machine2CSettleAmount: '0',
                Machine3CSettleCount: '0',
                Machine3CSettleAmount: '0',
                Machine4CSettleCount: '0',
                Machine4CSettleAmount: '0',
                Machine5CSettleCount: '0',
                Machine5CSettleAmount: '0',
                AdultTicketCount: 0,
                AdultSetTicketCount: 0,
                ChildTicketCount: 0,
                TicketCount: 0,
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
                TodayUn投入AmountUncertain: 0,
                TodayUn投入AmountCertain: 0,
                Deficiency: 0,
                YesterdayUn投入Amount: 0,
                ReceiptTotalAmount: 0,
                Remarks: '',
                EscortMale: 0,
                EscortFemale: 0,
                EscortChild: 0,
                SixTicketCount: 0,
                SixTicketSalesCount: 0,
                SixSalesNoStart: 0,
                SixSalesNoEnd: 0,
                SixMaleTicketCount: 0,
                SixFemaleTicketCount: 0,
                CouponCount: 0,
                RearRegisterAmount: 0,
                RearRegisterTicketAmount: 0,
                RearRegisterRelaxAmount: 0,
                ReportSpace: '',
            },
            Color: 'green',
            Found: false,
            Mode: '登録'
        },
        selectedDate: '2024-07-26', // Initialize with the default date
        loading: true,
        error: null,
        // フェッチするデータのURL。実際のサーバーエンドポイントに置き換えてください。
        // This URL should be replaced with your actual server endpoint.
        api_url: '/api/sales-data?date=',

        init() {
            // Watch for changes to the selectedDate property
            this.$watch('selectedDate', (newDate) => {
                this.fetchData(newDate);
            });

            // Initial data fetch
            this.fetchData(this.selectedDate);
        },

        // サーバーからデータを取得する関数
        async fetchData(date) {
            this.loading = true;
            this.error = null;
            try {
                const response = await fetch(`${this.api_url}${date}`);
                if (!response.ok) {
                    throw new Error('ネットワーク応答が正常ではありませんでした');
                }
                const rawData = await response.json();
                
                // Golang の []uint8 を JavaScript の文字列に変換
                if (rawData.Date) {
                    rawData.Date = new TextDecoder().decode(new Uint8Array(rawData.Date));
                }
                
                this.data = rawData;
                console.log('Fetched Data:', this.data);
            } catch (e) {
                this.error = `データの取得に失敗しました: ${e.message}`;
                console.error('Error fetching data:', e);
            } finally {
                this.loading = false;
            }
        },

        // 数値にカンマ区切りを追加するgetter
        get formattedData() {
            // if (this.loading || !this.data) {
            //     return {};
            // }

            const formatted = {};
            formatted.Color = this.data.Color || 'green'; // デフォルトの色を設定
            formatted.Found = this.data.Found || false; // デフォルトのFoundを設定
            formatted.Mode = this.data.Mode || '登録'; // デフォルトのModeを設定
            formatted.Record = {};
            for (const key in this.data.Record) {
                // `Date`のような文字列や`null`値をスキップ
                if (typeof this.data.Record[key] === 'number') {
                    // 日本語のロケールを使用してカンマを追加
                    formatted.Record[key] = this.data.Record[key].toLocaleString('ja-JP');
                    //console.log(`Formatted ${key}:`, formatted[key], `typeof ${this.data[key]} :`, typeof this.data[key]);
                } else {
                    // 数値でない場合はそのまま値をコピー
                    formatted.Record[key] = this.data.Record[key];
                }
            }
            return formatted;
        },
    
        // This is a calculated property, not a simple variable
        get CashAmountTotal() {
            let total = 0;
            for (let i = 1; i <= 5; i++) {
                let val = this.data.Record[`Machine${i}CashAmount`];
                total += Number(val) || 0;
            }
            return total.toLocaleString('ja-JP');
        },

        // New method to sync data with raw input
        updateRecord(key, value) {
            // Strip commas and convert to number
            const numValue = Number(value.toString().replace(/,/g, ''));
            if (!isNaN(numValue)) {
                this.data.Record[key] = numValue;
            }
            this.updateCashAmountTotal();
        }
        }));
});
