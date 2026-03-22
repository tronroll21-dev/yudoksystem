document.addEventListener('alpine:init', () => {
    Alpine.data('powerReadings', () => ({
        year: 2026,
        month: 3,
        readings: [],
        baseline: {
          label: "",
          power_reading: 0
        },
        isLoading: false,
        currentUser: null,
        manualTotalWithTax: 0,
        manualTax: 0,

        async init() {
            await this.fetchUser();
            await this.fetchData();
        },

        get title() {
            return `${this.year}年${this.month}月度`;
        },

        async fetchUser() {
            try {
                const res = await fetch('/api/user/me');
                if (res.ok) {
                    this.currentUser = await res.json();
                }
            } catch (error) {
                console.error('Error fetching user:', error);
            }
        },

        async fetchData() {
            this.isLoading = true;
            try {
                // Fetch data for the current fiscal month. 
                // The Go backend also returns the 21st of the previous month for the baseline.
                const res = await fetch(`/api/power-readings?year=${this.year}&month=${this.month}`);
                const data = await res.json();
                this.generateFiscalMonth(data.Record || []);
            } catch (error) {
                console.error('Error fetching data:', error);
            } finally {
                this.isLoading = false;
            }
        },

        generateFiscalMonth(apiData) {
            // Calculate previous month/year for the baseline (21st)
            let prevMonthNum = this.month - 1;
            let prevYearNum = this.year;
            if (prevMonthNum === 0) {
                prevMonthNum = 12;
                prevYearNum = this.year - 1;
            }

            // Find baseline record (21st of previous month)
            const baselineRecord = apiData.find(r => r.year === prevYearNum && r.month === prevMonthNum && r.day === 21);
            console.log('Baseline Record:', baselineRecord);

            const rows = [];
            // 1. Add "前月末" (Baseline)

            this.baseline = {
                label: "前月末",
                power_reading: baselineRecord.power_reading
            };

            // 2. Generate Fiscal Days: 22nd of prev month to 21st of current month
            // As per instruction: these records are fetched for the current "month" (e.g., month=3)
            const startDate = new Date(this.year, this.month - 2, 22); 
            const endDate = new Date(this.year, this.month - 1, 21);   
            
            let currentDate = new Date(startDate);

            while (currentDate <= endDate) {
                const d = currentDate.getDate();

                // Match by day within the records that belong to the current selected fiscal month
                const record = apiData.find(r => r.year === this.year && r.month === this.month && r.day === d);
                
                rows.push({
                    isBaseline: false,
                    date: new Date(currentDate),
                    year: this.year,
                    month: this.month,
                    day: d,
                    power_reading: record ? record.power_reading : null,
                    author: record ? record.author : '',
                    memo: record ? (record.memo && record.memo.Valid ? record.memo.String : '') : ''
                });

                currentDate.setDate(currentDate.getDate() + 1);
            }

            this.readings = rows;
        },

        async prevMonth() {
            if (this.month === 1) {
                this.month = 12;
                this.year--;
            } else {
                this.month--;
            }
            await this.fetchData();
        },

        async nextMonth() {
            if (this.month === 12) {
                this.month = 1;
                this.year++;
            } else {
                this.month++;
            }
            await this.fetchData();
        },

        updateAuthor(index) {
            if (this.currentUser && this.currentUser.username) {
                this.readings[index].author = this.currentUser.username;
            }
        },

        async saverow(data) {
            // Ensure we have an author
            if (!data.author && this.currentUser) {
                data.author = this.currentUser.username;
            }

            try {

                let id_month = (data.day >= 22 && data.month != 1) ? data.month - 1 : data.month;
                let mdd = id_month * 100 + data.day; // e.g., 122 for Jan 22, 301 for Mar 1
                let id_year = (mdd >= 922) ? data.year -1 : data.year;
                let id = `${id_year}${String(mdd).padStart(4, '0')}`; // e.g., "2023122" for Jan 22, 2024

                const response = await fetch('/api/power-readings', {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json',
                    },
                    body: JSON.stringify({
                        ID: data.ID || id,
                        year: parseInt(data.year),
                        month: parseInt(data.month),
                        day: parseInt(data.day),
                        power_reading: parseFloat(data.power_reading),
                        author: data.author || '',
                        memo: data.memo || ''
                    })
                });

                if (!response.ok) {
                    throw new Error('Failed to save data');
                }
                console.log('Row saved successfully');
            } catch (error) {
                console.error('Error saving row:', error);
                alert('データの保存に失敗しました。');
            }
        },

        getFormattedDate(row) {
            const day = row.day;
            const date = row.date;
            if (day === 22 || day === 1) {
                return `${date.getMonth() + 1}/${day}`;
            }
            return day.toString();
        },

        getJapaneseWeekday(date) {
            const kanji = ["日", "月", "火", "水", "木", "金", "土"];
            return kanji[date.getDay()];
        },

        getDiff(index) {
            if (this.readings[index].isBaseline) return 0;

            const current = parseFloat(this.readings[index].power_reading);
            if (isNaN(current)) return 0;

            const prevRow = index == 0 ? this.baseline : this.readings[index - 1];
            const prevValue = parseFloat(prevRow.power_reading);
            if (isNaN(prevValue)) return 0;

            const diff = current - prevValue;
            // Instruction: floor((current - previous) / 10) / 100
            if(this.year <= 2025 || (
                this.year === 2026 && (
                    this.readings[index].month >= 10 ||
                    this.readings[index].month === 1 || (
                        this.readings[index].month === 2 && (this.readings[index].day >= 22 || this.readings[index].day <= 3)
                    ) 
                )
            )) {
                return diff;
            }
            return Math.floor(diff / 10) / 100;
        },

        getValue(index) {
            return this.getDiff(index) * 600;
        },

        getRunningTotal(index) {
            let total = 0;
            // Sum from the first fiscal day (index 0) to the current index
            for (let i = 0; i <= index; i++) {
                total += this.getValue(i);
            }
            return total;
        },

        getAnnualTotal(index) {
            const previousAnnual = 4555914;
            return previousAnnual + this.getRunningTotal(index);
        },

        getFee(index) {
            if (this.readings[index].isBaseline) return 0;
            return Math.round(this.getValue(index) * 25.77);
        },

        getSumOfFees() {
            return this.readings.reduce((sum, _, i) => sum + (this.readings[i].isBaseline ? 0 : this.getFee(i)), 0);
        },

        getCalculatedTax() {
            return Math.floor(this.getSumOfFees() * 0.05);
        },

        getCalculatedTotalWithTax() {
            return this.getSumOfFees() + this.getCalculatedTax();
        },

        getManualTaxExcl() {
            return this.manualTotalWithTax - this.manualTax;
        }
    }));
});
