document.addEventListener('alpine:init', () => {
    Alpine.data('dailyReport', () => ({
        data: {
            Record: {
                ID: 0,
                DateString: '',
                WeatherCode: 0,
                StaffCode: 0,
                Machine1CashCount: 0,
                Machine1CashAmount: 0,
                Machine1CashCount: 0,
                Machine1CashAmount: 0,
                Machine1SettleCount: 0,
                Machine1SettleAmount: 0,
                Machine2CashCount: 0,
                Machine2CashAmount: 0,
                Machine2CashCount: 0,
                Machine2CashAmount: 0,
                Machine2SettleCount: 0,
                Machine2SettleAmount: 0,
                Machine3CashCount: 0,
                Machine3CashAmount: 0,
                Machine3CashCount: 0,
                Machine3CashAmount: 0,
                Machine3CashAmount: 0,
                Machine3SettleCount: 0,
                Machine3SettleAmount: 0,
                Machine4CashCount: 0,
                Machine4CashAmount: 0,
                Machine4CashCount: 0,
                Machine4CashAmount: 0,
                Machine4SettleCount: 0,
                Machine4SettleAmount: 0,
                Machine5CashCount: 0,
                Machine5CashAmount: 0,
                Machine5CashCount: 0,
                Machine5CashAmount: 0,
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
                TicketCount: 0,
                SixTicketCount: 0,
                MaleTicketCount: 0,
                FemaleTicketCount: 0,
                MaleTicketShare: '0',
                FemaleTicketShare: '0',
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
        toast: {
            show: false,
            message: '',
            type: 'success'
        },
        tenkiOptions: [ 
            "なし",
            "晴れ",
            "曇り",
            "雨",
            "晴れのち曇り",
            "曇りのち雨",
            "雨のち曇り",
            "雪",
            "雨台風" 
        ],                  
        // Datepicker state
        showModal: false,
        selectedDate: '', // This will hold 'YYYY-MM-DD'
        currentMonth: new Date().getMonth(),
        currentYear: new Date().getFullYear(),
        loading: true,
        error: null,
        tantoushas: {},
        // フェッチするデータのURL。実際のサーバーエンドポイントに置き換えてください。
        // This URL should be replaced with your actual server endpoint.
        api_url: '/api/sales-data?date=',
        api_tantousha_url: '/api/tantoushas',
        get tenkiIconUrl() {
            const code = this.data.Record.WeatherCode;
            if (code >= 1) {
                return `/assets/img/${this.tenkiOptions[code]}.svg`;
            }
            return `data:image/gif;base64,R0lGODlhAQABAIAAAP///wAAACH5BAEAAAAALAAAAAABAAEAAAICRAEAOw==`;
        },

        get currentMonthYear() {
            return `${this.currentYear}年${this.currentMonth + 1}月`;
        },

        get calendarDays() {
            const days = [];
            const firstDay = new Date(this.currentYear, this.currentMonth, 1).getDay();
            const daysInMonth = new Date(this.currentYear, this.currentMonth + 1, 0).getDate();
            
            // Month before empty cell
            for (let i = 0; i < firstDay; i++) {
                days.push({ date: null, label: '' });
            }
            
            // Month date
            for (let i = 1; i <= daysInMonth; i++) {
                days.push({ 
                    date: new Date(this.currentYear, this.currentMonth, i),
                    label: i 
                });
            }
            
            return days;
        },

        get selectedDateDisplay() {
            if (!this.selectedDate) return '日付を選択';
            const [year, month, day] = this.selectedDate.split('-');
            return `${year}年${parseInt(month)}月${parseInt(day)}日`;
        },
        
        isSelected(date) {
            if (!this.selectedDate) return false;
            const selectedDateObj = new Date(this.selectedDate);
            return date.getFullYear() === selectedDateObj.getFullYear() &&
                   date.getMonth() === selectedDateObj.getMonth() &&
                   date.getDate() === selectedDateObj.getDate();
        },
        
        isToday(date) {
            const today = new Date();
            return date.getFullYear() === today.getFullYear() &&
                   date.getMonth() === today.getMonth() &&
                   date.getDate() === today.getDate();
        },

        userId: null,

        async init() {
            const today = new Date();
            const year = today.getFullYear();
            const month = (today.getMonth() + 1).toString().padStart(2, '0');
            const day = today.getDate().toString().padStart(2, '0');
            this.selectedDate = `${year}-${month}-${day}`;

            this.data.Record.DateString = this.selectedDate;
            this.$dispatch('update-date', { data: this.selectedDate });
            
            this.fetchTantoushas();
            await this.fetchCurrentUser();
            this.fetchData(this.selectedDate);
        },

        selectDate(date) {
            const year = date.getFullYear();
            const month = (date.getMonth() + 1).toString().padStart(2, '0');
            const day = date.getDate().toString().padStart(2, '0');
            this.selectedDate = `${year}-${month}-${day}`;
            this.showModal = false;

            // Wait for the DOM to update (for the modal to close) before fetching
            this.$nextTick(() => {
                this.fetchData(this.selectedDate);
                this.data.Record.DateString = this.selectedDate;
                this.$dispatch('update-date', { data: this.selectedDate });
            });
        },
        
        previousMonth() {
            if (this.currentMonth === 0) {
                this.currentMonth = 11;
                this.currentYear--;
            } else {
                this.currentMonth--;
            }
        },
        
        nextMonth() {
            if (this.currentMonth === 11) {
                this.currentMonth = 0;
                this.currentYear++;
            } else {
                this.currentMonth++;
            }
        },

        focusNextTabbable(element) {
            //const tabbableElements = this.$el.querySelectorAll('[tabindex])');
            document.querySelector('[tabindex="' + (element.tabIndex + 1) + '"]').focus();
        },

        showToast(message, type = 'success') {
            this.toast.message = message;
            this.toast.type = type;
            this.toast.show = true;
            setTimeout(() => {
                this.toast.show = false;
            }, 2000);
        },

        async fetchCurrentUser() {
            console.log('Fetching current user data...');
            try {
                debugger;
                console.log('Before fetch current user response:');
                const response = await fetch('/api/user/me', { credentials: 'include' });
                console.log('Fetch current user response:', response);
                if (!response.ok) {
                    throw new Error('Not authenticated');
                }
                const userData = await response.json();
                console.log('Fetched user data:', userData);
                if (userData.username) {
                    // Assuming tantoushas is already loaded or will be loaded
                    // Find the user ID from the tantoushas list based on the username
                    const currentUser = Object.values(this.tantoushas).find(t => t.name === userData.username);
                    if (currentUser && currentUser.id) {
                        this.userId = currentUser.id;
                    } else {
                        console.warn('Current user found, but ID not found in tantoushas list.');
                        throw new Error('User ID not found in tantoushas list');
                    }
                } else {
                    throw new Error('Username not found in response');
                }
            } catch (e) {
                console.error('Failed to fetch user data, redirecting to login.', e);
                window.location.href = '/login';
            }
        },

        // サーバーからデータを取得する関数
        async fetchTantoushas() {
            this.loading = true;
            this.error = null;
            try {
                const response = await fetch(`${this.api_tantousha_url}`, { credentials: 'include' });
                if (!response.ok) {
                    throw new Error('ネットワーク応答が正常ではありませんでした');
                }
                const tantoushas = await response.json();
                               
                this.tantoushas = JSON.parse(tantoushas);
            } catch (e) {
                this.error = `データの取得に失敗しました: ${e.message}`;
                console.error('Error fetching data:', e);
            } finally {
                this.loading = false;
            }
        },

        async fetchData(date) {
            this.loading = true;
            this.error = null;
            try {
                const response = await fetch(`${this.api_url}${date}`, { credentials: 'include' });
                if (!response.ok) {
                    throw new Error('ネットワーク応答が正常ではありませんでした');
                }
                const rawData = await response.json();
                
                // Golang の []uint8 を JavaScript の文字列に変換
                // if (rawData.Date) {
                //     rawData.Date = new TextDecoder().decode(new Uint8Array(rawData.Date));
                // }
                if (rawData.Found && rawData.Record.DateString != date) {
                    console.error('Fetched record date does not match requested date.');
                }

                if (!rawData.Found) {
                    rawData.Record.DateString = date;
                }

                this.data = rawData;

                if (this.data.Found === false && this.data.Mode === '登録') {
                    this.data.Record.StaffCode = this.userId;
                }
                console.log('Fetched Data:', this.data);
            } catch (e) {
                this.error = `データの取得に失敗しました: ${e.message}`;
                console.error('Error fetching data:', e);
            } finally {
                this.loading = false;
            }
        },    

        get Totals() {
            let totals = {};
            let cashCountTotal = 0, cashAmountTotal = 0, settleCountTotal = 0, settleAmountTotal = 0,
                unsettledCountTotal = 0, unsettledAmountTotal = 0,
                qrCountTotal = 0, qrAmountTotal = 0,
                qrSettleCountTotal = 0, qrSettleAmountTotal = 0,
                eCountTotal = 0, eAmountTotal = 0,
                eSettleCountTotal = 0, eSettleAmountTotal = 0,
                cCountTotal = 0, cAmountTotal = 0,
                cSettleCountTotal = 0, cSettleAmountTotal = 0;
            let curCashCountTotal = 0, curCashAmountTotal = 0, curSettleCountTotal = 0, curSettleAmountTotal = 0,
                curUnsettledCountTotal = 0, curUnsettledAmountTotal = 0,
                curQrCountTotal = 0, curQrAmountTotal = 0,
                curQrSettleCountTotal = 0, curQrSettleAmountTotal = 0,
                curECountTotal = 0, curEAmountTotal = 0,
                curESettleCountTotal = 0, curESettleAmountTotal = 0,
                curCCountTotal = 0, curCAmountTotal = 0,
                curCSettleCountTotal = 0, curCSettleAmountTotal = 0;
            let netCashCountTotal = 0,
                netCashAmountTotal = 0,
                netCashCountMinusUnsettledTotal = 0,
                netCashAmountMinusUnsettledTotal = 0,
                netECountTotal = 0,
                netEAmountTotal = 0,
                netCCountTotal = 0,
                netCAmountTotal = 0,
                netQrCountTotal = 0,
                netQrAmountTotal = 0;
            for (let i = 1; i <= 5; i++) {
                curCashCountTotal = Number(this.data.Record[`Machine${i}CashCount`]) || 0;
                curCashAmountTotal = Number(this.data.Record[`Machine${i}CashAmount`]) || 0;
                curSettleCountTotal = Number(this.data.Record[`Machine${i}SettleCount`]) || 0;
                curSettleAmountTotal = Number(this.data.Record[`Machine${i}SettleAmount`]) || 0;
                curUnsettledCountTotal = Number(this.data.Record[`Machine${i}UnsettledCount`]) || 0;
                curUnsettledAmountTotal = Number(this.data.Record[`Machine${i}UnsettledAmount`]) || 0;
                curQrCountTotal = Number(this.data.Record[`Machine${i}QrCount`]) || 0;
                curQrAmountTotal = Number(this.data.Record[`Machine${i}QrAmount`]) || 0;
                curQrSettleCountTotal = Number(this.data.Record[`Machine${i}QrSettleCount`]) || 0;
                curQrSettleAmountTotal = Number(this.data.Record[`Machine${i}QrSettleAmount`]) || 0;
                curECountTotal = Number(this.data.Record[`Machine${i}ECount`]) || 0;
                curEAmountTotal = Number(this.data.Record[`Machine${i}EAmount`]) || 0;
                curESettleCountTotal = Number(this.data.Record[`Machine${i}ESettleCount`]) || 0;
                curESettleAmountTotal = Number(this.data.Record[`Machine${i}ESettleAmount`]) || 0;
                curCCountTotal = Number(this.data.Record[`Machine${i}CCount`]) || 0;
                curCAmountTotal = Number(this.data.Record[`Machine${i}CAmount`]) || 0;
                curCSettleCountTotal = Number(this.data.Record[`Machine${i}CSettleCount`]) || 0;
                curCSettleAmountTotal = Number(this.data.Record[`Machine${i}CSettleAmount`]) || 0;
                cashCountTotal += curCashCountTotal;
                cashAmountTotal += curCashAmountTotal;
                settleCountTotal += curSettleCountTotal;
                settleAmountTotal += curSettleAmountTotal;
                unsettledCountTotal += curUnsettledCountTotal;
                unsettledAmountTotal += curUnsettledAmountTotal;
                qrCountTotal += curQrCountTotal;
                qrAmountTotal += curQrAmountTotal;
                qrSettleCountTotal += curQrSettleCountTotal;
                qrSettleAmountTotal += curQrSettleAmountTotal;
                eCountTotal += curECountTotal;
                eAmountTotal += curEAmountTotal;
                eSettleCountTotal += curESettleCountTotal;
                eSettleAmountTotal += curESettleAmountTotal;
                cCountTotal += curCCountTotal;
                cAmountTotal += curCAmountTotal;
                cSettleCountTotal += curCSettleCountTotal;
                cSettleAmountTotal += curCSettleAmountTotal;
                netCashCountTotal += (curCashCountTotal - curSettleCountTotal);
                netCashAmountTotal += (curCashAmountTotal - curSettleAmountTotal);
                netCashCountMinusUnsettledTotal += (curCashCountTotal - curSettleCountTotal - curUnsettledCountTotal);
                netCashAmountMinusUnsettledTotal += (curCashAmountTotal - curSettleAmountTotal - curUnsettledAmountTotal);
                netECountTotal += (curECountTotal - curESettleCountTotal);
                netEAmountTotal += (curEAmountTotal - curESettleAmountTotal);
                netCCountTotal += (curCCountTotal - curCSettleCountTotal);
                netCAmountTotal += (curCAmountTotal - curCSettleAmountTotal);
                netQrCountTotal += (curQrCountTotal - curQrSettleCountTotal);
                netQrAmountTotal += (curQrAmountTotal - curQrSettleAmountTotal);
                totals[`Machine${i}NetCashCount`] = (curCashCountTotal - curSettleCountTotal).toLocaleString('ja-JP'); 
                totals[`Machine${i}NetCashAmount`] = (curCashAmountTotal - curSettleAmountTotal).toLocaleString('ja-JP'); 
                totals[`Machine${i}NetCashCountMinusUnsettled`] = (curCashCountTotal - curSettleCountTotal - curUnsettledCountTotal).toLocaleString('ja-JP'); 
                totals[`Machine${i}NetCashAmountMinusUnsettled`] = (curCashAmountTotal - curSettleAmountTotal - curUnsettledAmountTotal).toLocaleString('ja-JP'); 
                totals[`Machine${i}NetECount`] = (curECountTotal - curESettleCountTotal).toLocaleString('ja-JP'); 
                totals[`Machine${i}NetEAmount`] = (curEAmountTotal - curESettleAmountTotal).toLocaleString('ja-JP'); 
                totals[`Machine${i}NetCCount`] = (curCCountTotal - curCSettleCountTotal).toLocaleString('ja-JP'); 
                totals[`Machine${i}NetCAmount`] = (curCAmountTotal - curCSettleAmountTotal).toLocaleString('ja-JP'); 
                totals[`Machine${i}NetQrCount`] = (curQrCountTotal - curQrSettleCountTotal).toLocaleString('ja-JP'); 
                totals[`Machine${i}NetQrAmount`] = (curQrAmountTotal - curQrSettleAmountTotal).toLocaleString('ja-JP');
                totals[`Machine${i}CashlessTotal`] = (curQrAmountTotal - curQrSettleAmountTotal + curEAmountTotal - curESettleAmountTotal).toLocaleString('ja-JP');
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
            totals.TounyuuGoukei = (netCashAmountTotal + 
                Number(this.data.Record.Change) -
                (Number(this.data.Record.Machine1UnsettledAmount) +
                 Number(this.data.Record.Machine2UnsettledAmount) +
                 Number(this.data.Record.Machine3UnsettledAmount) +
                 Number(this.data.Record.Machine4UnsettledAmount) +
                 Number(this.data.Record.Machine5UnsettledAmount)) +
                Number(this.data.Record.PhoneFee) - 
                Number(this.data.Record.HonjitsuMitounyuuAmountUncertain) -
                Number(this.data.Record.HonjitsuMitounyuuAmountCertain) +
                Number(this.data.Record.Deficiency) +
                Number(this.data.Record.ZenjitsuMitounyuuAmount)
            ).toLocaleString('ja-JP');

            return totals;
        },

get TicketTotals() {
    const totals = {
 
        MaleTicketCountTotalNumber: Number(this.data.Record.MaleTicketCount) +
                                    Number(this.data.Record.SixMaleTicketCount),

        FemaleTicketCountTotalNumber: Number(this.data.Record.FemaleTicketCount) +
                                      Number(this.data.Record.SixFemaleTicketCount)
    };
 
    return {
        TicketCountTotal: (Number(this.data.Record.MaleTicketCount) +
                           Number(this.data.Record.FemaleTicketCount)).toLocaleString('ja-JP'),
        
        SixTicketCountTotal: (Number(this.data.Record.SixMaleTicketCount) +
                              Number(this.data.Record.SixFemaleTicketCount)).toLocaleString('ja-JP'),
        
        MaleTicketCountTotal: totals.MaleTicketCountTotalNumber.toLocaleString('ja-JP'),
        
        FemaleTicketCountTotal: totals.FemaleTicketCountTotalNumber.toLocaleString('ja-JP'),

        TicketCountGrandTotal: (totals.MaleTicketCountTotalNumber + totals.FemaleTicketCountTotalNumber).toLocaleString('ja-JP'),

       MaleTicketShare: (Math.round(totals.MaleTicketCountTotalNumber/(totals.MaleTicketCountTotalNumber + totals.FemaleTicketCountTotalNumber)*100)) ,
       FemaleTicketShare: (Math.round(totals.FemaleTicketCountTotalNumber/(totals.MaleTicketCountTotalNumber + totals.FemaleTicketCountTotalNumber)*100)),

       SCutTotal: (Number(this.data.Record.SCutMale) +
                   Number(this.data.Record.SCutFemale) +
                   Number(this.data.Record.SCutChild)).toLocaleString('ja-JP')
    }
},

        // New method to save record
        async saveRecord() {
            try {
                // Remove formatting from numbers before sending to backend
                const recordToSend = JSON.parse(JSON.stringify(this.data.Record)); // Deep copy

                // If Date is a string, convert it to a format the backend expects
                // if (typeof recordToSend.DateString === 'string') {
                //     recordToSend.DateString = this.selectDate;
                // }

                debugger;

                const response = await fetch('/api/sales-data', {
                    method: 'POST',
                    credentials: 'include',
                    headers: {
                        'Content-Type': 'application/json',
                    },
                    body: JSON.stringify({
                        Record: recordToSend,
                        Found: this.data.Found,
                    }),
                });

                if (!response.ok) {
                    throw new Error(`サーバーエラー: ${response.status} ${response.statusText}`);
                }

                const responseData = await response.json();
                this.data = responseData; // Update frontend data with the new record (including ID if newly created)
                console.log('Record saved successfully:', responseData);
                this.showToast('レコードが正常に保存されました！', 'success');
            } catch (error) {
                console.error('レコードの保存中にエラーが発生しました:', error);
                this.showToast(`レコードの保存に失敗しました: ${error.message}`, 'error');
            }
        },

        updateKenbaikiData(combinedData) {
            
            for(const key in combinedData.paymentStats) {
                if (combinedData.paymentStats.hasOwnProperty(key)) {
         
                    this.data.Record['Machine' + key + 'CashCount'] = combinedData.paymentStats[key].GenkinGrossMaisuu;
                    this.data.Record['Machine' + key + 'CashAmount'] = combinedData.paymentStats[key].GenkinGrossKingaku;
                    this.data.Record['Machine' + key + 'QrCount'] = combinedData.paymentStats[key].QrcodeGrossMaisuu;
                    this.data.Record['Machine' + key + 'QrAmount'] = combinedData.paymentStats[key].QrcodeGrossKingaku;
                    this.data.Record['Machine' + key + 'ECount'] = combinedData.paymentStats[key].DenshimaneeGrossMaisuu;
                    this.data.Record['Machine' + key + 'EAmount'] = combinedData.paymentStats[key].DenshimaneeGrossKingaku;

                    this.data.Record['Machine' + key + 'CashSettleCount'] = combinedData.paymentStats[key].GenkinSeisanMaisuu;
                    this.data.Record['Machine' + key + 'CashSettleAmount'] = combinedData.paymentStats[key].GenkinSeisanKingaku;
                    this.data.Record['Machine' + key + 'QrSettleCount'] = combinedData.paymentStats[key].QrcodeSeisanMaisuu;
                    this.data.Record['Machine' + key + 'QrSettleAmount'] = combinedData.paymentStats[key].QrcodeSeisanKingaku;
                }
            }

        }

        }));
});
